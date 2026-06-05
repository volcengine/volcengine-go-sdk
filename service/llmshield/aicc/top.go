package jeddak_secure_channel

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/sts"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

const timeFormatV4 = "20060102T150405Z"

// TopInfo holds configuration for the TOP API request.
type TopInfo struct {
	AK          string `json:"ak"`           // Access Key
	SK          string `json:"sk"`           // Secret Key
	Service     string `json:"service"`      // Service name (required)
	Region      string `json:"region"`       // Region (default: "cn-beijing")
	Method      string `json:"method"`       // HTTP method (default: "POST")
	Action      string `json:"action"`       // Deprecated: Use action parameter in RequestTop
	Version     string `json:"version"`      // API version (default: "2024-12-24")
	URL         string `json:"url"`          // Base URL (for signing)
	URLRewrite  string `json:"url_rewrite"`  // Optional URL rewrite (for actual request)
	AiccSaaSTrn string `json:"aicc_saas_trn"`// Optional: STS AssumeRole TRN
	TargetUID   string `json:"target_uid"`   // Optional: Target UID for header
}

// tempCredentials holds STS AssumeRole result
type tempCredentials struct {
	AK           string
	SK           string
	SessionToken string
}

type signingMetadata struct {
	algorithm       string
	date            string
	service         string
	region          string
	signedHeaders   string
	credentialScope string
}

// deserializeTopInfo parses a JSON reader into a TopInfo struct with defaults.
func deserializeTopInfo(reader io.Reader) (TopInfo, error) {
	var info TopInfo
	if err := json.NewDecoder(reader).Decode(&info); err != nil {
		return TopInfo{}, err
	}
	// Set default values
	if info.Region == "" {
		info.Region = "cn-beijing"
	}
	if info.Method == "" {
		info.Method = "POST"
	}
	if info.Version == "" {
		info.Version = "2024-12-24"
	}
	return info, nil
}

func requestTop(topInfo *TopInfo, action string, extraHeaders map[string]string, body []byte) (map[string]interface{}, error) {
	// Validate required fields
	if topInfo.AK == "" {
		return nil, errors.New("ak is required")
	}
	if topInfo.SK == "" {
		return nil, errors.New("sk is required")
	}
	if topInfo.Service == "" {
		return nil, errors.New("service is required")
	}

	// If aicc_saas_trn is present, get temporary credentials via STS AssumeRole
	effectiveAK := topInfo.AK
	effectiveSK := topInfo.SK
	var sessionToken string

	if topInfo.AiccSaaSTrn != "" {
		tempCreds, err := assumeRole(topInfo)
		if err != nil {
			return nil, fmt.Errorf("STS AssumeRole failed: %w", err)
		}
		effectiveAK = tempCreds.AK
		effectiveSK = tempCreds.SK
		sessionToken = tempCreds.SessionToken
	}

	// Build query string (Action and Version)
	query := url.Values{}
	query.Add("Action", action)
	query.Add("Version", topInfo.Version)

	// Determine URLs: use original URL for signing, URLRewrite for actual request
	signingURLStr := topInfo.URL
	requestURLStr := topInfo.URL
	if topInfo.URLRewrite != "" {
		requestURLStr = topInfo.URLRewrite
	}

	// Ensure URLs have scheme
	if !strings.HasPrefix(signingURLStr, "http://") && !strings.HasPrefix(signingURLStr, "https://") {
		signingURLStr = "https://" + signingURLStr
	}
	if !strings.HasPrefix(requestURLStr, "http://") && !strings.HasPrefix(requestURLStr, "https://") {
		requestURLStr = "https://" + requestURLStr
	}

	// Parse signing URL (for Host header and canonical request)
	signingURL, err := url.Parse(signingURLStr + "?" + query.Encode())
	if err != nil {
		return nil, fmt.Errorf("invalid signing URL: %w", err)
	}

	// Parse request URL (for actual HTTP request)
	requestURL, err := url.Parse(requestURLStr + "?" + query.Encode())
	if err != nil {
		return nil, fmt.Errorf("invalid request URL: %w", err)
	}

	// Build authentication headers (use signing URL for Host header)
	// Note: X-Security-Token is NOT included in signed headers (matches Python)
	headers, err := buildTopHeaders(effectiveAK, effectiveSK, topInfo.Service, topInfo.Region, topInfo.Method, signingURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to build headers: %w", err)
	}

	// Add X-Security-Token AFTER signing (matches Python behavior)
	if sessionToken != "" {
		headers["X-Security-Token"] = sessionToken
	}

	// Create HTTP request (use request URL)
	req, err := http.NewRequest(topInfo.Method, requestURL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	for k, v := range extraHeaders {
		req.Header.Set(k, v)
	}
	// Add UID header if target_uid is specified
	if topInfo.TargetUID != "" {
		req.Header.Set("UID", topInfo.TargetUID)
	}

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read and parse response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var respJson map[string]interface{}
	if err := json.Unmarshal(respBody, &respJson); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	result, ok := respJson["Result"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("missing or invalid Result in response: %s", string(respBody))
	}
	return result, nil
}

// assumeRole calls STS AssumeRole to get temporary credentials
func assumeRole(topInfo *TopInfo) (*tempCredentials, error) {
	// Create STS client with explicit credentials
	creds := credentials.NewStaticCredentials(topInfo.AK, topInfo.SK, "")
	sess, err := session.NewSession(&volcengine.Config{
		Credentials: creds,
		Region:      volcengine.String(topInfo.Region),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	stsClient := sts.New(sess)

	// Build AssumeRole request
	durationSeconds := int32(3600)
	roleSessionName := "temp_aicc_" + fmt.Sprintf("%d", time.Now().Unix())

	input := &sts.AssumeRoleInput{
		DurationSeconds: &durationSeconds,
		RoleSessionName: &roleSessionName,
		RoleTrn:         &topInfo.AiccSaaSTrn,
	}

	// Call AssumeRole
	output, err := stsClient.AssumeRole(input)
	if err != nil {
		return nil, fmt.Errorf("AssumeRole failed: %w", err)
	}

	if output.Credentials == nil {
		return nil, errors.New("AssumeRole returned nil credentials")
	}

	return &tempCredentials{
		AK:           *output.Credentials.AccessKeyId,
		SK:           *output.Credentials.SecretAccessKey,
		SessionToken: *output.Credentials.SessionToken,
	}, nil
}

// buildTopHeaders constructs headers required for TOP API authentication.
// This follows the Python implementation in request_bytedance_gateway.
// Signed headers: Content-Type, Host, X-Content-Sha256, X-Date
// Note: X-Security-Token is added AFTER signing (not part of signed headers)
func buildTopHeaders(ak, sk, service, region, method string, reqURL *url.URL, body []byte) (map[string]string, error) {
	now := time.Now().UTC()
	nowDateTime := now.Format(timeFormatV4)
	nowDate := nowDateTime[:8]

	// Initialize headers map - must match Python order for consistent signing
	// Python order: Content-Type, Host, X-Content-Sha256, X-Date
	headers := make(map[string]string)

	// Content-Type (must be included in signed headers - matches Python)
	headers["Content-Type"] = "application/json"

	// Set Host header
	host := reqURL.Host
	// Strip port if it's default (80 for http, 443 for https)
	if strings.Contains(host, ":") {
		parts := strings.Split(host, ":")
		port := parts[1]
		if (reqURL.Scheme == "http" && port == "80") || (reqURL.Scheme == "https" && port == "443") {
			host = parts[0]
		}
	}
	headers["Host"] = host

	// Compute body hash and set X-Content-Sha256
	payloadHash := hashSHA256(body)
	headers["X-Content-Sha256"] = payloadHash

	// Set X-Date
	headers["X-Date"] = nowDateTime

	// Determine which headers to sign (following Python logic)
	// Sign: Content-Type, Host, X-Content-Sha256, X-Date
	// Note: X-Security-Token is NOT signed (added after signing)
	var sortedHeaderKeys []string
	for key := range headers {
		keyLower := strings.ToLower(key)
		sortedHeaderKeys = append(sortedHeaderKeys, keyLower)
	}
	sort.Strings(sortedHeaderKeys)

	// Build canonical headers
	var canonicalHeaders strings.Builder
	for _, key := range sortedHeaderKeys {
		// Get the original case value from headers map
		var value string
		for k, v := range headers {
			if strings.ToLower(k) == key {
				value = strings.TrimSpace(v)
				break
			}
		}
		// Special handling for host (strip default port)
		if key == "host" {
			if strings.Contains(value, ":") {
				parts := strings.Split(value, ":")
				port := parts[1]
				if port == "80" || port == "443" {
					value = parts[0]
				}
			}
		}
		canonicalHeaders.WriteString(key + ":" + value + "\n")
	}

	signedHeaders := strings.Join(sortedHeaderKeys, ";")

	// Build canonical request - matches Python format
	// Format: METHOD\nPATH\nQUERY\nCANONICAL_HEADERS\nSIGNED_HEADERS\nPAYLOAD_HASH
	// Python uses "/" as path (hardcoded)
	canonicalPath := "/"
	canonicalQuery := normquery(reqURL.Query())

	canonicalRequest := method + "\n" +
		canonicalPath + "\n" +
		canonicalQuery + "\n" +
		canonicalHeaders.String() + "\n" +
		signedHeaders + "\n" +
		payloadHash

	hashedCanonicalRequest := hashSHA256([]byte(canonicalRequest))

	// Build credential scope
	credentialScope := nowDate + "/" + region + "/" + service + "/request"

	// Build string to sign
	stringToSign := "HMAC-SHA256\n" +
		nowDateTime + "\n" +
		credentialScope + "\n" +
		hashedCanonicalRequest

	// Compute signature
	signingKey := signingKeyV4(sk, nowDate, region, service)
	signature := hex.EncodeToString(hmacSHA256(signingKey, stringToSign))

	// Build Authorization header
	authorization := fmt.Sprintf(
		"HMAC-SHA256 Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		ak, credentialScope, signedHeaders, signature,
	)
	headers["Authorization"] = authorization

	return headers, nil
}

// Helper functions following volc-sdk-golang implementation

func hashSHA256(content []byte) string {
	h := sha256.New()
	h.Write(content)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hmacSHA256(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func signingKeyV4(secretKey, date, region, service string) []byte {
	kDate := hmacSHA256([]byte(secretKey), date)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "request")
	return kSigning
}

func normuri(uri string) string {
	parts := strings.Split(uri, "/")
	for i := range parts {
		parts[i] = encodePathFrag(parts[i])
	}
	return strings.Join(parts, "/")
}

func encodePathFrag(s string) string {
	hexCount := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			hexCount++
		}
	}
	t := make([]byte, len(s)+2*hexCount)
	j := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		} else {
			t[j] = c
			j++
		}
	}
	return string(t)
}

func shouldEscape(c byte) bool {
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
		return false
	}
	if '0' <= c && c <= '9' {
		return false
	}
	if c == '-' || c == '_' || c == '.' || c == '~' {
		return false
	}
	return true
}

func normquery(v url.Values) string {
	queryString := v.Encode()
	return strings.Replace(queryString, "+", "%20", -1)
}

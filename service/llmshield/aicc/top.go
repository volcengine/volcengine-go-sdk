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
)

// TopInfo holds configuration for the TOP API request.
type TopInfo struct {
	AK         string `json:"ak"`          // Access Key
	SK         string `json:"sk"`          // Secret Key
	Service    string `json:"service"`     // Service name (required)
	Region     string `json:"region"`      // Region (default: "cn-beijing")
	Method     string `json:"method"`      // HTTP method (default: "POST")
	Action     string `json:"action"`      // Deprecated: Use action parameter in RequestTop
	Version    string `json:"version"`     // API version (default: "2024-12-24")
	URL        string `json:"url"`         // Base URL
	URLRewrite string `json:"url_rewrite"` // Optional URL rewrite
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

	// Build query string (Action and Version)
	query := url.Values{}
	query.Add("Action", action)
	query.Add("Version", topInfo.Version)
	queryStr := query.Encode()

	// Parse base URL
	urlStr := topInfo.URL + "?" + queryStr
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	// Apply URL rewrite if specified
	if topInfo.URLRewrite != "" {
		urlStr = topInfo.URLRewrite + "?" + queryStr
		parsedURL, err = url.Parse(urlStr)
		if err != nil {
			return nil, fmt.Errorf("invalid URL rewrite: %w", err)
		}
	}

	// Build authentication headers
	headers, err := buildTopHeaders(topInfo, parsedURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to build headers: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest(topInfo.Method, parsedURL.String(), bytes.NewReader(body))
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

	// metadata := respJson["ResponseMetadata"].(map[string]interface{})
	// log.Printf("Response metadata: %+v", metadata)

	results := respJson["Result"].(map[string]interface{})
	return results, nil
}

// buildTopHeaders constructs headers required for TOP API authentication.
func buildTopHeaders(topInfo *TopInfo, url *url.URL, body []byte) (map[string]string, error) {
	now := time.Now().UTC()
	nowDateTime := now.Format("20060102T150405Z") // yyyyMMdd'T'HHmmss'Z'
	nowDate := now.Format("20060102")             // yyyyMMdd

	contentSha256 := sha256Hex(body)

	// Initial headers (will be signed)
	headers := map[string]string{
		"Content-Type":     "application/json",
		"Host":             url.Host,
		"X-Content-Sha256": contentSha256,
		"X-Date":           nowDateTime,
	}

	// Prepare signed headers (lowercase, sorted)
	lowerHeaders := make(map[string]string)
	for k, v := range headers {
		lowerHeaders[strings.ToLower(k)] = v
	}
	signedHeaders := make([]string, 0, len(lowerHeaders))
	for k := range lowerHeaders {
		signedHeaders = append(signedHeaders, k)
	}
	sort.Strings(signedHeaders)
	signedHeadersStr := strings.Join(signedHeaders, ";")

	// Canonical headers (lowercase key:value\n, sorted)
	var canonicalHeaders strings.Builder
	for _, k := range signedHeaders {
		canonicalHeaders.WriteString(k + ":" + lowerHeaders[k] + "\n")
	}
	canonicalHeaders.WriteString("\n") // Extra newline after headers

	// Canonical request: METHOD\nPATH\nQUERY\nCANONICAL_HEADERS\nSIGNED_HEADERS\nCONTENT_SHA256
	urlPathWithoutSlash, _ := strings.CutPrefix(url.Path, "/")
	canonicalRequest := fmt.Sprintf("%s\n/%s\n%s\n%s%s\n%s",
		topInfo.Method,
		urlPathWithoutSlash, // Slash is in the format string
		url.RawQuery,
		canonicalHeaders.String(),
		signedHeadersStr,
		contentSha256,
	)
	hashedCanonicalRequest := sha256Hex([]byte(canonicalRequest))

	// Credential scope: date/region/service/request
	credentialScope := fmt.Sprintf("%s/%s/%s/request", nowDate, topInfo.Region, topInfo.Service)

	// String to sign: HMAC-SHA256\nnowDateTime\ncredentialScope\nhashedCanonicalRequest
	stringToSign := fmt.Sprintf("HMAC-SHA256\n%s\n%s\n%s", nowDateTime, credentialScope, hashedCanonicalRequest)

	// Compute signature via iterative HMAC-SHA256
	signatureData := []string{topInfo.SK, nowDate, topInfo.Region, topInfo.Service, "request", stringToSign}
	signature, err := hmacSha256Iterative(signatureData)
	if err != nil {
		return nil, err
	}

	// Authorization header
	authorization := fmt.Sprintf(
		"HMAC-SHA256 Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		topInfo.AK, credentialScope, signedHeadersStr, signature,
	)
	headers["Authorization"] = authorization

	return headers, nil
}

// sha256Hex computes SHA-256 hash of data and returns hex string.
func sha256Hex(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// hmacSha256Iterative computes iterative HMAC-SHA256 (each step uses previous result as key).
func hmacSha256Iterative(data []string) (string, error) {
	if len(data) == 0 {
		return "", errors.New("data is empty")
	}
	key := []byte(data[0]) // Initial key (SK)
	for i := 1; i < len(data); i++ {
		h := hmac.New(sha256.New, key)
		if _, err := h.Write([]byte(data[i])); err != nil {
			return "", err
		}
		key = h.Sum(nil)
	}
	return hex.EncodeToString(key), nil
}

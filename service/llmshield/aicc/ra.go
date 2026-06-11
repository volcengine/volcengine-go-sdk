package jeddak_secure_channel

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// attestServer performs server attestation and returns a ClientSessionKey.
// Dispatches to local or TCA attestation based on ClientConfig.RaType.
func attestServer(token string, config ClientConfig) (*ClientSessionKey, error) {
	var serverPublicKey string
	var err error

	if config.RaType == RA_TYPE_LOCAL {
		serverPublicKey, err = attestLocal(token, config)
	} else {
		serverPublicKey, err = attestTopTca(token, config)
	}
	if err != nil {
		return nil, fmt.Errorf("attestation failed: %w", err)
	}

	return LoadClientSessionKey(serverPublicKey)
}

// attestLocal performs attestation via direct HTTP request to config.RaUrl.
func attestLocal(token string, config ClientConfig) (string, error) {
	reqBody, err := buildRaRequest(config)
	if err != nil {
		return "", fmt.Errorf("build request: %w", err)
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, config.RaURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Token", token)
	}
	nonce, err := generateNonce(3)
	if err != nil {
		return "", err
	}
	req.Header.Set("Nonce", nonce)

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("attest failed (status %d): %s", resp.StatusCode, resp.Status)
	}

	// Parse response
	raResults, err := parseResponse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("parse response: %w", err)
	}

	return extractPublicKey(raResults)
}

// attestTopTca performs attestation via TopInfo (TCA backend).
func attestTopTca(token string, config ClientConfig) (string, error) {
	topInfo, err := config.parseTopInfo()
	if err != nil {
		return "", fmt.Errorf("parse top info: %w", err)
	}

	reqBody, err := buildRaRequest(config)
	if err != nil {
		return "", fmt.Errorf("build request: %w", err)
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	// Build headers
	headers := map[string]string{
		"Content-Type": "application/json",
		"UID":          config.RaUID,
	}
	if token != "" {
		headers["Token"] = token
	}

	// Send Top request (placeholder implementation)
	raResults, err := requestTop(&topInfo, "GetAttestationBackend", headers, jsonBody)
	if err != nil {
		return "", fmt.Errorf("top request: %w", err)
	}

	return extractPublicKey(raResults)
}

// buildRaRequest constructs the RA request JSON body.
func buildRaRequest(config ClientConfig) (RaRequest, error) {
	nonce, err := generateNonce(8)
	if err != nil {
		return RaRequest{}, fmt.Errorf("generate nonce: %w", err)
	}

	return RaRequest{
		PolicyId:       config.RaPolicyId,
		ServiceName:    config.RaServiceName,
		KeyNegotiation: *config.RaKeyNegotiation,
		Token:          *config.RaNeedToken,
		Nonce:          nonce,
	}, nil
}

// RaRequest represents the JSON structure for RA attestation requests.
type RaRequest struct {
	PolicyId       string `json:"PolicyID"`
	ServiceName    string `json:"ServiceName"`
	KeyNegotiation bool   `json:"KeyNegotiation"`
	Token          bool   `json:"Token"`
	Nonce          string `json:"Nonce"`
}

func addBase64Padding(tokenData string) string {
	// 计算需要补充的填充符数量（和Python逻辑完全一致）
	padding := (4 - len(tokenData)%4) % 4
	// 拼接填充符（Go中用strings.Repeat实现字符串重复）
	return tokenData + strings.Repeat("=", padding)
}

// extractPublicKey parses the RA response to find the server's public key.
func extractPublicKey(raResults map[string]interface{}) (string, error) {
	for _, entry := range raResults {
		raInfo, ok := entry.(map[string]interface{})
		if !ok {
			continue
		}

		validateRaInfo(raInfo)

		// Extract public key from key_info
		keyInfo, ok := raInfo["key_info"].(map[string]interface{})
		if !ok {
			log.Println("RA result missing key_info")
			continue
		}

		pubKey, ok := keyInfo["pub_key_info"].(string)
		if !ok {
			log.Println("key_info missing pub_key_info")
			continue
		}

		return pubKey, nil
	}

	return "", errors.New("no valid RA results found")
}

// validateRaInfo checks RA evidence and token validity.
func validateRaInfo(raInfo map[string]interface{}) bool {
	// Check for evidence
	if evidence, ok := raInfo["evidence"].(string); ok {
		var evidenceJSON map[string]interface{}
		if err := json.Unmarshal([]byte(evidence), &evidenceJSON); err != nil {
			log.Printf("Invalid RA evidence: %v", err)
		}
	} else {
		log.Println("RA response missing evidence")
	}

	// Check token validity
	token, ok := raInfo["token"].(string)
	if !ok {
		log.Println("RA response missing token")
		return false
	}

	// Validate token format (JWT-like X.Y.Z)
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		log.Println("Invalid token format (expected X.Y.Z)")
		return false
	}

	// Decode and parse token body
	paddedTokenData := addBase64Padding(parts[1])
	tokenBody, err := base64.URLEncoding.DecodeString(paddedTokenData)
	if err != nil {
		log.Printf("Failed to decode token: %v", err)
		return false
	}

	var tokenClaims map[string]interface{}
	if err := json.Unmarshal(tokenBody, &tokenClaims); err != nil {
		log.Printf("Failed to parse token claims: %v", err)
		return false
	}

	// Check policy match
	policies, ok := tokenClaims["policies_matched"].([]interface{})
	if !ok || len(policies) == 0 {
		log.Println("Token missing or empty policies_matched")
		return false
	}

	log.Println("RA validation successful")
	return true
}

// generateNonce creates a secure random nonce (hex-encoded).
func generateNonce(numBytes int) (string, error) {
	b := make([]byte, numBytes)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("generate random bytes: %w", err)
	}
	return hex.EncodeToString(b), nil
}

// parseResponse reads and parses JSON from an HTTP response body.
func parseResponse(body io.Reader) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return nil, fmt.Errorf("json decode: %w", err)
	}
	return data, nil
}

package jeddak_secure_channel

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

// Constants for RA types
const (
	RA_TYPE_LOCAL = "local"
	RA_TYPE_TCA   = "tca"
)

// ClientConfig holds configuration parameters for the secure channel client.
type ClientConfig struct {
	RaURL            string   `json:"ra_url"`             // Default: empty string
	RaType           string   `json:"ra_type"`            // Default: RA_TYPE_TCA
	RaServiceName    string   `json:"ra_service_name"`    // Default: empty string
	RaUID            string   `json:"ra_uid"`             // Default: empty string
	RaKeyNegotiation *bool    `json:"ra_key_negotiation"` // Default: true (nil means not set in JSON)
	RaNeedToken      *bool    `json:"ra_need_token"`      // Default: true (nil means not set in JSON)
	RaPolicyId       string   `json:"ra_policy_id"`       // Default: empty string
	BytedanceTopInfo string   `json:"bytedance_top_info"` // Default: empty string
	AttestInterval   *float64 `json:"attest_interval"`    // Default: nil (no periodic attestation)
}

// ClientConfigFromJson parses a JSON string into a ClientConfig.
func ClientConfigFromJson(jsonStr string) (ClientConfig, error) {
	return ClientConfigFromReader(strings.NewReader(jsonStr))
}

// ClientConfigFromReader parses a JSON stream from an io.Reader into a ClientConfig.
// Applies default values for missing fields.
func ClientConfigFromReader(reader io.Reader) (ClientConfig, error) {
	var config ClientConfig
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&config); err != nil {
		return ClientConfig{}, err
	}

	// Apply default values for missing fields
	config.setDefaults()
	return config, nil
}

// ClientConfigFromFile reads a file and parses its contents into a ClientConfig.
func ClientConfigFromFile(path string) (ClientConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return ClientConfig{}, err
	}
	defer file.Close()

	return ClientConfigFromReader(file)
}

// parseTopInfo deserializes and processes the `bytedanceTopInfo` JSON string.
// Handles URL fallback to `raUrl` and scheme defaulting to https.
func (c *ClientConfig) parseTopInfo() (TopInfo, error) {
	// Deserialize TopInfo from bytedanceTopInfo JSON string
	topInfo, err := deserializeTopInfo(strings.NewReader(c.BytedanceTopInfo))
	if err != nil {
		return TopInfo{}, err
	}

	// Fallback to raUrl if TopInfo URL is empty (backward compatibility)
	if topInfo.URL == "" {
		topInfo.URL = c.RaURL
	}

	// Ensure URL has a scheme (default to https)
	if !strings.HasPrefix(topInfo.URL, "http:") && !strings.HasPrefix(topInfo.URL, "https:") {
		topInfo.URL = "https://" + topInfo.URL
	}

	return topInfo, nil
}

// setDefaults applies default values to missing configuration fields.
func (c *ClientConfig) setDefaults() {
	// Default RA type to TCA if not specified
	if c.RaType == "" {
		c.RaType = RA_TYPE_TCA
	}

	// Default key negotiation to true if not specified
	if c.RaKeyNegotiation == nil {
		trueVal := true
		c.RaKeyNegotiation = &trueVal
	}

	// Default need token to true if not specified
	if c.RaNeedToken == nil {
		trueVal := true
		c.RaNeedToken = &trueVal
	}
}

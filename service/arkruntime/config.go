package arkruntime

import (
	"net/http"
	"strings"
	"time"
)

const (
	defaultBaseUrl                 = "https://ark.cn-beijing.volces.com/api/v3"
	defaultRegion                  = "cn-beijing"
	defaultEmptyMessagesLimit uint = 300
)

type tokenInfo struct {
	token       string
	expiredTime int64
}

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	apiKey string
	ak     string
	sk     string
	region string

	BaseURL            string
	HTTPClient         *http.Client
	EmptyMessagesLimit uint
	RetryTimes         int
}

func NewClientConfig(apiKey, ak, sk string, setters ...configOption) ClientConfig {
	config := ClientConfig{
		apiKey:             apiKey,
		ak:                 ak,
		sk:                 sk,
		region:             defaultRegion,
		BaseURL:            defaultBaseUrl,
		HTTPClient:         &http.Client{},
		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}

	for _, setter := range setters {
		setter(&config)
	}
	return config
}

type configOption func(*ClientConfig)

func WithRegion(region string) configOption {
	return func(config *ClientConfig) {
		config.region = region
	}
}

func WithBaseUrl(url string) configOption {
	return func(config *ClientConfig) {
		config.BaseURL = url
		if strings.HasSuffix(url, "/") {
			config.BaseURL = strings.TrimSuffix(url, "/")
		}
	}
}

func WithRetryTimes(retryTimes int) configOption {
	return func(config *ClientConfig) {
		config.RetryTimes = retryTimes
	}
}

func WithTimeout(timeout time.Duration) configOption {
	return func(config *ClientConfig) {
		config.HTTPClient = &http.Client{
			Timeout: timeout,
		}
	}
}

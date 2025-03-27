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
	defaultRetryTimes         int  = 2
	defaultTimeout                 = 10 * time.Minute
	defaultBatchMaxParallel        = 3000
)

const (
	resourceTypeEndpoint = "endpoint"
	resourceTypeBot      = "bot"
	stsTokenKeyPattern   = "%s#%s"
)

type tokenInfo struct {
	token       string
	expiredTime int64
}

// ClientConfig is a configuration of a client.
type clientConfig struct {
	apiKey string
	ak     string
	sk     string
	region string

	BaseURL            string
	HTTPClient         *http.Client
	EmptyMessagesLimit uint
	RetryTimes         int

	batchMaxParallel int
}

func NewClientConfig(apiKey, ak, sk string, setters ...ConfigOption) clientConfig {
	config := clientConfig{
		apiKey:  apiKey,
		ak:      ak,
		sk:      sk,
		region:  defaultRegion,
		BaseURL: defaultBaseUrl,
		HTTPClient: &http.Client{
			Timeout: defaultTimeout,
		},
		EmptyMessagesLimit: defaultEmptyMessagesLimit,
		RetryTimes:         defaultRetryTimes,
		batchMaxParallel:   defaultBatchMaxParallel,
	}

	for _, setter := range setters {
		setter(&config)
	}
	return config
}

type ConfigOption func(*clientConfig)

func WithRegion(region string) ConfigOption {
	return func(config *clientConfig) {
		config.region = region
	}
}

func WithBaseUrl(url string) ConfigOption {
	return func(config *clientConfig) {
		config.BaseURL = url
		if strings.HasSuffix(url, "/") {
			config.BaseURL = strings.TrimSuffix(url, "/")
		}
	}
}

func WithRetryTimes(retryTimes int) ConfigOption {
	return func(config *clientConfig) {
		config.RetryTimes = retryTimes
	}
}

func WithTimeout(timeout time.Duration) ConfigOption {
	return func(config *clientConfig) {
		config.HTTPClient.Timeout = timeout
	}
}

func WithHTTPClient(client *http.Client) ConfigOption {
	return func(config *clientConfig) {
		config.HTTPClient = client
	}
}

// WithBatchMaxParallel sets the max parallel for batch requests.
func WithBatchMaxParallel(maxParallel int) ConfigOption {
	return func(config *clientConfig) {
		config.batchMaxParallel = maxParallel
	}
}

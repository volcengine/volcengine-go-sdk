package custom

import (
	"context"
	"net/http"
	"net/url"

	"github.com/volcengine/volcstack-go-sdk/volcstack/credentials"
)

type RequestMetadata struct {
	ServiceName string
	Version     string
	Action      string
	HttpMethod  string
	Region      string
	Request     *http.Request
	RawQuery    *url.Values
}

type ExtendHttpRequest func(ctx context.Context, request *http.Request)

type ExtendHttpRequestWithMeta func(ctx context.Context, request *http.Request, meta RequestMetadata)

type ExtraHttpParameters func(ctx context.Context) map[string]string

type ExtraHttpParametersWithMeta func(ctx context.Context, meta RequestMetadata) map[string]string

type ExtraHttpJsonBody func(ctx context.Context, input *map[string]interface{}, meta RequestMetadata)

type LogAccount func(ctx context.Context) *string

type DynamicCredentials func(ctx context.Context) (*credentials.Credentials, *string)

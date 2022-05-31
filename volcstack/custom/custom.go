package custom

import (
	"context"
	"net/http"

	"github.com/volcengine/volcstack-go-sdk/volcstack/credentials"
)

type RequestMetadata struct {
	ServiceName string
	Version     string
	Action      string
	HttpMethod  string
	Region      string
}

type ExtendHttpRequest func(ctx context.Context, request *http.Request)

type ExtendHttpRequestWithMeta func(ctx context.Context, request *http.Request, meta RequestMetadata)

type ExtraHttpParameters func(ctx context.Context) map[string]string

type ExtraHttpParametersWithMeta func(ctx context.Context, meta RequestMetadata) map[string]string

type LogAccount func(ctx context.Context) *string

type DynamicCredentials func(ctx context.Context) (*credentials.Credentials, *string)

package custom

import (
	"context"
	"net/http"
	"net/url"

	"github.com/volcengine/volcstack-go-sdk/volcstack/client/metadata"
)

type SdkInterceptor struct {
	Context    context.Context
	Request    *http.Request
	Name       string
	Method     string
	ClientInfo metadata.ClientInfo
	URI        string
	Header     http.Header
	URL        *url.URL
	Input      interface{}
	Output     interface{}
	Holders    []interface{}
}

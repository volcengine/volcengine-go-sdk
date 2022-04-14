package custom

import (
	"context"
	"net/http"
	"net/url"
)

type SdkInterceptor struct {
	Context context.Context
	Request *http.Request
	Name    string
	Method  string
	URI     string
	Header  http.Header
	URL     *url.URL
	Input   interface{}
	Output  interface{}
	Holders []interface{}
}

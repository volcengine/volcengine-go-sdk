package request

import "github.com/volcengine/volcstack-go-sdk/volcstack"

// setContext updates the Request to use the passed in context for cancellation.
// Context will also be used for request retry delay.
//
// Creates shallow copy of the http.Request with the WithContext method.
func setRequestContext(r *Request, ctx volcstack.Context) {
	r.context = ctx
	r.HTTPRequest = r.HTTPRequest.WithContext(ctx)
}

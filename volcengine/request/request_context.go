package request

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import "github.com/volcengine/volcengine-go-sdk/volcengine"

// setContext updates the Request to use the passed in context for cancellation.
// Context will also be used for request retry delay.
//
// Creates shallow copy of the http.Request with the WithContext method.
func setRequestContext(r *Request, ctx volcengine.Context) {
	r.context = ctx
	r.HTTPRequest = r.HTTPRequest.WithContext(ctx)
}

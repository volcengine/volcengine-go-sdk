package corehandlers

import "github.com/volcengine/volcstack-go-sdk/volcstack/request"

var CustomerRequestHandler = request.NamedHandler{
	Name: "core.CustomerRequestHandler",
	Fn: func(r *request.Request) {
		if r.Config.ExtendHttpRequest != nil {
			r.Config.ExtendHttpRequest(r.Context(), r.HTTPRequest)
		}
	},
}

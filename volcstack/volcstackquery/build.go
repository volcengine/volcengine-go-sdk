package volcstackquery

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"net/url"
	"strings"

	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackbody"
)

// BuildHandler is a named request handler for building volcstackquery protocol requests
var BuildHandler = request.NamedHandler{Name: "volcstacksdk.volcstackquery.Build", Fn: Build}

// Build builds a request for a Volcstack Query service.
func Build(r *request.Request) {
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.ClientInfo.APIVersion},
	}
	//r.HTTPRequest.Header.Add("Accept", "application/json")
	//method := strings.ToUpper(r.HTTPRequest.Method)

	if r.Config.ExtraUserAgent != nil && *r.Config.ExtraUserAgent != "" {
		if strings.HasPrefix(*r.Config.ExtraUserAgent, "/") {
			request.AddToUserAgent(r, *r.Config.ExtraUserAgent)
		} else {
			request.AddToUserAgent(r, "/"+*r.Config.ExtraUserAgent)
		}

	}
	r.HTTPRequest.Host = r.HTTPRequest.URL.Host
	v := r.HTTPRequest.Header.Get("Content-Type")
	if (strings.ToUpper(r.HTTPRequest.Method) == "PUT" ||
		strings.ToUpper(r.HTTPRequest.Method) == "POST" ||
		strings.ToUpper(r.HTTPRequest.Method) == "DELETE" ||
		strings.ToUpper(r.HTTPRequest.Method) == "PATCH") &&
		strings.Contains(strings.ToLower(v), "application/json") {
		r.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
		volcstackbody.BodyJson(&body, r)
	} else {
		volcstackbody.BodyParam(&body, r)
	}
}

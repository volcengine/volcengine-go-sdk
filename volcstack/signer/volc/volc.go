package volc

import (
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/request"
)

var SignRequestHandler = request.NamedHandler{
	Name: "volc.SignRequestHandler", Fn: SignSDKRequest,
}

func SignSDKRequest(req *request.Request) {

	region := req.ClientInfo.SigningRegion
	if region == "" {
		region = volcstack.StringValue(req.Config.Region)
	}

	name := req.ClientInfo.SigningName
	if name == "" {
		name = req.ClientInfo.ServiceID
	}

	credentials := req.Config.Credentials.GetBase(region, name)

	r := credentials.Sign(req.HTTPRequest)
	req.HTTPRequest.Header = r.Header
}

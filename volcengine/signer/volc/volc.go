package volc

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

var SignRequestHandler = request.NamedHandler{
	Name: "volc.SignRequestHandler", Fn: SignSDKRequest,
}

func SignSDKRequest(req *request.Request) {

	region := req.ClientInfo.SigningRegion

	var (
		_credentials *credentials.Credentials
		_region      *string
		c1           base.Credentials
	)

	if req.Config.DynamicCredentials != nil {
		_credentials, _region = req.Config.DynamicCredentials(req.Context())
		if volcengine.StringValue(_region) == "" {
			req.Error = volcengine.ErrMissingRegion
		}
		region = volcengine.StringValue(_region)
	} else if region == "" {
		region = volcengine.StringValue(req.Config.Region)
	}

	name := req.ClientInfo.SigningName
	if name == "" {
		name = req.ClientInfo.ServiceID
	}

	if _credentials == nil {
		c1 = req.Config.Credentials.GetBase(region, name)
	} else {
		c1 = _credentials.GetBase(region, name)
	}
	r := c1.Sign(req.HTTPRequest)
	req.HTTPRequest.Header = r.Header
}

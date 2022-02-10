package hmac

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
)

var SignRequestHandler = request.NamedHandler{
	Name: "hmac.SignRequestHandler", Fn: SignSDKRequest,
}

func SignSDKRequest(req *request.Request) {

}

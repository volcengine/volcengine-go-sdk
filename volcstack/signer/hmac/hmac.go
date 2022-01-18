package hmac

import (
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/request"
)

var SignRequestHandler = request.NamedHandler{
	Name: "hmac.SignRequestHandler", Fn: SignSDKRequest,
}

func SignSDKRequest(req *request.Request) {

}

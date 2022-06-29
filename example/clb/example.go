package main

import (
	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

func main() {
	ak := "access key"
	sk := "secret key"
	var region = "cn-north-3"

	//if use env Credentials
	//please
	//export VOLCSTACK_ACCESS_KEY=AK
	//export VOLCSTACK_SECRET_KEY=SK
	// and WithCredentials(credentials.NewEnvCredentials())

	// Init client
	config := volcengine.NewConfig().
		WithRegion(region).
		WithAkSk(ak, sk).
		//WithCredentials(credentials.NewEnvCredentials()).
		WithDisableSSL(true).
		//WithLogLevel(volcengine.LogDebugWithHTTPBody).
		WithEndpoint(volcengineutil.NewEndpoint().GetEndpoint())
	sess, _ := session.NewSession(config)
	svc := clb.New(sess)

	// Call OpenAPI
	CreateLoadBalancer(svc)
}

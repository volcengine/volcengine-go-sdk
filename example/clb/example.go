package main

import (
	"github.com/volcengine/volcstack-go-sdk/service/clb"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/session"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
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
	config := volcstack.NewConfig().
		WithRegion(region).
		WithAkSk(ak, sk).
		//WithCredentials(credentials.NewEnvCredentials()).
		WithDisableSSL(true).
		//WithLogLevel(volcstack.LogDebugWithHTTPBody).
		WithEndpoint(volcstackutil.NewEndpoint().GetEndpoint())
	sess, _ := session.NewSession(config)
	svc := clb.New(sess)

	// Call OpenAPI
	CreateLoadBalancer(svc)
}

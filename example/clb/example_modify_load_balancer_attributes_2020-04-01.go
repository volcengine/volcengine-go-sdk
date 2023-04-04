package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyLoadBalancerAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	modifyLoadBalancerAttributesInput := &clb.ModifyLoadBalancerAttributesInput{
		Description: volcengine.String("test"),
		LoadBalancerId: volcengine.String("clb-bp1b6c719dfa08ex****"),
		LoadBalancerName: volcengine.String("clb-test1"),
		LoadBalancerSpec: volcengine.String("small_1"),
		ModificationProtectionReason: volcengine.String("实例托管"),
		ModificationProtectionStatus: volcengine.String("ConsoleProtection"),
	}

	resp, err := svc.ModifyLoadBalancerAttributes(modifyLoadBalancerAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

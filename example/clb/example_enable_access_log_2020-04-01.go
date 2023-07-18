// Example Code generated by Beijing Volcanoengine Technology.
package clbexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func EnableAccessLog() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	enableAccessLogInput := &clb.EnableAccessLogInput{
		BucketName:     volcengine.String("clb-tos1"),
		LoadBalancerId: volcengine.String("clb-bp1b6c719dfa08ex****"),
	}

	resp, err := svc.EnableAccessLog(enableAccessLogInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

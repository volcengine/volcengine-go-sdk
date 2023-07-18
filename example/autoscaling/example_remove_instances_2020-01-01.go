// Example Code generated by Beijing Volcanoengine Technology.
package autoscalingexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/autoscaling"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func RemoveInstances() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := autoscaling.New(sess)
	removeInstancesInput := &autoscaling.RemoveInstancesInput{
		InstanceIds:    volcengine.StringSlice([]string{"i-ybmike5l70l8j1ha****"}),
		ScalingGroupId: volcengine.String("scg-ybmssdnnhn5pkgyd****"),
	}

	resp, err := svc.RemoveInstances(removeInstancesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

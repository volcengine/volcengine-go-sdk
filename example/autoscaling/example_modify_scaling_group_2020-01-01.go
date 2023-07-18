// Example Code generated by Beijing Volcanoengine Technology.
package autoscalingexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/autoscaling"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyScalingGroup() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := autoscaling.New(sess)
	modifyScalingGroupInput := &autoscaling.ModifyScalingGroupInput{
		ScalingGroupId: volcengine.String("scg-yblhryq64wgh9zmp****ActiveScalingConfigurationIdscc-yblhz5oxzml8j1gv****"),
	}

	resp, err := svc.ModifyScalingGroup(modifyScalingGroupInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

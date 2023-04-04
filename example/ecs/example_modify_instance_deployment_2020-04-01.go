package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ModifyInstanceDeployment() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	modifyInstanceDeploymentInput := &ecs.ModifyInstanceDeploymentInput{
		DeploymentSetId: volcengine.String("dps-yc1o9aahks5m57nk****"),
		InstanceId: volcengine.String("i-3tigy72q3u3vj0x2****"),
	}

	resp, err := svc.ModifyInstanceDeployment(modifyInstanceDeploymentInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

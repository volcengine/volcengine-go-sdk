package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DeleteDeploymentSet() {
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
	deleteDeploymentSetInput := &ecs.DeleteDeploymentSetInput{
		DeploymentSetId: volcengine.String("dps-yc1o9aahks5m57nk****"),
	}

	resp, err := svc.DeleteDeploymentSet(deleteDeploymentSetInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

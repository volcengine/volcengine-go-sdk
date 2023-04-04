package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ModifyDeploymentSetAttribute() {
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
	modifyDeploymentSetAttributeInput := &ecs.ModifyDeploymentSetAttributeInput{
		DeploymentSetId: volcengine.String("dps-yc1o9aahks5m57nk****"),
		DeploymentSetName: volcengine.String("testDeploymentSet"),
	}

	resp, err := svc.ModifyDeploymentSetAttribute(modifyDeploymentSetAttributeInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

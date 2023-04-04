package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func CreateDeploymentSet() {
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
	createDeploymentSetInput := &ecs.CreateDeploymentSetInput{
		DeploymentSetName: volcengine.String("testDeploymentSetName"),
		Granularity: volcengine.String("host"),
		Strategy: volcengine.String("Availability"),
	}

	resp, err := svc.CreateDeploymentSet(createDeploymentSetInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DisassociateInstancesIamRole() {
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
	disassociateInstancesIamRoleInput := &ecs.DisassociateInstancesIamRoleInput{
		IamRoleName: volcengine.String("EcsTestRole"),
		InstanceIds: volcengine.StringSlice([]string{"i-3tiefmkskq3vj0******"}),
	}

	resp, err := svc.DisassociateInstancesIamRole(disassociateInstancesIamRoleInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

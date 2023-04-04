package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func AssociateInstancesIamRole() {
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
	associateInstancesIamRoleInput := &ecs.AssociateInstancesIamRoleInput{
		IamRoleName: volcengine.String("EcsTestRole"),
		InstanceIds: volcengine.StringSlice([]string{"i-3tiefmkskq3vj0******"}),
	}

	resp, err := svc.AssociateInstancesIamRole(associateInstancesIamRoleInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

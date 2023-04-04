package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func RunInstances() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	reqNetworkInterfaces0 := &ecs.NetworkInterfaceForRunInstancesInput{
		SecurityGroupIds: volcengine.StringSlice([]string{"sg-3ti78x9h8t4bw*****"}),
		SubnetId: volcengine.String("subnet-3tispp1nai4e8i****"),
	}
	reqVolumes0 := &ecs.VolumeForRunInstancesInput{
		Size: volcengine.Int32(40),
		VolumeType: volcengine.String("ESSD_PL0"),
	}
	runInstancesInput := &ecs.RunInstancesInput{
		Count: volcengine.Int32(1),
		ImageId: volcengine.String("image-3tefr6wgx63vj0******"),
		InstanceName: volcengine.String("instance-test"),
		InstanceTypeId: volcengine.String("ecs.g1ie.xlarge"),
		NetworkInterfaces: []*ecs.NetworkInterfaceForRunInstancesInput{reqNetworkInterfaces0},
		Password: volcengine.String("password@123"),
		Volumes: []*ecs.VolumeForRunInstancesInput{reqVolumes0},
		ZoneId: volcengine.String("cn-beijing-a"),
	}

	resp, err := svc.RunInstances(runInstancesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

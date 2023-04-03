package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	describeNetworkInterfacesInput := &vpc.DescribeNetworkInterfacesInput{
		VpcId: volcengine.String("vpc-bp15zckdt37pq72zv****"),
	}

	resp, err := svc.DescribeNetworkInterfaces(describeNetworkInterfacesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
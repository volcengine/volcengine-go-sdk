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
	describeEipAddressAttributesInput := &vpc.DescribeEipAddressAttributesInput{
		AllocationId: volcengine.String("eip-2ze7ujxscd****"),
	}

	resp, err := svc.DescribeEipAddressAttributes(describeEipAddressAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
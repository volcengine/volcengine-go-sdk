package main

import (
	"fmt"
	"github.com/volcengine/volcstack-go-sdk/service/clb"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
)

func EnableAccessLog(svc *clb.CLB) {
	input := &clb.EnableAccessLogInput{
		LoadBalancerId: volcstack.String("clb-xxx"),
		BucketName:     volcstack.String("clb-tos1"),
	}

	resp, err := svc.EnableAccessLog(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DisableAccessLog(svc *clb.CLB) {
	input := &clb.DisableAccessLogInput{
		LoadBalancerId: volcstack.String("clb-xxx"),
	}

	resp, err := svc.DisableAccessLog(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

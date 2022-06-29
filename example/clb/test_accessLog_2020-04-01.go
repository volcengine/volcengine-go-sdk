package main

import (
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func EnableAccessLog(svc *clb.CLB) {
	input := &clb.EnableAccessLogInput{
		LoadBalancerId: volcengine.String("clb-xxx"),
		BucketName:     volcengine.String("clb-tos1"),
	}

	resp, err := svc.EnableAccessLog(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DisableAccessLog(svc *clb.CLB) {
	input := &clb.DisableAccessLogInput{
		LoadBalancerId: volcengine.String("clb-xxx"),
	}

	resp, err := svc.DisableAccessLog(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

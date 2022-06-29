package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func CreateLoadBalancer(svc *clb.CLB) {
	input := &clb.CreateLoadBalancerInput{
		Type:             volcengine.String("private"),
		SubnetId:         volcengine.String("subnet-xxx"),
		LoadBalancerSpec: volcengine.String("small_1"),
		RegionId:         volcengine.String("cn-north-3"),
	}

	resp, err := svc.CreateLoadBalancer(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func ModifyLoadBalancerAttributes(svc *clb.CLB) {
	input := &clb.ModifyLoadBalancerAttributesInput{
		LoadBalancerId:   volcengine.String("clb-xxx"),
		LoadBalancerName: volcengine.String("clb-name"),
	}

	resp, err := svc.ModifyLoadBalancerAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeLoadBalancers(svc *clb.CLB) {
	input := &clb.DescribeLoadBalancersInput{
		PageSize:   volcengine.Int64(100),
		PageNumber: volcengine.Int64(1),
	}

	resp, err := svc.DescribeLoadBalancers(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeLoadBalancerAttributes(svc *clb.CLB) {
	input := &clb.DescribeLoadBalancerAttributesInput{
		LoadBalancerId: volcengine.String("clb-xxx"),
	}

	resp, err := svc.DescribeLoadBalancerAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DeleteLoadBalancer(svc *clb.CLB) {
	input := &clb.DeleteLoadBalancerInput{
		LoadBalancerId: volcengine.String("clb-xxx"),
	}

	resp, err := svc.DeleteLoadBalancer(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}
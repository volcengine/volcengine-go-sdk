package main

import (
	"fmt"

	"github.com/volcengine/volcstack-go-sdk/service/clb"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
)

func CreateLoadBalancer(svc *clb.CLB) {
	input := &clb.CreateLoadBalancerInput{
		Type:             volcstack.String("private"),
		SubnetId:         volcstack.String("subnet-xxx"),
		LoadBalancerSpec: volcstack.String("small_1"),
		RegionId:         volcstack.String("cn-north-3"),
	}

	resp, err := svc.CreateLoadBalancer(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func ModifyLoadBalancerAttributes(svc *clb.CLB) {
	input := &clb.ModifyLoadBalancerAttributesInput{
		LoadBalancerId:   volcstack.String("clb-xxx"),
		LoadBalancerName: volcstack.String("clb-name"),
	}

	resp, err := svc.ModifyLoadBalancerAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeLoadBalancers(svc *clb.CLB) {
	input := &clb.DescribeLoadBalancersInput{
		PageSize:   volcstack.Int64(100),
		PageNumber: volcstack.Int64(1),
	}

	resp, err := svc.DescribeLoadBalancers(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeLoadBalancerAttributes(svc *clb.CLB) {
	input := &clb.DescribeLoadBalancerAttributesInput{
		LoadBalancerId: volcstack.String("clb-xxx"),
	}

	resp, err := svc.DescribeLoadBalancerAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DeleteLoadBalancer(svc *clb.CLB) {
	input := &clb.DeleteLoadBalancerInput{
		LoadBalancerId: volcstack.String("clb-xxx"),
	}

	resp, err := svc.DeleteLoadBalancer(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}
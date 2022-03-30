package main

import (
	"fmt"
	"github.com/volcengine/volcstack-go-sdk/service/clb"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
)

func CreateServerGroup(svc *clb.CLB) {
	input := &clb.CreateServerGroupInput{
		LoadBalancerId: volcstack.String("clb-xxx"),
	}

	resp, err := svc.CreateServerGroup(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func ModifyServerGroupAttributes(svc *clb.CLB) {
	input := &clb.ModifyServerGroupAttributesInput{
		ServerGroupId: volcstack.String("rsp-xxx"),
	}

	resp, err := svc.ModifyServerGroupAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeServerGroups(svc *clb.CLB) {
	input := &clb.DescribeServerGroupsInput{
		PageSize:   volcstack.Int64(100),
		PageNumber: volcstack.Int64(1),
	}

	resp, err := svc.DescribeServerGroups(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeServerGroupAttributes(svc *clb.CLB) {
	input := &clb.DescribeServerGroupAttributesInput{
		ServerGroupId: volcstack.String("rsp-xxx"),
	}

	resp, err := svc.DescribeServerGroupAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func AddServerGroupBackendServers(svc *clb.CLB) {
	input := &clb.AddServerGroupBackendServersInput{
		ServerGroupId: volcstack.String("rsp-xxx"),
		Servers: []*clb.ServerForAddServerGroupBackendServersInput{
			{
				InstanceId:  volcstack.String("i-xxx"),
				Type:        volcstack.String("ecs"),
				Weight:      volcstack.Int64(100),
				Ip:          volcstack.String("192.168.1.1"),
				Port:        volcstack.Int64(80),
				Description: volcstack.String("ecs description"),
			},
		},
	}

	resp, err := svc.AddServerGroupBackendServers(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func RemoveServerGroupBackendServers(svc *clb.CLB) {
	input := &clb.RemoveServerGroupBackendServersInput{
		ServerGroupId: volcstack.String("rsp-xxx"),
		ServerIds: []*string{
			volcstack.String("rsp-xxx"),
			volcstack.String("rsp-xxx"),
		},
	}

	resp, err := svc.RemoveServerGroupBackendServers(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DeleteServerGroup(svc *clb.CLB) {
	input := &clb.DeleteServerGroupInput{
		ServerGroupId: volcstack.String("rsp-xxx"),
	}

	resp, err := svc.DeleteServerGroup(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

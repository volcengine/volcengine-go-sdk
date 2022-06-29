package main

import (
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func CreateServerGroup(svc *clb.CLB) {
	input := &clb.CreateServerGroupInput{
		LoadBalancerId: volcengine.String("clb-xxx"),
	}

	resp, err := svc.CreateServerGroup(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func ModifyServerGroupAttributes(svc *clb.CLB) {
	input := &clb.ModifyServerGroupAttributesInput{
		ServerGroupId: volcengine.String("rsp-xxx"),
	}

	resp, err := svc.ModifyServerGroupAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeServerGroups(svc *clb.CLB) {
	input := &clb.DescribeServerGroupsInput{
		PageSize:   volcengine.Int64(100),
		PageNumber: volcengine.Int64(1),
	}

	resp, err := svc.DescribeServerGroups(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeServerGroupAttributes(svc *clb.CLB) {
	input := &clb.DescribeServerGroupAttributesInput{
		ServerGroupId: volcengine.String("rsp-xxx"),
	}

	resp, err := svc.DescribeServerGroupAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func AddServerGroupBackendServers(svc *clb.CLB) {
	input := &clb.AddServerGroupBackendServersInput{
		ServerGroupId: volcengine.String("rsp-xxx"),
		Servers: []*clb.ServerForAddServerGroupBackendServersInput{
			{
				InstanceId:  volcengine.String("i-xxx"),
				Type:        volcengine.String("ecs"),
				Weight:      volcengine.Int64(100),
				Ip:          volcengine.String("192.168.1.1"),
				Port:        volcengine.Int64(80),
				Description: volcengine.String("ecs description"),
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
		ServerGroupId: volcengine.String("rsp-xxx"),
		ServerIds: []*string{
			volcengine.String("rsp-xxx"),
			volcengine.String("rsp-xxx"),
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
		ServerGroupId: volcengine.String("rsp-xxx"),
	}

	resp, err := svc.DeleteServerGroup(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

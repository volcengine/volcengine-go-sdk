package main

import (
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func CreateListener(svc *clb.CLB) {
	input := &clb.CreateListenerInput{
		LoadBalancerId: volcengine.String("clb-xxx"),
		Protocol:       volcengine.String("TCP"),
		Port:           volcengine.Int64(80),
		ServerGroupId:  volcengine.String("rsp-xxx"),
	}

	resp, err := svc.CreateListener(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func ModifyListenerAttributes(svc *clb.CLB) {
	input := &clb.ModifyListenerAttributesInput{
		ListenerId:   volcengine.String("lsn-xxx"),
		ListenerName: volcengine.String("listener-name"),
	}

	resp, err := svc.ModifyListenerAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeListeners(svc *clb.CLB) {
	input := &clb.DescribeListenersInput{
		LoadBalancerId: volcengine.String("clb-xxx"),
	}

	resp, err := svc.DescribeListeners(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeListenerAttributes(svc *clb.CLB) {
	input := &clb.DescribeListenerAttributesInput{
		ListenerId: volcengine.String("lsn-xxx"),
	}

	resp, err := svc.DescribeListenerAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeListenerHealth(svc *clb.CLB) {
	input := &clb.DescribeListenerHealthInput{
		ListenerId: volcengine.String("lsn-xxx"),
	}

	resp, err := svc.DescribeListenerHealth(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DeleteListener(svc *clb.CLB) {
	input := &clb.DeleteListenerInput{
		ListenerId: volcengine.String("lsn-xxx"),
	}

	resp, err := svc.DeleteListener(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

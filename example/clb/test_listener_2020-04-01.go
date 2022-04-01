package main

import (
	"fmt"
	"github.com/volcengine/volcstack-go-sdk/service/clb"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
)

func CreateListener(svc *clb.CLB) {
	input := &clb.CreateListenerInput{
		LoadBalancerId: volcstack.String("clb-xxx"),
		Protocol:       volcstack.String("TCP"),
		Port:           volcstack.Int64(80),
		ServerGroupId:  volcstack.String("rsp-xxx"),
	}

	resp, err := svc.CreateListener(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func ModifyListenerAttributes(svc *clb.CLB) {
	input := &clb.ModifyListenerAttributesInput{
		ListenerId:   volcstack.String("lsn-xxx"),
		ListenerName: volcstack.String("listener-name"),
	}

	resp, err := svc.ModifyListenerAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeListeners(svc *clb.CLB) {
	input := &clb.DescribeListenersInput{
		LoadBalancerId: volcstack.String("clb-xxx"),
	}

	resp, err := svc.DescribeListeners(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeListenerAttributes(svc *clb.CLB) {
	input := &clb.DescribeListenerAttributesInput{
		ListenerId: volcstack.String("lsn-xxx"),
	}

	resp, err := svc.DescribeListenerAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeListenerHealth(svc *clb.CLB) {
	input := &clb.DescribeListenerHealthInput{
		ListenerId: volcstack.String("lsn-xxx"),
	}

	resp, err := svc.DescribeListenerHealth(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DeleteListener(svc *clb.CLB) {
	input := &clb.DeleteListenerInput{
		ListenerId: volcstack.String("lsn-xxx"),
	}

	resp, err := svc.DeleteListener(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

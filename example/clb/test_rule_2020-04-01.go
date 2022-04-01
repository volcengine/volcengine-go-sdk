package main

import (
	"fmt"
	"github.com/volcengine/volcstack-go-sdk/service/clb"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
)

func CreateRules(svc *clb.CLB) {
	input := &clb.CreateRulesInput{
		ListenerId: volcstack.String("lsn-xxx"),
		Rules: []*clb.RuleForCreateRulesInput{
			{
				ServerGroupId: volcstack.String("rsp-xxx"),
				Domain:        volcstack.String("*.com"),
				Description:   volcstack.String("rule description"),
			},
		},
	}

	resp, err := svc.CreateRules(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func ModifyRules(svc *clb.CLB) {
	input := &clb.ModifyRulesInput{
		ListenerId: volcstack.String("rsp-xxx"),
		Rules: []*clb.RuleForModifyRulesInput{
			{
				RuleId:        volcstack.String("rule-xxx"),
				ServerGroupId: volcstack.String("rsp-xxx"),
				Description:   volcstack.String("rule description"),
			},
		},
	}

	resp, err := svc.ModifyRules(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeRules(svc *clb.CLB) {
	input := &clb.DescribeRulesInput{
		ListenerId: volcstack.String("lsn-xxx"),
	}

	resp, err := svc.DescribeRules(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DeleteRules(svc *clb.CLB) {
	input := &clb.DeleteRulesInput{
		ListenerId: volcstack.String("lsn-xxx"),
		RuleIds: []*string{
			volcstack.String("rule-xxx"),
			volcstack.String("rule-xxx"),
		},
	}

	resp, err := svc.DeleteRules(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

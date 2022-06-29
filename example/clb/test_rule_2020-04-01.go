package main

import (
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func CreateRules(svc *clb.CLB) {
	input := &clb.CreateRulesInput{
		ListenerId: volcengine.String("lsn-xxx"),
		Rules: []*clb.RuleForCreateRulesInput{
			{
				ServerGroupId: volcengine.String("rsp-xxx"),
				Domain:        volcengine.String("*.com"),
				Description:   volcengine.String("rule description"),
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
		ListenerId: volcengine.String("rsp-xxx"),
		Rules: []*clb.RuleForModifyRulesInput{
			{
				RuleId:        volcengine.String("rule-xxx"),
				ServerGroupId: volcengine.String("rsp-xxx"),
				Description:   volcengine.String("rule description"),
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
		ListenerId: volcengine.String("lsn-xxx"),
	}

	resp, err := svc.DescribeRules(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DeleteRules(svc *clb.CLB) {
	input := &clb.DeleteRulesInput{
		ListenerId: volcengine.String("lsn-xxx"),
		RuleIds: []*string{
			volcengine.String("rule-xxx"),
			volcengine.String("rule-xxx"),
		},
	}

	resp, err := svc.DeleteRules(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

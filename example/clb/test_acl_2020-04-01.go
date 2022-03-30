package main

import (
	"fmt"
	"github.com/volcengine/volcstack-go-sdk/service/clb"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
)

func CreateAcl(svc *clb.CLB) {
	input := &clb.CreateAclInput{
		AclName:     volcstack.String("acl-name"),
		Description: volcstack.String("acl description"),
	}

	resp, err := svc.CreateAcl(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func ModifyAclAttributes(svc *clb.CLB) {
	input := &clb.ModifyAclAttributesInput{
		AclId:       volcstack.String("acl-xxx"),
		AclName:     volcstack.String("acl-name"),
		Description: volcstack.String("acl description"),
	}

	resp, err := svc.ModifyAclAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeAcls(svc *clb.CLB) {
	input := &clb.DescribeAclsInput{
		PageSize:   volcstack.Int64(100),
		PageNumber: volcstack.Int64(1),
	}

	resp, err := svc.DescribeAcls(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DescribeAclAttributes(svc *clb.CLB) {
	input := &clb.DescribeAclAttributesInput{
		AclId: volcstack.String("acl-xxx"),
	}

	resp, err := svc.DescribeAclAttributes(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func AddAclEntries(svc *clb.CLB) {
	input := &clb.AddAclEntriesInput{
		AclId: volcstack.String("acl-xxx"),
		AclEntries: []*clb.AclEntryForAddAclEntriesInput{
			{
				Entry:       volcstack.String("192.XX.XX.7/16"),
				Description: volcstack.String("acl entry description"),
			},
		},
	}

	resp, err := svc.AddAclEntries(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func RemoveAclEntries(svc *clb.CLB) {
	input := &clb.RemoveAclEntriesInput{
		AclId: volcstack.String("acl-xxx"),
		Entries: []*string{
			volcstack.String("192.XX.XX.7/16"),
		},
	}

	resp, err := svc.RemoveAclEntries(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

func DeleteAcl(svc *clb.CLB) {
	input := &clb.DeleteAclInput{
		AclId: volcstack.String("acl-xxx"),
	}

	resp, err := svc.DeleteAcl(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

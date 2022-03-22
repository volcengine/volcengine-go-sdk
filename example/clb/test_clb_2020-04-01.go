package main

import (
	"fmt"
	"github.com/volcengine/volcstack-go-sdk/service/clb"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/session"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

func main() {
	ak := "access key"
	sk := "secret key"
	var region = "cn-north-3"

	//if use env Credentials
	//please
	//export VOLCSTACK_ACCESS_KEY=AK
	//export VOLCSTACK_SECRET_KEY=SK
	// and WithCredentials(credentials.NewEnvCredentials())

	config := volcstack.NewConfig().
		WithRegion(region).
		WithAkSk(ak, sk).
		//WithCredentials(credentials.NewEnvCredentials()).
		WithDisableSSL(true).
		//WithLogLevel(volcstack.LogDebugWithHTTPBody).
		WithEndpoint(volcstackutil.NewEndpoint().GetEndpoint())
	sess, _ := session.NewSession(config)
	svc := clb.New(sess)

	// CreateLoadBalancer
	input := &clb.CreateLoadBalancerInput{}

	input.Type = volcstack.String("private")
	input.SubnetId = volcstack.String("subnet-xxx")
	input.LoadBalancerSpec = volcstack.String("small_1")
	input.RegionId = volcstack.String(region)

	resp, err := svc.CreateLoadBalancer(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)

	// ModifyLoadBalancerAttributes
	//input := &clb.ModifyLoadBalancerAttributesInput{}
	//
	//input.LoadBalancerId = volcstack.String("clb-xxx")
	//input.LoadBalancerName = volcstack.String("clb-name")
	//
	//resp, err := svc.ModifyLoadBalancerAttributes(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DescribeLoadBalancers
	//input := &clb.DescribeLoadBalancersInput{}
	//
	//input.PageSize = volcstack.Int64(100)
	//input.PageNumber = volcstack.Int64(1)
	//
	//resp, err := svc.DescribeLoadBalancers(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DescribeLoadBalancerAttributes
	//input := &clb.DescribeLoadBalancerAttributesInput{}
	//
	//input.LoadBalancerId = volcstack.String("clb-xxx")
	//
	//resp, err := svc.DescribeLoadBalancerAttributes(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DeleteLoadBalancer
	//input := &clb.DeleteLoadBalancerInput{}
	//
	//input.LoadBalancerId = volcstack.String("clb-xxx")
	//
	//resp, err := svc.DeleteLoadBalancer(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ConvertLoadBalancerBillingType

	// RenewLoadBalancer

	// SetLoadBalancerRenewal

	// DescribeLoadBalancersBilling

	// CreateListener
	//input := &clb.CreateListenerInput{}
	//
	//input.LoadBalancerId = volcstack.String("clb-xxx")
	//input.Protocol = volcstack.String("TCP")
	//input.Port = volcstack.Int64(80)
	//input.ServerGroupId = volcstack.String("rsp-xxx")
	//
	//resp, err := svc.CreateListener(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ModifyListenerAttributes
	//input := &clb.ModifyListenerAttributesInput{}
	//
	//input.ListenerId = volcstack.String("lsn-xxx")
	//input.ListenerName = volcstack.String("listener-name")
	//
	//resp, err := svc.ModifyListenerAttributes(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DescribeListeners
	//input := &clb.DescribeListenersInput{}
	//
	//input.LoadBalancerId = volcstack.String("clb-xxx")
	//
	//resp, err := svc.DescribeListeners(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DescribeListenerAttributes
	//input := &clb.DescribeListenerAttributesInput{}
	//
	//input.ListenerId = volcstack.String("lsn-xxx")
	//
	//resp, err := svc.DescribeListenerAttributes(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DescribeListenerHealth
	//input := &clb.DescribeListenerHealthInput{}
	//
	//input.ListenerId = volcstack.String("lsn-xxx")
	//
	//resp, err := svc.DescribeListenerHealth(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DeleteListener
	//input := &clb.DeleteListenerInput{}
	//
	//input.ListenerId = volcstack.String("lsn-xxx")
	//
	//resp, err := svc.DeleteListener(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// CreateServerGroup
	//input := &clb.CreateServerGroupInput{}
	//
	//input.LoadBalancerId = volcstack.String("clb-xxx")
	//
	//resp, err := svc.CreateServerGroup(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ModifyServerGroupAttributes
	//input := &clb.ModifyServerGroupAttributesInput{}
	//
	//input.ServerGroupId = volcstack.String("rsp-xxx")
	//
	//resp, err := svc.ModifyServerGroupAttributes(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DescribeServerGroups
	//input := &clb.DescribeServerGroupsInput{}
	//
	//input.PageSize = volcstack.Int64(100)
	//input.PageNumber = volcstack.Int64(1)
	//
	//resp, err := svc.DescribeServerGroups(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DescribeServerGroupAttributes
	//input := &clb.DescribeServerGroupAttributesInput{}
	//
	//input.ServerGroupId = volcstack.String("rsp-xxx")
	//
	//resp, err := svc.DescribeServerGroupAttributes(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// AddServerGroupBackendServers
	//input := &clb.AddServerGroupBackendServersInput{}
	//
	//input.ServerGroupId = volcstack.String("rsp-xxx")
	//input.Servers = []*clb.ServerForAddServerGroupBackendServersInput{
	//	{
	//		InstanceId:  volcstack.String("i-xxx"),
	//		Type:        volcstack.String("ecs"),
	//		Weight:      volcstack.Int64(100),
	//		Ip:          volcstack.String("192.168.1.1"),
	//		Port:        volcstack.Int64(80),
	//		Description: volcstack.String("ecs description"),
	//	},
	//}
	//
	//resp, err := svc.AddServerGroupBackendServers(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RemoveServerGroupBackendServers
	//input := &clb.RemoveServerGroupBackendServersInput{}
	//
	//input.ServerGroupId = volcstack.String("rsp-xxx")
	//input.ServerIds = []*string{
	//	volcstack.String("rsp-xxx"),
	//	volcstack.String("rsp-xxx"),
	//}
	//
	//resp, err := svc.RemoveServerGroupBackendServers(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DeleteServerGroup
	//input := &clb.DeleteServerGroupInput{}
	//
	//input.ServerGroupId = volcstack.String("rsp-xxx")
	//
	//resp, err := svc.DeleteServerGroup(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// CreateAcl
	//input := &clb.CreateAclInput{}
	//
	//input.AclName = volcstack.String("acl-name")
	//input.Description = volcstack.String("acl description")
	//
	//resp, err := svc.CreateAcl(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ModifyAclAttributes
	//input := &clb.ModifyAclAttributesInput{}
	//
	//input.AclId = volcstack.String("acl-xxx")
	//input.AclName = volcstack.String("acl-name")
	//input.Description = volcstack.String("acl description")
	//
	//resp, err := svc.ModifyAclAttributes(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DescribeAcls
	//input := &clb.DescribeAclsInput{}
	//
	//input.PageSize = volcstack.Int64(100)
	//input.PageNumber = volcstack.Int64(1)
	//
	//resp, err := svc.DescribeAcls(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DescribeAclAttributes
	//input := &clb.DescribeAclAttributesInput{}
	//
	//input.AclId = volcstack.String("acl-xxx")
	//
	//resp, err := svc.DescribeAclAttributes(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// AddAclEntries
	//input := &clb.AddAclEntriesInput{}
	//
	//input.AclId = volcstack.String("acl-xxx")
	//input.AclEntries = []*clb.AclEntryForAddAclEntriesInput{
	//	{
	//		Entry:       volcstack.String("192.XX.XX.7/16"),
	//		Description: volcstack.String("acl entry description"),
	//	},
	//}
	//
	//resp, err := svc.AddAclEntries(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RemoveAclEntries
	//input := &clb.RemoveAclEntriesInput{}
	//
	//input.AclId = volcstack.String("acl-xxx")
	//input.Entries = []*string{
	//	volcstack.String("192.XX.XX.7/16"),
	//}
	//
	//resp, err := svc.RemoveAclEntries(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DeleteAcl
	//input := &clb.DeleteAclInput{}
	//
	//input.AclId = volcstack.String("acl-xxx")
	//
	//resp, err := svc.DeleteAcl(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// CreateRules
	//input := &clb.CreateRulesInput{}
	//
	//input.ListenerId = volcstack.String("lsn-xxx")
	//input.Rules = []*clb.RuleForCreateRulesInput{
	//	{
	//		ServerGroupId: volcstack.String("rsp-xxx"),
	//		Domain:        volcstack.String("*.com"),
	//		Description:   volcstack.String("rule description"),
	//	},
	//}
	//
	//resp, err := svc.CreateRules(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ModifyRules
	//input := &clb.ModifyRulesInput{}

	//input.ListenerId = volcstack.String("rsp-xxx")
	//input.Rules = []*clb.RuleForModifyRulesInput{
	//	{
	//		RuleId:        volcstack.String("rule-xxx"),
	//		ServerGroupId: volcstack.String("rsp-xxx"),
	//		Description:   volcstack.String("rule description"),
	//	},
	//}
	//
	//resp, err := svc.ModifyRules(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DescribeRules
	//input := &clb.DescribeRulesInput{}
	//
	//input.ListenerId = volcstack.String("lsn-xxx")
	//
	//resp, err := svc.DescribeRules(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DeleteRules
	//input := &clb.DeleteRulesInput{}
	//
	//input.ListenerId = volcstack.String("lsn-xxx")
	//input.RuleIds = []*string{
	//	volcstack.String("rule-xxx"),
	//	volcstack.String("rule-xxx"),
	//}
	//
	//resp, err := svc.DeleteRules(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// UploadCertificate
	//input := &clb.UploadCertificateInput{}
	//
	//input.CertificateName = volcstack.String("cert-name")
	//input.Description = volcstack.String("cert description")
	//input.PrivateKey = volcstack.String("xxx")
	//input.PublicKey = volcstack.String("xxx")
	//
	//resp, err := svc.UploadCertificate(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DescribeCertificates
	//input := &clb.DescribeCertificatesInput{}
	//
	//input.PageSize = volcstack.Int64(100)
	//input.PageNumber = volcstack.Int64(1)
	//
	//resp, err := svc.DescribeCertificates(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DeleteCertificate
	//input := &clb.DeleteCertificateInput{}
	//
	//input.CertificateId = volcstack.String("cert-xxx")
	//
	//resp, err := svc.DeleteCertificate(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// EnableAccessLog
	//input := &clb.EnableAccessLogInput{}
	//
	//input.LoadBalancerId = volcstack.String("clb-xxx")
	//input.BucketName = volcstack.String("clb-tos1")
	//
	//resp, err := svc.EnableAccessLog(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// DisableAccessLog
	//input := &clb.DisableAccessLogInput{}
	//
	//input.LoadBalancerId = volcstack.String("clb-xxx")
	//
	//resp, err := svc.DisableAccessLog(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)
}

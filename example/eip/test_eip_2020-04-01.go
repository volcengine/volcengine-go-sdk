package main

import (
	"fmt"

	"github.com/volcengine/volcstack-go-sdk/service/vpc"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/session"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

func main() {
	ak := ""
	sk := ""
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
	svc := vpc.New(sess)

	// eip AllocateEipAddress
	//in := &vpc.AllocateEipAddressInput{
	//	BillingType: volcstack.Int64(2),
	//	Bandwidth:   volcstack.Int64(1),
	//	ISP:         volcstack.String("BGP"),
	//	Name:        volcstack.String("test"),
	//	Description: volcstack.String("test"),
	//	PeriodUnit:  volcstack.Int64(1),
	//	Period:   	 volcstack.Int64(1),
	//}
	//resp, err := svc.AllocateEipAddress(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// eip DescribeEipAddresses
	//in := &vpc.DescribeEipAddressesInput{
	//	Status:      			 volcstack.String("Available"),
	//	EipAddresses: 			 []*string{volcstack.String("10.227.135.195")},
	//	AllocationIds: 			 []*string{volcstack.String("eip-274shqcruf08w7fap8sbtslqu")},
	//	PageNumber:     		 volcstack.Int64(1),
	//	PageSize:     			 volcstack.Int64(20),
	//	Name:        			 volcstack.String("test"),
	//	ISP:         			 volcstack.String("BGP"),
	//	//AssociatedInstanceId:    volcstack.String(""),
	//	//AssociatedInstanceType:  volcstack.String("EcsInstance"),
	//}
	//resp, err := svc.DescribeEipAddresses(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// eip AssociateEipAddress
	//in := &vpc.AssociateEipAddressInput{
	//	AllocationId:     volcstack.String("eip-274shqcruf08w7fap8sbtslqu"),
	//	InstanceId:    	  volcstack.String("i-72q1xwbl4t5zogd59ls9"),
	//	InstanceType:     volcstack.String("EcsInstance"),
	//	PrivateIpAddress: volcstack.String("192.168.0.191"),
	//}
	//resp, err := svc.AssociateEipAddress(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// eip DisassociateEipAddress
	//in := &vpc.DisassociateEipAddressInput{
	//	AllocationId:     volcstack.String("eip-274shqcruf08w7fap8sbtslqu"),
	//	InstanceId:    	  volcstack.String("i-72q1xwbl4t5zogd59ls9"),
	//	InstanceType:     volcstack.String("EcsInstance"),
	//}
	//resp, err := svc.DisassociateEipAddress(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// eip DisassociateEipAddress
	//in := &vpc.ReleaseEipAddressInput{
	//	AllocationId:     volcstack.String("eip-274shqcruf08w7fap8sbtslqu"),
	//}
	//resp, err := svc.ReleaseEipAddress(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// eip ModifyEipAddressAttributes
	//in := &vpc.ModifyEipAddressAttributesInput{
	//	AllocationId:   volcstack.String("eip-274t1gtv5fugw7fap8scozv4s"),
	//	Bandwidth:     	volcstack.Int64(2),
	//	Name:     		volcstack.String("test"),
	//	Description:    volcstack.String("test"),
	//}
	//resp, err := svc.ModifyEipAddressAttributes(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// eip DescribeEipAddressAttributes
	in := &vpc.DescribeEipAddressAttributesInput{
		AllocationId: volcstack.String("eip-274t1gtv5fugw7fap8scozv4s"),
	}
	resp, err := svc.DescribeEipAddressAttributes(in)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v", resp)
}

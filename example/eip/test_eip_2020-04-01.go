package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
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
	config := volcengine.NewConfig().
		WithRegion(region).
		WithAkSk(ak, sk).
		//WithCredentials(credentials.NewEnvCredentials()).
		WithDisableSSL(true).
		//WithLogLevel(volcengine.LogDebugWithHTTPBody).
		WithEndpoint(volcengineutil.NewEndpoint().GetEndpoint())
	sess, _ := session.NewSession(config)
	svc := vpc.New(sess)

	// eip AllocateEipAddress
	//in := &vpc.AllocateEipAddressInput{
	//	BillingType: volcengine.Int64(2),
	//	Bandwidth:   volcengine.Int64(1),
	//	ISP:         volcengine.String("BGP"),
	//	Name:        volcengine.String("test"),
	//	Description: volcengine.String("test"),
	//	PeriodUnit:  volcengine.Int64(1),
	//	Period:   	 volcengine.Int64(1),
	//}
	//resp, err := svc.AllocateEipAddress(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// eip DescribeEipAddresses
	//in := &vpc.DescribeEipAddressesInput{
	//	Status:      			 volcengine.String("Available"),
	//	EipAddresses: 			 []*string{volcengine.String("10.227.135.195")},
	//	AllocationIds: 			 []*string{volcengine.String("eip-274shqcruf08w7fap8sbtslqu")},
	//	PageNumber:     		 volcengine.Int64(1),
	//	PageSize:     			 volcengine.Int64(20),
	//	Name:        			 volcengine.String("test"),
	//	ISP:         			 volcengine.String("BGP"),
	//	//AssociatedInstanceId:    volcengine.String(""),
	//	//AssociatedInstanceType:  volcengine.String("EcsInstance"),
	//}
	//resp, err := svc.DescribeEipAddresses(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// eip AssociateEipAddress
	//in := &vpc.AssociateEipAddressInput{
	//	AllocationId:     volcengine.String("eip-274shqcruf08w7fap8sbtslqu"),
	//	InstanceId:    	  volcengine.String("i-72q1xwbl4t5zogd59ls9"),
	//	InstanceType:     volcengine.String("EcsInstance"),
	//	PrivateIpAddress: volcengine.String("192.168.0.191"),
	//}
	//resp, err := svc.AssociateEipAddress(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// eip DisassociateEipAddress
	//in := &vpc.DisassociateEipAddressInput{
	//	AllocationId:     volcengine.String("eip-274shqcruf08w7fap8sbtslqu"),
	//	InstanceId:    	  volcengine.String("i-72q1xwbl4t5zogd59ls9"),
	//	InstanceType:     volcengine.String("EcsInstance"),
	//}
	//resp, err := svc.DisassociateEipAddress(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// eip DisassociateEipAddress
	//in := &vpc.ReleaseEipAddressInput{
	//	AllocationId:     volcengine.String("eip-274shqcruf08w7fap8sbtslqu"),
	//}
	//resp, err := svc.ReleaseEipAddress(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// eip ModifyEipAddressAttributes
	//in := &vpc.ModifyEipAddressAttributesInput{
	//	AllocationId:   volcengine.String("eip-274t1gtv5fugw7fap8scozv4s"),
	//	Bandwidth:     	volcengine.Int64(2),
	//	Name:     		volcengine.String("test"),
	//	Description:    volcengine.String("test"),
	//}
	//resp, err := svc.ModifyEipAddressAttributes(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// eip DescribeEipAddressAttributes
	in := &vpc.DescribeEipAddressAttributesInput{
		AllocationId: volcengine.String("eip-274t1gtv5fugw7fap8scozv4s"),
	}
	resp, err := svc.DescribeEipAddressAttributes(in)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v", resp)
}

package main

import (
	"fmt"

	"github.com/volcengine/volcstack-go-sdk/service/vpc"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/session"
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
		WithDisableSSL(true)
	//WithLogLevel(volcstack.LogDebugWithHTTPBody).
	//WithEndpoint(volcstackutil.NewEndpoint().GetEndpoint())
	sess, _ := session.NewSession(config)
	svc := vpc.New(sess)

	// ENI CreateNetworkInterface
	//query := &vpc.CreateNetworkInterfaceInput{
	//	SubnetId:     			volcstack.String("subnet-274xwulapzc3k7fap8thue3pl"),
	//	SecurityGroupIds: 		[]*string{volcstack.String("sg-2744odkqfhkhs7fap8tnp49r6")},
	//	PrimaryIpAddress: 		volcstack.String("192.168.1.2"),
	//	NetworkInterfaceName: 	volcstack.String("test"),
	//	Description: 			volcstack.String("test"),
	//	PortSecurityEnabled: 	volcstack.Bool(false),
	//}
	//resp, err := svc.CreateNetworkInterface(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI ModifyNetworkInterfaceAttributes
	//query := &vpc.ModifyNetworkInterfaceAttributesInput{
	//	NetworkInterfaceId:     volcstack.String("eni-274t9vh8ndm9s7fap8shjokny"),
	//	SecurityGroupIds: 		[]*string{volcstack.String("sg-2744odkqfhkhs7fap8tnp49r6"), volcstack.String("sg-274k348d2mhog7fap8thr09gd")},
	//	NetworkInterfaceName: 	volcstack.String("test-up"),
	//	Description: 			volcstack.String("test-up"),
	//}
	//resp, err := svc.ModifyNetworkInterfaceAttributes(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI DescribeNetworkInterfaces
	//query := &vpc.DescribeNetworkInterfacesInput{
	//	Type: 					volcstack.String("secondary"),
	//	Status: 			    volcstack.String("Available"),
	//	VpcId:                  volcstack.String("vpc-2744odetc5ekg7fap8tlh9dl3"),
	//	SubnetId:     			volcstack.String("subnet-274xwulapzc3k7fap8thue3pl"),
	//	//InstanceId:     	    volcstack.String("subnet-274xwulapzc3k7fap8thue3pl"),
	//	PrimaryIpAddresses:     []*string{volcstack.String("192.168.1.2")},
	//	SecurityGroupId:        volcstack.String("sg-2744odkqfhkhs7fap8tnp49r6"),
	//	NetworkInterfaceIds:    []*string{volcstack.String("eni-274t9vh8ndm9s7fap8shjokny")},
	//	NetworkInterfaceName: 	volcstack.String("test-up"),
	//	PageNumber:             volcstack.Int64(1),
	//	PageSize:             	volcstack.Int64(20),
	//	//PrivateIpAddresses:     []*string{volcstack.String("192.168.1.2")},
	//	//ZoneId:                 volcstack.String("cn-lingqiu-a"),
	//}
	//resp, err := svc.DescribeNetworkInterfaces(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI DescribeNetworkInterfaceAttributes
	//query := &vpc.DescribeNetworkInterfaceAttributesInput{
	//	NetworkInterfaceId:    volcstack.String("eni-274ngjt4wp2io7fap8sfxgfzn"),
	//}
	//resp, err := svc.DescribeNetworkInterfaceAttributes(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI DeleteNetworkInterface
	//query := &vpc.DeleteNetworkInterfaceInput{
	//	NetworkInterfaceId:    volcstack.String("eni-274t9vh8ndm9s7fap8shjokny"),
	//}
	//resp, err := svc.DeleteNetworkInterface(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI AttachNetworkInterface
	//query := &vpc.AttachNetworkInterfaceInput{
	//	NetworkInterfaceId:    volcstack.String("eni-2few40gfjx2bk59gp68wu110z"),
	//	InstanceId:    		   volcstack.String("i-ybkbqfzie1l8j1or2ke5"),
	//}
	//resp, err := svc.AttachNetworkInterface(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI DetachNetworkInterface
	//query := &vpc.DetachNetworkInterfaceInput{
	//	NetworkInterfaceId:    volcstack.String("eni-2few40gfjx2bk59gp68wu110z"),
	//	InstanceId:    		   volcstack.String("i-ybkbqfzie1l8j1or2ke5"),
	//}
	//resp, err := svc.DetachNetworkInterface(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// todo 增加AssignPrivateIpAddresses、UnassignPrivateIpAddresses

	// RouteTable CreateRouteTable
	//query := &vpc.CreateRouteTableInput{
	//	VpcId:     			volcstack.String("vpc-2fevwte3frk0059gp67d69r4j"),
	//	RouteTableName: 	volcstack.String("test"),
	//	Description: 	    volcstack.String("test"),
	//}
	//resp, err := svc.CreateRouteTable(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteTable DeleteRouteTable
	//query := &vpc.DeleteRouteTableInput{
	//	RouteTableId:  volcstack.String("vtb-12bhqhvptkl4w17q7y1pv9qmc"),
	//}
	//resp, err := svc.DeleteRouteTable(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteTable ModifyRouteTableAttributes
	//query := &vpc.ModifyRouteTableAttributesInput{
	//	RouteTableId:   volcstack.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	RouteTableName: volcstack.String("test"),
	//	Description:   	volcstack.String("test"),
	//}
	//resp, err := svc.ModifyRouteTableAttributes(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteTable DescribeRouteTableList
	//query := &vpc.DescribeRouteTableListInput{
	//	VpcId:          volcstack.String("vpc-2fevwte3frk0059gp67d69r4j"),
	//	RouteTableId:   volcstack.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	RouteTableName: volcstack.String("test"),
	//	PageSize:   	volcstack.Int64(20),
	//	PageNumber:   	volcstack.Int64(1),
	//}
	//resp, err := svc.DescribeRouteTableList(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteTable AssociateRouteTable
	//query := &vpc.AssociateRouteTableInput{
	//	RouteTableId:   volcstack.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	SubnetId: 		volcstack.String("subnet-12b0gepopmdxc17q7y1ros7up"),
	//}
	//resp, err := svc.AssociateRouteTable(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteTable DisassociateRouteTable
	//query := &vpc.DisassociateRouteTableInput{
	//	RouteTableId:   volcstack.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	SubnetId: 		volcstack.String("subnet-12b0gepopmdxc17q7y1ros7up"),
	//}
	//resp, err := svc.DisassociateRouteTable(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteEntry CreateRouteEntry
	//query := &vpc.CreateRouteEntryInput{
	//	RouteTableId:   		volcstack.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	DestinationCidrBlock:  	volcstack.String("192.168.0.0/16"),
	//	NextHopType:   			volcstack.String("Instance"),
	//	NextHopId:   			volcstack.String("i-ybkbqfzie1l8j1or2ke5"),
	//	RouteEntryName:   		volcstack.String("test"),
	//	Description:   			volcstack.String("test"),
	//}
	//resp, err := svc.CreateRouteEntry(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteEntry ModifyRouteEntry
	//query := &vpc.ModifyRouteEntryInput{
	//	RouteEntryId:   	    volcstack.String("rte-12bhqix7op81s17q7y39jk947"),
	//	RouteEntryName:   		volcstack.String("test-up"),
	//	Description:   			volcstack.String("test-up"),
	//}
	//resp, err := svc.ModifyRouteEntry(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteEntry DescribeRouteEntryList
	//query := &vpc.DescribeRouteEntryListInput{
	//	RouteTableId:   	    volcstack.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	RouteEntryType:   	    volcstack.String("Custom"),
	//	RouteEntryId:   	    volcstack.String("rte-12bhqix7op81s17q7y39jk947"),
	//	RouteEntryName:   		volcstack.String("test-up"),
	//	DestinationCidrBlock:  	volcstack.String("192.168.0.0/16"),
	//	NextHopType:   			volcstack.String("Instance"),
	//	NextHopId:   			volcstack.String("i-ybkbqfzie1l8j1or2ke5"),
	//	PageSize:   			volcstack.Int64(20),
	//	PageNumber:   			volcstack.Int64(1),
	//}
	//resp, err := svc.DescribeRouteEntryList(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteEntry DeleteRouteEntry
	//query := &vpc.DeleteRouteEntryInput{
	//	RouteEntryId:   volcstack.String("rte-12bhqix7op81s17q7y39jk947"),
	//}
	//resp, err := svc.DeleteRouteEntry(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip CreateHaVip
	//query := &vpc.CreateHaVipInput{
	//	SubnetId:    volcstack.String("subnet-12b0gepopmdxc17q7y1ros7up"),
	//	IpAddress:   volcstack.String("192.168.1.21"),
	//	HaVipName:   volcstack.String("test"),
	//	Description: volcstack.String("test"),
	//}
	//resp, err := svc.CreateHaVip(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip ModifyHaVipAttributes
	//query := &vpc.ModifyHaVipAttributesInput{
	//	HaVipId:     volcstack.String("havip-2few45rudt62o59gp67nybytl"),
	//	HaVipName:   volcstack.String("test-up"),
	//	Description: volcstack.String("test-up"),
	//}
	//resp, err := svc.ModifyHaVipAttributes(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip DeleteHaVip
	//query := &vpc.DeleteHaVipInput{
	//	HaVipId:     volcstack.String("havip-2few45rudt62o59gp67nybytl"),
	//}
	//resp, err := svc.DeleteHaVip(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip AssociateHaVip
	//query := &vpc.AssociateHaVipInput{
	//	HaVipId:     	volcstack.String("havip-imv9e3nn8q2o8gbssxfxarvx"),
	//	InstanceId:     volcstack.String("i-l8j1k9sk6cjww1u29v09"),
	//	InstanceType:   volcstack.String("EcsInstance"),
	//}
	//resp, err := svc.AssociateHaVip(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip DisassociateHaVip
	//query := &vpc.DisassociateHaVipInput{
	//	HaVipId:     	volcstack.String("havip-imv9e3nn8q2o8gbssxfxarvx"),
	//	InstanceId:     volcstack.String("i-l8j1k9sk6cjww1u29v09"),
	//	InstanceType:   volcstack.String("EcsInstance"),
	//}
	//resp, err := svc.DisassociateHaVip(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip DescribeHaVips
	query := &vpc.DescribeHaVipsInput{
		VpcId:      volcstack.String("vpc-2fevwte3frk0059gp67d69r4j"),
		SubnetId:   volcstack.String("subnet-12b0gepopmdxc17q7y1ros7up"),
		IpAddress:  volcstack.String("192.168.1.22"),
		HaVipIds:   []*string{volcstack.String("havip-imv9e3nn8q2o8gbssxfxarvx")},
		HaVipName:  volcstack.String("tf-test"),
		Status:     volcstack.String("Available"),
		PageSize:   volcstack.Int64(20),
		PageNumber: volcstack.Int64(1),
	}
	resp, err := svc.DescribeHaVips(query)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

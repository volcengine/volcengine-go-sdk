package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	ak := "your ak"
	sk := "your sk"
	var region = "cn-beijing"

	//if use env Credentials
	//please
	//export VOLCSTACK_ACCESS_KEY=AK
	//export VOLCSTACK_SECRET_KEY=SK
	// and WithCredentials(credentials.NewEnvCredentials())
	config := volcengine.NewConfig().
		WithRegion(region).
		WithAkSk(ak, sk).
		//WithCredentials(credentials.NewEnvCredentials()).
		WithDisableSSL(true)
	//WithLogLevel(volcengine.LogDebugWithHTTPBody).
	//WithEndpoint(volcengineutil.NewEndpoint().GetEndpoint())
	sess, _ := session.NewSession(config)
	svc := vpc.New(sess)

	// ENI CreateNetworkInterface
	//query := &vpc.CreateNetworkInterfaceInput{
	//	SubnetId:     			volcengine.String("subnet-274xwulapzc3k7fap8thue3pl"),
	//	SecurityGroupIds: 		[]*string{volcengine.String("sg-2744odkqfhkhs7fap8tnp49r6")},
	//	PrimaryIpAddress: 		volcengine.String("192.168.1.2"),
	//	NetworkInterfaceName: 	volcengine.String("test"),
	//	Description: 			volcengine.String("test"),
	//	PortSecurityEnabled: 	volcengine.Bool(false),
	//}
	//resp, err := svc.CreateNetworkInterface(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI ModifyNetworkInterfaceAttributes
	//query := &vpc.ModifyNetworkInterfaceAttributesInput{
	//	NetworkInterfaceId:     volcengine.String("eni-274t9vh8ndm9s7fap8shjokny"),
	//	SecurityGroupIds: 		[]*string{volcengine.String("sg-2744odkqfhkhs7fap8tnp49r6"), volcengine.String("sg-274k348d2mhog7fap8thr09gd")},
	//	NetworkInterfaceName: 	volcengine.String("test-up"),
	//	Description: 			volcengine.String("test-up"),
	//}
	//resp, err := svc.ModifyNetworkInterfaceAttributes(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI DescribeNetworkInterfaces
	//query := &vpc.DescribeNetworkInterfacesInput{
	//	Type: 					volcengine.String("secondary"),
	//	Status: 			    volcengine.String("Available"),
	//	VpcId:                  volcengine.String("vpc-2744odetc5ekg7fap8tlh9dl3"),
	//	SubnetId:     			volcengine.String("subnet-274xwulapzc3k7fap8thue3pl"),
	//	//InstanceId:     	    volcengine.String("subnet-274xwulapzc3k7fap8thue3pl"),
	//	PrimaryIpAddresses:     []*string{volcengine.String("192.168.1.2")},
	//	SecurityGroupId:        volcengine.String("sg-2744odkqfhkhs7fap8tnp49r6"),
	//	NetworkInterfaceIds:    []*string{volcengine.String("eni-274t9vh8ndm9s7fap8shjokny")},
	//	NetworkInterfaceName: 	volcengine.String("test-up"),
	//	PageNumber:             volcengine.Int64(1),
	//	PageSize:             	volcengine.Int64(20),
	//	//PrivateIpAddresses:     []*string{volcengine.String("192.168.1.2")},
	//	//ZoneId:                 volcengine.String("cn-lingqiu-a"),
	//}
	//resp, err := svc.DescribeNetworkInterfaces(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI DescribeNetworkInterfaceAttributes
	//query := &vpc.DescribeNetworkInterfaceAttributesInput{
	//	NetworkInterfaceId:    volcengine.String("eni-274ngjt4wp2io7fap8sfxgfzn"),
	//}
	//resp, err := svc.DescribeNetworkInterfaceAttributes(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI DeleteNetworkInterface
	//query := &vpc.DeleteNetworkInterfaceInput{
	//	NetworkInterfaceId:    volcengine.String("eni-274t9vh8ndm9s7fap8shjokny"),
	//}
	//resp, err := svc.DeleteNetworkInterface(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI AttachNetworkInterface
	//query := &vpc.AttachNetworkInterfaceInput{
	//	NetworkInterfaceId:    volcengine.String("eni-2few40gfjx2bk59gp68wu110z"),
	//	InstanceId:    		   volcengine.String("i-ybkbqfzie1l8j1or2ke5"),
	//}
	//resp, err := svc.AttachNetworkInterface(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// ENI DetachNetworkInterface
	//query := &vpc.DetachNetworkInterfaceInput{
	//	NetworkInterfaceId:    volcengine.String("eni-2few40gfjx2bk59gp68wu110z"),
	//	InstanceId:    		   volcengine.String("i-ybkbqfzie1l8j1or2ke5"),
	//}
	//resp, err := svc.DetachNetworkInterface(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// todo 增加AssignPrivateIpAddresses、UnassignPrivateIpAddresses

	// RouteTable CreateRouteTable
	//query := &vpc.CreateRouteTableInput{
	//	VpcId:     			volcengine.String("vpc-2fevwte3frk0059gp67d69r4j"),
	//	RouteTableName: 	volcengine.String("test"),
	//	Description: 	    volcengine.String("test"),
	//}
	//resp, err := svc.CreateRouteTable(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteTable DeleteRouteTable
	//query := &vpc.DeleteRouteTableInput{
	//	RouteTableId:  volcengine.String("vtb-12bhqhvptkl4w17q7y1pv9qmc"),
	//}
	//resp, err := svc.DeleteRouteTable(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteTable ModifyRouteTableAttributes
	//query := &vpc.ModifyRouteTableAttributesInput{
	//	RouteTableId:   volcengine.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	RouteTableName: volcengine.String("test"),
	//	Description:   	volcengine.String("test"),
	//}
	//resp, err := svc.ModifyRouteTableAttributes(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteTable DescribeRouteTableList
	//query := &vpc.DescribeRouteTableListInput{
	//	VpcId:          volcengine.String("vpc-2fevwte3frk0059gp67d69r4j"),
	//	RouteTableId:   volcengine.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	RouteTableName: volcengine.String("test"),
	//	PageSize:   	volcengine.Int64(20),
	//	PageNumber:   	volcengine.Int64(1),
	//}
	//resp, err := svc.DescribeRouteTableList(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteTable AssociateRouteTable
	//query := &vpc.AssociateRouteTableInput{
	//	RouteTableId:   volcengine.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	SubnetId: 		volcengine.String("subnet-12b0gepopmdxc17q7y1ros7up"),
	//}
	//resp, err := svc.AssociateRouteTable(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteTable DisassociateRouteTable
	//query := &vpc.DisassociateRouteTableInput{
	//	RouteTableId:   volcengine.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	SubnetId: 		volcengine.String("subnet-12b0gepopmdxc17q7y1ros7up"),
	//}
	//resp, err := svc.DisassociateRouteTable(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteEntry CreateRouteEntry
	//query := &vpc.CreateRouteEntryInput{
	//	RouteTableId:   		volcengine.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	DestinationCidrBlock:  	volcengine.String("192.168.0.0/16"),
	//	NextHopType:   			volcengine.String("Instance"),
	//	NextHopId:   			volcengine.String("i-ybkbqfzie1l8j1or2ke5"),
	//	RouteEntryName:   		volcengine.String("test"),
	//	Description:   			volcengine.String("test"),
	//}
	//resp, err := svc.CreateRouteEntry(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteEntry ModifyRouteEntry
	//query := &vpc.ModifyRouteEntryInput{
	//	RouteEntryId:   	    volcengine.String("rte-12bhqix7op81s17q7y39jk947"),
	//	RouteEntryName:   		volcengine.String("test-up"),
	//	Description:   			volcengine.String("test-up"),
	//}
	//resp, err := svc.ModifyRouteEntry(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteEntry DescribeRouteEntryList
	//query := &vpc.DescribeRouteEntryListInput{
	//	RouteTableId:   	    volcengine.String("vtb-12bhqi3ly0t1c17q7y2aham35"),
	//	RouteEntryType:   	    volcengine.String("Custom"),
	//	RouteEntryId:   	    volcengine.String("rte-12bhqix7op81s17q7y39jk947"),
	//	RouteEntryName:   		volcengine.String("test-up"),
	//	DestinationCidrBlock:  	volcengine.String("192.168.0.0/16"),
	//	NextHopType:   			volcengine.String("Instance"),
	//	NextHopId:   			volcengine.String("i-ybkbqfzie1l8j1or2ke5"),
	//	PageSize:   			volcengine.Int64(20),
	//	PageNumber:   			volcengine.Int64(1),
	//}
	//resp, err := svc.DescribeRouteEntryList(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// RouteEntry DeleteRouteEntry
	//query := &vpc.DeleteRouteEntryInput{
	//	RouteEntryId:   volcengine.String("rte-12bhqix7op81s17q7y39jk947"),
	//}
	//resp, err := svc.DeleteRouteEntry(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip CreateHaVip
	//query := &vpc.CreateHaVipInput{
	//	SubnetId:    volcengine.String("subnet-12b0gepopmdxc17q7y1ros7up"),
	//	IpAddress:   volcengine.String("192.168.1.21"),
	//	HaVipName:   volcengine.String("test"),
	//	Description: volcengine.String("test"),
	//}
	//resp, err := svc.CreateHaVip(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip ModifyHaVipAttributes
	//query := &vpc.ModifyHaVipAttributesInput{
	//	HaVipId:     volcengine.String("havip-2few45rudt62o59gp67nybytl"),
	//	HaVipName:   volcengine.String("test-up"),
	//	Description: volcengine.String("test-up"),
	//}
	//resp, err := svc.ModifyHaVipAttributes(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip DeleteHaVip
	//query := &vpc.DeleteHaVipInput{
	//	HaVipId:     volcengine.String("havip-2few45rudt62o59gp67nybytl"),
	//}
	//resp, err := svc.DeleteHaVip(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip AssociateHaVip
	//query := &vpc.AssociateHaVipInput{
	//	HaVipId:     	volcengine.String("havip-imv9e3nn8q2o8gbssxfxarvx"),
	//	InstanceId:     volcengine.String("i-l8j1k9sk6cjww1u29v09"),
	//	InstanceType:   volcengine.String("EcsInstance"),
	//}
	//resp, err := svc.AssociateHaVip(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip DisassociateHaVip
	//query := &vpc.DisassociateHaVipInput{
	//	HaVipId:     	volcengine.String("havip-imv9e3nn8q2o8gbssxfxarvx"),
	//	InstanceId:     volcengine.String("i-l8j1k9sk6cjww1u29v09"),
	//	InstanceType:   volcengine.String("EcsInstance"),
	//}
	//resp, err := svc.DisassociateHaVip(query)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)

	// HaVip DescribeHaVips
	query := &vpc.DescribeHaVipsInput{
		VpcId:      volcengine.String("vpc-2fevwte3frk0059gp67d69r4j"),
		SubnetId:   volcengine.String("subnet-12b0gepopmdxc17q7y1ros7up"),
		IpAddress:  volcengine.String("192.168.1.22"),
		HaVipIds:   []*string{volcengine.String("havip-imv9e3nn8q2o8gbssxfxarvx")},
		HaVipName:  volcengine.String("tf-test"),
		Status:     volcengine.String("Available"),
		PageSize:   volcengine.Int64(20),
		PageNumber: volcengine.Int64(1),
	}
	resp, err := svc.DescribeHaVips(query)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
}

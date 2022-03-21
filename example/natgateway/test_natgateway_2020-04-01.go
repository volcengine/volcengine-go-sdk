package main

import (
	"fmt"

	"github.com/volcengine/volcstack-go-sdk/service/natgateway"
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
	svc := natgateway.New(sess)

	// nat CreateSnatEntry
	//in := &natgateway.CreateSnatEntryInput{
	//	NatGatewayId: 	 volcstack.String("ngw-274zlbpdrwz5s7fap8sr2zc2i"),
	//	SubnetId: 	  	 volcstack.String("subnet-2740cym8mv9q87fap8u3hfx4i"),
	//	EipId: 	 		 volcstack.String("eip-274t1gtv5fugw7fap8scozv4s"),
	//	SnatEntryName: 	 volcstack.String("test"),
	//}
	//resp, err := svc.CreateSnatEntry(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// nat DescribeSnatEntryAttributes
	//in := &natgateway.DescribeSnatEntryAttributesInput{
	//	SnatEntryId: 	 volcstack.String("snat-274t6q0ye8g007fap8tv89xs4"),
	//}
	//resp, err := svc.DescribeSnatEntryAttributes(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// nat ModifySnatEntryAttributes
	//in := &natgateway.ModifySnatEntryAttributesInput{
	//	SnatEntryId: 	 volcstack.String("snat-274t6q0ye8g007fap8tv89xs4"),
	//	SnatEntryName: 	 volcstack.String("test"),
	//	EipId: 	 		 volcstack.String("eip-274t1gtv5fugw7fap8scozv4s"),
	//}
	//resp, err := svc.ModifySnatEntryAttributes(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// nat DescribeSnatEntries
	//in := &natgateway.DescribeSnatEntriesInput{
	//	PageSize:        volcstack.Int64(20),
	//	PageNumber:      volcstack.Int64(1),
	//	SnatEntryIds: 	 []*string{volcstack.String("snat-274t6q0ye8g007fap8tv89xs4")},
	//	SnatEntryName: 	 volcstack.String("test"),
	//	NatGatewayId: 	 volcstack.String("ngw-274zlbpdrwz5s7fap8sr2zc2i"),
	//	SubnetId: 	 	 volcstack.String("subnet-2740cym8mv9q87fap8u3hfx4i"),
	//	EipId: 	 		 volcstack.String("eip-274t1gtv5fugw7fap8scozv4s"),
	//}
	//resp, err := svc.DescribeSnatEntries(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// nat DeleteSnatEntry
	in := &natgateway.DeleteSnatEntryInput{
		SnatEntryId: volcstack.String("snat-274t6q0ye8g007fap8tv89xs4"),
	}
	resp, err := svc.DeleteSnatEntry(in)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v", resp)
}

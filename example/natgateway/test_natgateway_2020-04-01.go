package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
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
	svc := natgateway.New(sess)

	// nat CreateSnatEntry
	//in := &natgateway.CreateSnatEntryInput{
	//	NatGatewayId: 	 volcengine.String("ngw-274zlbpdrwz5s7fap8sr2zc2i"),
	//	SubnetId: 	  	 volcengine.String("subnet-2740cym8mv9q87fap8u3hfx4i"),
	//	EipId: 	 		 volcengine.String("eip-274t1gtv5fugw7fap8scozv4s"),
	//	SnatEntryName: 	 volcengine.String("test"),
	//}
	//resp, err := svc.CreateSnatEntry(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// nat DescribeSnatEntryAttributes
	//in := &natgateway.DescribeSnatEntryAttributesInput{
	//	SnatEntryId: 	 volcengine.String("snat-274t6q0ye8g007fap8tv89xs4"),
	//}
	//resp, err := svc.DescribeSnatEntryAttributes(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// nat ModifySnatEntryAttributes
	//in := &natgateway.ModifySnatEntryAttributesInput{
	//	SnatEntryId: 	 volcengine.String("snat-274t6q0ye8g007fap8tv89xs4"),
	//	SnatEntryName: 	 volcengine.String("test"),
	//	EipId: 	 		 volcengine.String("eip-274t1gtv5fugw7fap8scozv4s"),
	//}
	//resp, err := svc.ModifySnatEntryAttributes(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// nat DescribeSnatEntries
	//in := &natgateway.DescribeSnatEntriesInput{
	//	PageSize:        volcengine.Int64(20),
	//	PageNumber:      volcengine.Int64(1),
	//	SnatEntryIds: 	 []*string{volcengine.String("snat-274t6q0ye8g007fap8tv89xs4")},
	//	SnatEntryName: 	 volcengine.String("test"),
	//	NatGatewayId: 	 volcengine.String("ngw-274zlbpdrwz5s7fap8sr2zc2i"),
	//	SubnetId: 	 	 volcengine.String("subnet-2740cym8mv9q87fap8u3hfx4i"),
	//	EipId: 	 		 volcengine.String("eip-274t1gtv5fugw7fap8scozv4s"),
	//}
	//resp, err := svc.DescribeSnatEntries(in)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v", resp)

	// nat DeleteSnatEntry
	in := &natgateway.DeleteSnatEntryInput{
		SnatEntryId: volcengine.String("snat-274t6q0ye8g007fap8tv89xs4"),
	}
	resp, err := svc.DeleteSnatEntry(in)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v", resp)
}

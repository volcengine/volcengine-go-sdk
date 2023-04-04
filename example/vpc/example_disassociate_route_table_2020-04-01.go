package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DisassociateRouteTable() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	disassociateRouteTableInput := &vpc.DisassociateRouteTableInput{
		RouteTableId: volcengine.String("vpc-2fdz8g5nruzuo5oxr******"),
		SubnetId: volcengine.String("ubnet-2fdzaou4liw3k5o******"),
	}

	resp, err := svc.DisassociateRouteTable(disassociateRouteTableInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

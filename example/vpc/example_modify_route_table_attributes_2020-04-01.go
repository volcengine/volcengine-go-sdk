package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyRouteTableAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	modifyRouteTableAttributesInput := &vpc.ModifyRouteTableAttributesInput{
		RouteTableId: volcengine.String("vtb-2fdzao4h726f45****"),
	}

	resp, err := svc.ModifyRouteTableAttributes(modifyRouteTableAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

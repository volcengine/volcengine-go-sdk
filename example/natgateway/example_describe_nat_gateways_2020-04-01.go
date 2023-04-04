package natgateway_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeNatGateways() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := natgateway.New(sess)
	describeNatGatewaysInput := &natgateway.DescribeNatGatewaysInput{
		NatGatewayIds: volcengine.StringSlice([]string{"ngw-2fesmko5zhdz459gp67sc****"}),
		PageNumber:    volcengine.Int64(1),
		PageSize:      volcengine.Int64(5),
	}

	resp, err := svc.DescribeNatGateways(describeNatGatewaysInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

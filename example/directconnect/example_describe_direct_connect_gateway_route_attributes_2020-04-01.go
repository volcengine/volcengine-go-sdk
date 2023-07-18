// Example Code generated by Beijing Volcanoengine Technology.
package directconnectexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/directconnect"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeDirectConnectGatewayRouteAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := directconnect.New(sess)
	describeDirectConnectGatewayRouteAttributesInput := &directconnect.DescribeDirectConnectGatewayRouteAttributesInput{
		DirectConnectGatewayRouteId: volcengine.String("dcr-2752hzaqrwfeo7f****"),
	}

	resp, err := svc.DescribeDirectConnectGatewayRouteAttributes(describeDirectConnectGatewayRouteAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

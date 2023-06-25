// Example Code generated by Beijing Volcanoengine Technology.
package directconnectemample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/directconnect"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeDirectConnectConnectionAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := directconnect.New(sess)
	describeDirectConnectConnectionAttributesInput := &directconnect.DescribeDirectConnectConnectionAttributesInput{
		DirectConnectConnectionId: volcengine.String("dcc-7qthudw0ll6jmc****"),
	}

	resp, err := svc.DescribeDirectConnectConnectionAttributes(describeDirectConnectConnectionAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

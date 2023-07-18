// Example Code generated by Beijing Volcanoengine Technology.
package directconnectexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/directconnect"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyDirectConnectVirtualInterfaceAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := directconnect.New(sess)
	modifyDirectConnectVirtualInterfaceAttributesInput := &directconnect.ModifyDirectConnectVirtualInterfaceAttributesInput{
		Description:          volcengine.String("test"),
		VirtualInterfaceId:   volcengine.String("dcv-7qthudw0ll6jmc****"),
		VirtualInterfaceName: volcengine.String("test"),
	}

	resp, err := svc.ModifyDirectConnectVirtualInterfaceAttributes(modifyDirectConnectVirtualInterfaceAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

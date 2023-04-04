package natgateway_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func CreateDnatEntry() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := natgateway.New(sess)
	createDnatEntryInput := &natgateway.CreateDnatEntryInput{
		ExternalIp: volcengine.String("12.XX.XX.34"),
		ExternalPort: volcengine.String("34"),
		InternalIp: volcengine.String("192.XX.XX.88"),
		InternalPort: volcengine.String("12"),
		NatGatewayId: volcengine.String("ngw-2feq5xhimd88w59gp686****"),
		Protocol: volcengine.String("tcp"),
	}

	resp, err := svc.CreateDnatEntry(createDnatEntryInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

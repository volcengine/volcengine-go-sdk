package natgateway_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeDnatEntries() {
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
	describeDnatEntriesInput := &natgateway.DescribeDnatEntriesInput{
		DnatEntryIds: volcengine.StringSlice([]string{"dnat-342abc3bc3****"}),
		NatGatewayId: volcengine.String("ngw-2feq5xhimd88w59gp686****"),
	}

	resp, err := svc.DescribeDnatEntries(describeDnatEntriesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package natgateway_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeSnatEntries() {
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
	describeSnatEntriesInput := &natgateway.DescribeSnatEntriesInput{
		NatGatewayId: volcengine.String("ngw-2fedgzyvtzaio59gp675l****"),
		SnatEntryIds: volcengine.StringSlice([]string{"snat-2fedi22b12iv459gp68****","snat-2fedhzdlyknb459gp676o****"}),
	}

	resp, err := svc.DescribeSnatEntries(describeSnatEntriesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

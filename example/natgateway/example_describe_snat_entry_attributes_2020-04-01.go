// Example Code generated by Beijing Volcanoengine Technology.
package natgatewayexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeSnatEntryAttributes() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := natgateway.New(sess)
	describeSnatEntryAttributesInput := &natgateway.DescribeSnatEntryAttributesInput{
		SnatEntryId: volcengine.String("snat-2fedi096gdiww59gp680****"),
	}

	resp, err := svc.DescribeSnatEntryAttributes(describeSnatEntryAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

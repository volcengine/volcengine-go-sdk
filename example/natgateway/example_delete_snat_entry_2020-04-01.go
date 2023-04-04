package natgatewayexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DeleteSnatEntry() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := natgateway.New(sess)
	deleteSnatEntryInput := &natgateway.DeleteSnatEntryInput{
		SnatEntryId: volcengine.String("snat-2fedi096gdiww59gp680****"),
	}

	resp, err := svc.DeleteSnatEntry(deleteSnatEntryInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package natgatewayexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DeleteDnatEntry() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := natgateway.New(sess)
	deleteDnatEntryInput := &natgateway.DeleteDnatEntryInput{
		DnatEntryId: volcengine.String("dnat-342abc3bc3****"),
	}

	resp, err := svc.DeleteDnatEntry(deleteDnatEntryInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
package billingexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/billing"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func UnsubscribeInstance() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := billing.New(sess)
	unsubscribeInstanceInput := &billing.UnsubscribeInstanceInput{
		ClientToken:                volcengine.String("uuid-xxxxxxxxxxxx"),
		InstanceID:                 volcengine.String("vol-924610xxxxx"),
		Product:                    volcengine.String("volume"),
		UnsubscribeRelatedInstance: volcengine.Bool(true),
	}

	resp, err := svc.UnsubscribeInstance(unsubscribeInstanceInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

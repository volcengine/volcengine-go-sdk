package ecsexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateSubscription() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	createSubscriptionInput := &ecs.CreateSubscriptionInput{
		EventTypes: volcengine.StringSlice([]string{"SystemFailure.Stop:Succeeded", "SystemFailure.Stop:Succeeded"}),
		Type:       volcengine.String("Notification"),
	}

	resp, err := svc.CreateSubscription(createSubscriptionInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package ecsexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func AttachKeyPair() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	attachKeyPairInput := &ecs.AttachKeyPairInput{
		InstanceIds: volcengine.StringSlice([]string{"i-ahipakt2gcg95jpv****", "i-ahipakt2gdg95lbe****"}),
		KeyPairName: volcengine.String("ssh_key_pair1"),
	}

	resp, err := svc.AttachKeyPair(attachKeyPairInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

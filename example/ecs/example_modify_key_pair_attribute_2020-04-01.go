package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ModifyKeyPairAttribute() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	modifyKeyPairAttributeInput := &ecs.ModifyKeyPairAttributeInput{
		Description: volcengine.String("ssh_key_test_description"),
		KeyPairName: volcengine.String("ssh_key_pair"),
	}

	resp, err := svc.ModifyKeyPairAttribute(modifyKeyPairAttributeInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

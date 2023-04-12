package ecsexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DeleteKeyPairs() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	deleteKeyPairsInput := &ecs.DeleteKeyPairsInput{
		KeyPairNames: volcengine.StringSlice([]string{"ssh_key_pair1", "ssh_key_pair2"}),
	}

	resp, err := svc.DeleteKeyPairs(deleteKeyPairsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

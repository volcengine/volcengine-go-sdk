package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ImportKeyPair() {
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
	importKeyPairInput := &ecs.ImportKeyPairInput{
		KeyPairName: volcengine.String("ssh_key_pair"),
		PublicKey: volcengine.String("ssh-rsa AaaAAB3NzaC1yc2EAAAADAQ******"),
	}

	resp, err := svc.ImportKeyPair(importKeyPairInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
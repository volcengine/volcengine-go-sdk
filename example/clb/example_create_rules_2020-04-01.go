package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func CreateRules() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	reqRules0 := &clb.RuleForCreateRulesInput{
		Domain: volcengine.String("test.com"),
		ServerGroupId: volcengine.String("rsp-bp1o94dp5i6ea****"),
		Url: volcengine.String("/test"),
	}
	createRulesInput := &clb.CreateRulesInput{
		ListenerId: volcengine.String("lsn-2fea4ayvu2g3k5oxruuz****"),
		Rules: []*clb.RuleForCreateRulesInput{reqRules0},
	}

	resp, err := svc.CreateRules(createRulesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
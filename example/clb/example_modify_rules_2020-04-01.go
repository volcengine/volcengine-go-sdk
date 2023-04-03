package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
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
	reqRules0 := &clb.RuleForModifyRulesInput{
		Description: volcengine.String("test"),
		RuleId: volcengine.String("rule-2fegss1cplxxc5oxruvvq****"),
		ServerGroupId: volcengine.String("rsp-bp1o94dp5i6ea****"),
	}
	modifyRulesInput := &clb.ModifyRulesInput{
		ListenerId: volcengine.String("lsn-2fek3rgsxhrsw5oxruwec****"),
		Rules: []*clb.RuleForModifyRulesInput{reqRules0},
	}

	resp, err := svc.ModifyRules(modifyRulesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
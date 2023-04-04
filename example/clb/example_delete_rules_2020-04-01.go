package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DeleteRules() {
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
	deleteRulesInput := &clb.DeleteRulesInput{
		ListenerId: volcengine.String("lsn-2fek3rgsxhrsw5oxruwec****"),
		RuleIds: volcengine.StringSlice([]string{"rule-2fegss1cplxxc5oxruvvq****"}),
	}

	resp, err := svc.DeleteRules(deleteRulesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

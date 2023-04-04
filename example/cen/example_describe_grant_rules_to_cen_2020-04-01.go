package cen_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeGrantRulesToCen() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	describeGrantRulesToCenInput := &cen.DescribeGrantRulesToCenInput{
		CenId: volcengine.String("cen-uf6o8d1dj8sjwxi6o****"),
	}

	resp, err := svc.DescribeGrantRulesToCen(describeGrantRulesToCenInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

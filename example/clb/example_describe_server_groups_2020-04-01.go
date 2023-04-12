package clbexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeServerGroups() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	describeServerGroupsInput := &clb.DescribeServerGroupsInput{
		ServerGroupIds: volcengine.StringSlice([]string{"rsp-2fel9y8pxr56o5oxruuuu****"}),
	}

	resp, err := svc.DescribeServerGroups(describeServerGroupsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func RemoveAclEntries() {
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
	removeAclEntriesInput := &clb.RemoveAclEntriesInput{
		AclId: volcengine.String("acl-3cj44nv0jhhxc6c6rrtet****"),
		Entries: volcengine.StringSlice([]string{"192.XX.XX.7/16","172.XX.XX.0/24"}),
	}

	resp, err := svc.RemoveAclEntries(removeAclEntriesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

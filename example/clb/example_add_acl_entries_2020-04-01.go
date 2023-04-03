package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func AddAclEntries() {
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
	reqAclEntries0 := &clb.AclEntryForAddAclEntriesInput{
		Entry: volcengine.String("192.XX.XX.7/16"),
	}
	reqAclEntries1 := &clb.AclEntryForAddAclEntriesInput{
		Entry: volcengine.String("192.XX.XX.0/16"),
	}
	addAclEntriesInput := &clb.AddAclEntriesInput{
		AclEntries: []*clb.AclEntryForAddAclEntriesInput{reqAclEntries0,reqAclEntries1},
		AclId: volcengine.String("acl-3cj44nv0jhhxc6c6rrtet****"),
	}

	resp, err := svc.AddAclEntries(addAclEntriesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

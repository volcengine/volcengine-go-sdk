package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func UpdateNetworkAclEntries() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	reqEgressAclEntries0 := &vpc.EgressAclEntryForUpdateNetworkAclEntriesInput{
		Description: volcengine.String("ThisisEgressAclEntries01."),
		DestinationCidrIp: volcengine.String("10.XX.XX.0/24"),
		NetworkAclEntryId: volcengine.String("nae-2zecs97e0brcge46****"),
		NetworkAclEntryName: volcengine.String("acl-01"),
		Policy: volcengine.String("accept"),
		Port: volcengine.String("-1/-1"),
		Protocol: volcengine.String("all"),
	}
	reqEgressAclEntries1 := &vpc.EgressAclEntryForUpdateNetworkAclEntriesInput{
		Description: volcengine.String("ThisisEgressAclEntries02."),
		DestinationCidrIp: volcengine.String("10.XX.XX.0/24"),
		NetworkAclEntryId: volcengine.String("nae-0iswk4d88jvds****"),
		NetworkAclEntryName: volcengine.String("acl-02"),
		Policy: volcengine.String("accept"),
		Port: volcengine.String("80/80"),
		Protocol: volcengine.String("icmp"),
	}
	reqIngressAclEntries0 := &vpc.IngressAclEntryForUpdateNetworkAclEntriesInput{
		Description: volcengine.String("ThisisIngressAclEntries01."),
		NetworkAclEntryId: volcengine.String("nae-2zepn32de59j8m4****"),
		NetworkAclEntryName: volcengine.String("acl-3***"),
		Policy: volcengine.String("accept"),
		Port: volcengine.String("22/22"),
		Protocol: volcengine.String("all"),
		SourceCidrIp: volcengine.String("10.XX.XX.0/24"),
	}
	reqIngressAclEntries1 := &vpc.IngressAclEntryForUpdateNetworkAclEntriesInput{
		Description: volcengine.String("ThisisIngressAclEntries02."),
		NetworkAclEntryId: volcengine.String("nae-xyz2dmndek90e****"),
		NetworkAclEntryName: volcengine.String("acl-es***"),
		Policy: volcengine.String(""),
		Port: volcengine.String("80/80"),
		Protocol: volcengine.String("tcp"),
		SourceCidrIp: volcengine.String("10.XX.XX.0/24"),
	}
	updateNetworkAclEntriesInput := &vpc.UpdateNetworkAclEntriesInput{
		EgressAclEntries: []*vpc.EgressAclEntryForUpdateNetworkAclEntriesInput{reqEgressAclEntries0,reqEgressAclEntries1},
		IngressAclEntries: []*vpc.IngressAclEntryForUpdateNetworkAclEntriesInput{reqIngressAclEntries0,reqIngressAclEntries1},
		NetworkAclId: volcengine.String("nacl-bp1fg655nh68xyz9****"),
		UpdateEgressAclEntries: volcengine.Bool(false),
		UpdateIngressAclEntries: volcengine.Bool(false),
	}

	resp, err := svc.UpdateNetworkAclEntries(updateNetworkAclEntriesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

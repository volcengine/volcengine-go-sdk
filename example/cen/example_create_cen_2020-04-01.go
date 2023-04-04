package cenexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateCen() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	createCenInput := &cen.CreateCenInput{
		CenName:     volcengine.String("cen-test"),
		Description: volcengine.String("test"),
	}

	resp, err := svc.CreateCen(createCenInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

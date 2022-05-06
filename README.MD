### Volcstack SDK for Go
####  概述
1. 火山引擎-云计算-Go语言SDK
2. Go版本最低1.5+ 推荐是用1.12+ 
3. service目录下的文件请不要修改，都是由代码生成器自动生成
4. 建议使用go mod方式进行依赖管理

##### 代码示例
代码示例在example目录下

	package main

    import (
    "fmt"

	"github.com/volcengine/volcstack-go-sdk/service/vpc"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/credentials"
	"github.com/volcengine/volcstack-go-sdk/volcstack/session"
    )

    func main() {
     var (
      ak     string
      sk     string
      region string
      config *volcstack.Config
      sess   *session.Session
      client *vpc.VPC
      resp   *vpc.DescribeVpcsOutput
      err    error
     )
     ak = "your ak"
     sk = "your sk"
     region = "cn-beijing"
     config = volcstack.NewConfig().
             WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
             WithRegion(region)

	 sess, _ = session.NewSession(config)
	 client = vpc.New(sess)

	 resp, err = client.DescribeVpcs(&vpc.DescribeVpcsInput{
		PageNumber: volcstack.Int64(1),
		PageSize:   volcstack.Int64(10),
	 })

	 if err != nil {
		panic(err)
	 }

	 fmt.Println(&resp)

     }

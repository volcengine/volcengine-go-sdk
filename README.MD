### Volcengine SDK for Go
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

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
    )

    func main() {
     var (
      ak     string
      sk     string
      region string
      config *volcengine.Config
      sess   *session.Session
      client *vpc.VPC
      resp   *vpc.DescribeVpcsOutput
      err    error
     )
     ak = "your ak"
     sk = "your sk"
     region = "cn-beijing"
     config = volcengine.NewConfig().
             WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
             WithRegion(region)

	 sess, err = session.NewSession(config)
     if err != nil {
		panic(err)
	 }

	 client = vpc.New(sess)

	 resp, err = client.DescribeVpcs(&vpc.DescribeVpcsInput{
		PageNumber: volcengine.Int64(1),
		PageSize:   volcengine.Int64(10),
	 })

	 if err != nil {
		panic(err)
	 }

	 fmt.Println(&resp)

     }

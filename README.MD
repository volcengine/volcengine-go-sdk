### Volcengine SDK for Go
####  概述
1. 火山引擎-云计算-Go语言SDK
2. Go版本最低1.5+ 推荐是用1.12+ 
3. service目录下的文件请不要修改，都是由代码生成器自动生成
4. 建议使用go mod方式进行依赖管理

#### Endpoint设置

如果您要自定义SDK的Endpoint，可以按照以下示例代码设置：
```go
config = volcengine.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithRegion(region).WithEndpoint("ecs.cn-beijing-autodriving.volcengineapi.com")
```

火山引擎标准的Endpoint规则说明：

| Regional 服务                                                                                                                          | Global 服务                                                                        |
|--------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------|
| `{service}.{region}.volcengineapi.com` <br> 例如：云服务ecs在cn-beijing-autodriving Region域名为： `ecs.cn-beijing-autodriving.volcengineapi.com` | `{service}.volcengineapi.com` <br> 例如：访问控制iam为Global服务，域名为：`iam.volcengineapi.com` |

注：

- Service中存在_符号时，Endpoint时需转为-符号。存在大写字母时需转成小写。
- 并非所有云产品和Region都支持标准域名，具体请前往您所使用的产品-API参考中查看。


#### 代码示例
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

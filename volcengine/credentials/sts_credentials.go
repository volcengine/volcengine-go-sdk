package credentials

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/sts"
)

const (
	defaultSTSRegion = "cn-beijing"
	defaultSTSHost   = "sts.volcengineapi.com"
)

type StsAssumeRoleProvider struct {
	AccessKey       string
	SecurityKey     string
	SessionToken    string
	RoleName        string
	AccountId       string
	Host            string
	Region          string
	Schema          string
	Timeout         time.Duration
	DurationSeconds int
	Policy          string
	// MaxRetries controls AssumeRole retry attempts.
	// Zero or negative values fall back to DefaultRetryerMaxNumRetries.
	MaxRetries    int
	RetryInterval time.Duration

	// ProxyURL 指定访问 STS 端点时使用的正向代理地址，适用于客户端位于企业内网、
	// 需经统一出口代理访问公网的场景。
	//
	// 格式为 scheme://[user:password@]host:port，例如：
	//   - 无认证：       "http://10.0.0.1:8080"
	//   - 需账号密码：   "http://user:pass@10.0.0.1:8080"
	//
	// 代理认证信息（用户名/密码）以 Basic 方式随请求发送。若含特殊字符，
	// 请按 RFC 3986 进行百分号编码。
	//
	// 留空时先使用绑定客户端中与 STS 协议对应的 HTTPProxy/HTTPSProxy，
	// 再保留原传输层代理行为（标准默认传输层处理 HTTP_PROXY/HTTPS_PROXY/NO_PROXY）。
	// 直接调用 StsAssumeRole 不包含客户端配置兜底层。
	ProxyURL string

	// CAFile 指定用于校验 STS 服务端证书的自定义 CA 证书文件路径（PEM 格式）。
	//
	// 适用于企业代理对 HTTPS 流量做 TLS 拦截（中间人）、使用私有 CA 重新签发证书，
	// 且该 CA 未安装到操作系统信任库的场景。设置后，SDK 将以此 CA 校验证书链。
	//
	// 留空时保留原传输层的根证书配置，标准默认值使用系统根证书。
	CAFile string

	// InsecureSkipVerify 控制是否跳过 STS 服务端 TLS 证书校验。
	//
	// 置为 true 时将不校验证书链与主机名，可用于自签证书或代理拦截等联调场景。
	//
	// 警告：此选项会使连接易受中间人攻击，切勿在生产环境启用。
	// 生产环境请改用 CAFile 显式信任所需的 CA 证书。
	InsecureSkipVerify bool
}

type StsAssumeRoleTime struct {
	CurrentTime string
	ExpiredTime string
}

// StsAssumeRoleOptions contains optional configuration for STS AssumeRole.
type StsAssumeRoleOptions struct {
	SessionToken    string
	Host            string
	Region          string
	Schema          string
	Timeout         time.Duration
	DurationSeconds int
	Policy          string
	// MaxRetries controls AssumeRole retry attempts.
	// Zero or negative values fall back to DefaultRetryerMaxNumRetries.
	MaxRetries    int
	RetryInterval time.Duration

	// ProxyURL 指定访问 STS 端点时使用的正向代理地址，格式
	// scheme://[user:password@]host:port。留空时先使用绑定客户端中与 STS
	// 协议对应的 HTTPProxy/HTTPSProxy，再保留原传输层代理行为。
	// 直接调用 StsAssumeRoleWithOptions 不包含客户端配置兜底层。
	ProxyURL string
	// CAFile 指定用于校验 STS 服务端证书的自定义 CA 证书文件路径（PEM 格式）。
	CAFile string
	// InsecureSkipVerify 控制是否跳过 STS 服务端 TLS 证书校验，仅限联调，生产禁用。
	InsecureSkipVerify bool
}

// StsAssumeRoleWithOptions creates a one-shot AssumeRole call with required
// parameters and optional functional options.
func StsAssumeRoleWithOptions(accessKey, securityKey, roleName, accountId string, optFns ...func(*StsAssumeRoleOptions)) (*Credentials, *StsAssumeRoleTime, error) {
	opts := StsAssumeRoleOptions{
		DurationSeconds: 3600,
	}
	for _, fn := range optFns {
		fn(&opts)
	}
	p := &StsAssumeRoleProvider{
		AccessKey:          accessKey,
		SecurityKey:        securityKey,
		SessionToken:       opts.SessionToken,
		RoleName:           roleName,
		AccountId:          accountId,
		Host:               opts.Host,
		Region:             opts.Region,
		Schema:             opts.Schema,
		Timeout:            opts.Timeout,
		DurationSeconds:    opts.DurationSeconds,
		Policy:             opts.Policy,
		MaxRetries:         opts.MaxRetries,
		RetryInterval:      opts.RetryInterval,
		ProxyURL:           opts.ProxyURL,
		CAFile:             opts.CAFile,
		InsecureSkipVerify: opts.InsecureSkipVerify,
	}
	return stsAssumeRoleInternal(p)
}

func StsAssumeRole(p *StsAssumeRoleProvider) (*Credentials, *StsAssumeRoleTime, error) {
	return stsAssumeRoleInternal(p)
}

// stsAssumeRoleInternal is the shared implementation for StsAssumeRole and
// StsAssumeRoleWithOptions.
func stsAssumeRoleInternal(p *StsAssumeRoleProvider) (*Credentials, *StsAssumeRoleTime, error) {
	ins := sts.NewInstance()
	ins.Client.ServiceInfo.Credentials.Region = defaultSTSRegion
	ins.SetHost(defaultSTSHost)
	if p.Region != "" {
		ins.Client.ServiceInfo.Credentials.Region = p.Region
	}
	if p.Host != "" {
		ins.SetHost(p.Host)
	}
	if p.Schema != "" {
		ins.SetSchema(p.Schema)
	}
	if p.Timeout > 0 {
		ins.Client.SetTimeout(p.Timeout)
	}
	// 仅当显式配置了代理/TLS 时才覆盖传输层：浅拷贝底层 http.Client 并只替换
	// Transport，保留 _GlobalClient 的超时等设置，且不污染全局实例。
	if transport, err := buildProxyTransport(ProxyConfig{
		ProxyURL:           p.ProxyURL,
		CAFile:             p.CAFile,
		InsecureSkipVerify: p.InsecureSkipVerify,
		base:               ins.Client.Client.Transport,
	}); err != nil {
		return nil, nil, err
	} else if transport != nil {
		httpClient := *ins.Client.Client
		httpClient.Transport = transport
		ins.Client.Client = &httpClient
		defer transport.CloseIdleConnections()
	}

	ins.Client.SetAccessKey(p.AccessKey)
	ins.Client.SetSecretKey(p.SecurityKey)
	if p.SessionToken != "" {
		ins.Client.SetSessionToken(p.SessionToken)
	}
	input := &sts.AssumeRoleRequest{
		DurationSeconds: p.DurationSeconds,
		RoleTrn:         fmt.Sprintf("trn:iam::%s:role/%s", p.AccountId, p.RoleName),
		RoleSessionName: uuid.New().String(),
		Policy:          p.Policy,
	}
	var maxRetriesPtr *int
	if p.MaxRetries > 0 {
		maxRetriesPtr = &p.MaxRetries
	}
	output, statusCode, err := assumeRoleWithRetry(ins, input, maxRetriesPtr, p.RetryInterval)
	var reqId string
	if output != nil {
		reqId = output.ResponseMetadata.RequestId
	}
	if err != nil {
		return nil, nil, fmt.Errorf("AssumeRole error,httpcode is %v and reqId is %s error is %s", statusCode, reqId, err.Error())
	}
	if output != nil && output.ResponseMetadata.Error != nil {
		return nil, nil, fmt.Errorf("AssumeRole service error, reqId: %s, code: %s, message: %s",
			reqId, output.ResponseMetadata.Error.Code, output.ResponseMetadata.Error.Message)
	}
	if statusCode >= 300 || statusCode < 200 {
		return nil, nil, fmt.Errorf("AssumeRole error,httpcode is %v and reqId is %s", statusCode, reqId)
	}
	if output == nil || output.Result == nil || output.Result.Credentials == nil {
		return nil, nil, fmt.Errorf("AssumeRole response error,httpcode is %v and reqId is %s", statusCode, reqId)
	}
	return NewCredentials(&StaticProvider{Value: Value{
			AccessKeyID:     output.Result.Credentials.AccessKeyId,
			SecretAccessKey: output.Result.Credentials.SecretAccessKey,
			SessionToken:    output.Result.Credentials.SessionToken,
		}}), &StsAssumeRoleTime{
			CurrentTime: output.Result.Credentials.CurrentTime,
			ExpiredTime: output.Result.Credentials.ExpiredTime,
		}, nil
}

// assumeRoleWithRetry invokes sts.AssumeRole with retries
func assumeRoleWithRetry(ins *sts.STS, input *sts.AssumeRoleRequest, maxRetries *int, retryInterval time.Duration) (*sts.AssumeRoleResp, int, error) {
	resolvedMaxRetries := resolveCredentialMaxRetries(maxRetries)
	if retryInterval <= 0 {
		retryInterval = DefaultRetryerRetryDelay
	}
	var (
		output     *sts.AssumeRoleResp
		statusCode int
		err        error
	)
	for attempt := 0; attempt <= resolvedMaxRetries; attempt++ {
		output, statusCode, err = ins.AssumeRole(input)
		failed := err != nil ||
			statusCode < 200 || statusCode >= 300 ||
			output == nil ||
			output.ResponseMetadata.Error != nil
		if !failed {
			return output, statusCode, err
		}
		// Determine STS error code for retryability check.
		var errCode string
		if output != nil && output.ResponseMetadata.Error != nil {
			errCode = output.ResponseMetadata.Error.Code
		}
		httpStatusCode := statusCode
		if err != nil && statusCode == 0 {
			httpStatusCode = 0
		}
		if !isRetryableSTSError(httpStatusCode, errCode) {
			return output, statusCode, err
		}
		if attempt < resolvedMaxRetries {
			time.Sleep(retryInterval)
		}
	}
	return output, statusCode, err
}

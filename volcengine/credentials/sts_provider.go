package credentials

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/sts"
)

type StsValue StsAssumeRoleProvider

type StsProvider struct {
	Expiry
	StsValue
	stsConfig STSConfig
}

func (s *StsProvider) Retrieve() (Value, error) {
	ins := sts.NewInstance()
	ins.Client.ServiceInfo.Credentials.Region = defaultSTSRegion
	ins.SetHost(defaultSTSHost)
	if s.Region != "" {
		ins.Client.ServiceInfo.Credentials.Region = s.Region
	} else if s.stsConfig.Region != "" {
		ins.Client.ServiceInfo.Credentials.Region = s.stsConfig.Region
	}
	if s.Host != "" {
		ins.SetHost(s.Host)
	} else if s.stsConfig.Host != "" {
		ins.SetHost(s.stsConfig.Host)
	}
	if s.Schema != "" {
		ins.SetSchema(s.Schema)
	}
	if s.Timeout > 0 {
		ins.Client.SetTimeout(s.Timeout)
	}
	// 仅当显式配置了代理/TLS 时才覆盖传输层：浅拷贝底层 http.Client 并只替换
	// Transport，保留 _GlobalClient 的超时等设置，且不污染全局实例。
	if transport, err := buildProxyTransport(ProxyConfig{
		ProxyURL:           s.ProxyURL,
		CAFile:             s.CAFile,
		InsecureSkipVerify: s.InsecureSkipVerify,
		httpProxy:          s.stsConfig.HTTPProxy,
		httpsProxy:         s.stsConfig.HTTPSProxy,
		base:               ins.Client.Client.Transport,
	}); err != nil {
		return Value{}, err
	} else if transport != nil {
		httpClient := *ins.Client.Client
		httpClient.Transport = transport
		ins.Client.Client = &httpClient
		defer transport.CloseIdleConnections()
	}

	if s.DurationSeconds < 900 {
		return Value{}, fmt.Errorf("DurationSeconds must greater than 900 seconds ")
	}

	ins.Client.SetAccessKey(s.AccessKey)
	ins.Client.SetSecretKey(s.SecurityKey)
	if s.SessionToken != "" {
		ins.Client.SetSessionToken(s.SessionToken)
	}
	input := &sts.AssumeRoleRequest{
		DurationSeconds: s.DurationSeconds,
		RoleTrn:         fmt.Sprintf("trn:iam::%s:role/%s", s.AccountId, s.RoleName),
		RoleSessionName: uuid.New().String(),
		Policy:          s.Policy,
	}
	t := time.Now().Add(time.Duration(s.DurationSeconds-60) * time.Second)
	var maxRetriesPtr *int
	if s.MaxRetries > 0 {
		maxRetriesPtr = &s.MaxRetries
	}
	output, _, err := assumeRoleWithRetry(ins, input, maxRetriesPtr, s.RetryInterval)
	if err != nil {
		return Value{}, err
	}
	if output == nil {
		return Value{}, fmt.Errorf("StsProvider: AssumeRole returned nil response")
	}
	if output.ResponseMetadata.Error != nil {
		bb, _err := json.Marshal(output.ResponseMetadata.Error)
		if _err != nil {
			return Value{}, _err
		}
		return Value{}, fmt.Errorf(string(bb))
	}
	if output.Result == nil || output.Result.Credentials == nil {
		return Value{}, fmt.Errorf("StsProvider: AssumeRole returned empty credentials")
	}
	v := Value{
		AccessKeyID:     output.Result.Credentials.AccessKeyId,
		SecretAccessKey: output.Result.Credentials.SecretAccessKey,
		SessionToken:    output.Result.Credentials.SessionToken,
		ProviderName:    "StsProvider",
	}
	s.SetExpiration(t, 0)
	return v, nil
}

func (s *StsProvider) IsExpired() bool {
	return s.Expiry.IsExpired()
}

// NewStsCredentialsWithOptions constructs an StsProvider-backed Credentials
// with required parameters and optional functional options.
func NewStsCredentialsWithOptions(accessKey, securityKey, roleName, accountId string, optFns ...func(*StsAssumeRoleOptions)) *Credentials {
	opts := StsAssumeRoleOptions{
		DurationSeconds: 3600,
	}
	for _, fn := range optFns {
		fn(&opts)
	}
	cfg := StsAssumeRoleProvider{
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
	p := &StsProvider{
		StsValue: StsValue(cfg),
		Expiry:   Expiry{},
	}
	return NewExpireAbleCredentials(p)
}

func NewStsCredentials(value StsValue) *Credentials {

	p := &StsProvider{
		StsValue: value,
		Expiry:   Expiry{},
	}
	return NewExpireAbleCredentials(p)
}

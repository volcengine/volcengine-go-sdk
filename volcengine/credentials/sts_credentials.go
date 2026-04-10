package credentials

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/sts"
)

type StsAssumeRoleProvider struct {
	AccessKey       string
	SecurityKey     string
	RoleName        string
	AccountId       string
	Host            string
	Region          string
	Schema          string
	Timeout         time.Duration
	DurationSeconds int
	Policy          string
	MaxRetries      *int
	RetryInterval   time.Duration
}

type StsAssumeRoleTime struct {
	CurrentTime string
	ExpiredTime string
}

// StsAssumeRoleOptions contains optional configuration for STS AssumeRole.
type StsAssumeRoleOptions struct {
	Host            string
	Region          string
	Schema          string
	Timeout         time.Duration
	DurationSeconds int
	Policy          string
	MaxRetries      *int
	RetryInterval   time.Duration
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
		AccessKey:       accessKey,
		SecurityKey:     securityKey,
		RoleName:        roleName,
		AccountId:       accountId,
		Host:            opts.Host,
		Region:          opts.Region,
		Schema:          opts.Schema,
		Timeout:         opts.Timeout,
		DurationSeconds: opts.DurationSeconds,
		Policy:          opts.Policy,
		MaxRetries:      opts.MaxRetries,
		RetryInterval:   opts.RetryInterval,
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

	ins.Client.SetAccessKey(p.AccessKey)
	ins.Client.SetSecretKey(p.SecurityKey)
	input := &sts.AssumeRoleRequest{
		DurationSeconds: p.DurationSeconds,
		RoleTrn:         fmt.Sprintf("trn:iam::%s:role/%s", p.AccountId, p.RoleName),
		RoleSessionName: uuid.New().String(),
		Policy:          p.Policy,
	}
	output, statusCode, err := assumeRoleWithRetry(ins, input, p.MaxRetries, p.RetryInterval)
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
		retryInterval = DefaultRetryerMinRetryDelay
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
		if attempt < resolvedMaxRetries {
			time.Sleep(retryInterval)
		}
	}
	return output, statusCode, err
}

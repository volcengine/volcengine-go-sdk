// Package defaults is a collection of helpers to retrieve the SDK's default
// configuration and handlers.
//
// Generally this package shouldn't be used directly, but session.Session
// instead. This package is useful when you need to reset the defaults
// of a session or service client to the SDK defaults before setting
// additional parameters.
package defaults

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/corehandlers"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials/clicreds"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials/endpointcreds"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// A Defaults provides a collection of default values for SDK clients.
type Defaults struct {
	Config   *volcengine.Config
	Handlers request.Handlers
}

// Get returns the SDK's default values with Config and handlers pre-configured.
func Get() Defaults {
	cfg := Config()
	handlers := Handlers()
	cfg.Credentials = CredChain(cfg, handlers)

	return Defaults{
		Config:   cfg,
		Handlers: handlers,
	}
}

// Config returns the default configuration without credentials.
// To retrieve a config with credentials also included use
// `defaults.Get().Config` instead.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the configuration of an
// existing service client or session.
func Config() *volcengine.Config {
	return volcengine.NewConfig().
		WithCredentials(credentials.AnonymousCredentials).
		WithRegion(os.Getenv("VOLCSTACK_REGION")).
		WithHTTPClient(http.DefaultClient).
		WithMaxRetries(volcengine.UseServiceDefaultRetries).
		WithLogger(volcengine.NewDefaultLogger()).
		WithLogLevel(volcengine.LogOff)
}

// Handlers returns the default request handlers.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the request handlers of an
// existing service client or session.
func Handlers() request.Handlers {
	var handlers request.Handlers

	handlers.Validate.PushBackNamed(corehandlers.ValidateEndpointHandler)
	handlers.Validate.AfterEachFn = request.HandlerListStopOnError
	handlers.Build.PushBackNamed(corehandlers.CustomerRequestHandler)
	handlers.Build.AfterEachFn = request.HandlerListStopOnError
	handlers.Sign.PushBackNamed(corehandlers.BuildContentLengthHandler)
	handlers.Send.PushBackNamed(corehandlers.ValidateReqSigHandler)
	handlers.Send.PushBackNamed(corehandlers.SendHandler)
	handlers.AfterRetry.PushBackNamed(corehandlers.AfterRetryHandler)
	handlers.ValidateResponse.PushBackNamed(corehandlers.ValidateResponseHandler)

	return handlers
}

// CredChain returns the default credential chain.
//
// Generally you shouldn't need to use this method directly, but
// is available if you need to reset the credentials of an
// existing service client or session's Config.
//
// Note: this now uses DefaultCredentialProvider instead of ChainProvider.
// Error messages always include details from every provider in the chain
// (the old CredentialsChainVerboseErrors config is no longer consulted).
func CredChain(cfg *volcengine.Config, handlers request.Handlers) *credentials.Credentials {
	return credentials.NewDefaultCredentialProviderFromProviders(
		CredProviders(cfg, handlers),
		true, // reuseLastProviderEnabled
	)
}

// CredProviders returns the slice of providers used in
// the default credential chain.
//
// The 5-step default chain:
//  1. EnvProvider (AK/SK/STS from environment variables)
//  2. OIDCCredentialsProvider (from environment variables)
//  3. CliProvider (from ~/.volcengine/config.json)
//  4. SharedCredentialsProvider (from ~/.volcengine/credentials)
//  5. EcsRoleProvider (from IMDS)
func CredProviders(cfg *volcengine.Config, handlers request.Handlers) []credentials.Provider {
	return []credentials.Provider{
		&credentials.EnvProvider{},
		credentials.NewOIDCCredentialsProviderFromEnv(),
		clicreds.NewCliProvider("", ""),
		&credentials.SharedCredentialsProvider{Filename: "", Profile: ""},
		credentials.NewEcsRoleProvider(""),
	}
}

// NewDefaultCredentialProvider creates a default credential chain with the
// given options. This is the primary entry point for users who want to
// customize the default chain (e.g., specify an ECS role name or profile).
//
// Example:
//
//	creds := defaults.NewDefaultCredentialProvider(
//	    func(o *credentials.DefaultCredentialProviderOptions) {
//	        o.RoleName = "my-ecs-role"
//	    },
//	)
func NewDefaultCredentialProvider(optFns ...func(*credentials.DefaultCredentialProviderOptions)) *credentials.Credentials {
	opts := credentials.DefaultCredentialProviderOptions{}
	for _, fn := range optFns {
		fn(&opts)
	}

	providers := []credentials.Provider{
		&credentials.EnvProvider{},
		credentials.NewOIDCCredentialsProviderFromEnv(),
		clicreds.NewCliProvider("", opts.ProfileName),
		&credentials.SharedCredentialsProvider{Filename: "", Profile: opts.ProfileName},
		credentials.NewEcsRoleProvider(opts.RoleName),
	}

	return credentials.NewDefaultCredentialProviderFromProviders(
		providers,
		opts.IsReuseEnabled(),
	)
}

const (
	httpProviderAuthorizationEnvVar = "VOLCSTACK_CONTAINER_AUTHORIZATION_TOKEN"
	httpProviderEnvVar              = "VOLCSTACK_CONTAINER_CREDENTIALS_FULL_URI"
)

var lookupHostFn = net.LookupHost

func isLoopbackHost(host string) (bool, error) {
	ip := net.ParseIP(host)
	if ip != nil {
		return ip.IsLoopback(), nil
	}

	// Host is not an ip, perform lookup
	addrs, err := lookupHostFn(host)
	if err != nil {
		return false, err
	}
	for _, addr := range addrs {
		if !net.ParseIP(addr).IsLoopback() {
			return false, nil
		}
	}

	return true, nil
}

func localHTTPCredProvider(cfg volcengine.Config, handlers request.Handlers, u string) credentials.Provider {
	var errMsg string

	parsed, err := url.Parse(u)
	if err != nil {
		errMsg = fmt.Sprintf("invalid URL, %v", err)
	} else {
		host := volcengine.URLHostname(parsed)
		if len(host) == 0 {
			errMsg = "unable to parse host from local HTTP cred provider URL"
		} else if isLoopback, loopbackErr := isLoopbackHost(host); loopbackErr != nil {
			errMsg = fmt.Sprintf("failed to resolve host %q, %v", host, loopbackErr)
		} else if !isLoopback {
			errMsg = fmt.Sprintf("invalid endpoint host, %q, only loopback hosts are allowed.", host)
		}
	}

	if len(errMsg) > 0 {
		if cfg.Logger != nil {
			cfg.Logger.Error("Ignoring, HTTP credential provider", errMsg, err)
		}
		return credentials.ErrorProvider{
			Err:          volcengineerr.New("CredentialsEndpointError", errMsg, err),
			ProviderName: endpointcreds.ProviderName,
		}
	}

	return httpCredProvider(cfg, handlers, u)
}

func httpCredProvider(cfg volcengine.Config, handlers request.Handlers, u string) credentials.Provider {
	return endpointcreds.NewProviderClient(cfg, handlers, u,
		func(p *endpointcreds.Provider) {
			p.ExpiryWindow = 5 * time.Minute
			p.AuthorizationToken = os.Getenv(httpProviderAuthorizationEnvVar)
		},
	)
}

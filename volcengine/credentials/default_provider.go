package credentials

import (
	"sync/atomic"

	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

// DefaultProviderName is the name used by DefaultCredentialProvider.
const DefaultProviderName = "DefaultCredentialProvider"

// DefaultCredentialProviderOptions holds configuration for the default
// credential chain.
type DefaultCredentialProviderOptions struct {
	// RoleName is the ECS IMDS role name. If empty, it will be resolved
	// from env var VOLCENGINE_ECS_METADATA or auto-detected.
	RoleName string

	// ProfileName is used by CLI config and SharedCredentialsProvider.
	// If empty, each provider falls back to its own default logic.
	ProfileName string

	// ReuseLastProviderEnabled controls whether the chain caches and
	// reuses the last successful provider on subsequent calls.
	// Default is true.
	ReuseLastProviderEnabled *bool
}

// IsReuseEnabled returns whether the reuse-last-provider optimization is
// enabled. Defaults to true if not explicitly set.
func (o *DefaultCredentialProviderOptions) IsReuseEnabled() bool {
	if o.ReuseLastProviderEnabled == nil {
		return true
	}
	return *o.ReuseLastProviderEnabled
}

// DefaultCredentialProvider implements Provider and walks through a chain of
// credential providers. When reuseLastProviderEnabled is true (default), after
// the first successful resolution the chain remembers the provider and reuses
// it for subsequent calls, falling back to full traversal if that provider
// later fails.
type DefaultCredentialProvider struct {
	Providers    []Provider
	lastProvider atomic.Value // stores Provider
	reuseEnabled bool
}

// NewDefaultCredentialProviderFromProviders creates a DefaultCredentialProvider
// from an explicit list of providers. This is intended to be called from the
// defaults package which assembles the full 5-step chain.
func NewDefaultCredentialProviderFromProviders(providers []Provider, reuseEnabled bool) *Credentials {
	p := &DefaultCredentialProvider{
		Providers:    append([]Provider{}, providers...),
		reuseEnabled: reuseEnabled,
	}
	return NewCredentials(p)
}

// Retrieve walks the chain and returns the first successful credentials.
func (d *DefaultCredentialProvider) Retrieve() (Value, error) {
	// Fast path: reuse last successful provider.
	if d.reuseEnabled {
		if last, ok := d.lastProvider.Load().(Provider); ok && last != nil {
			creds, err := last.Retrieve()
			if err == nil {
				return creds, nil
			}
			// Clear cached provider and fall through to full traversal.
			d.lastProvider.Store(Provider(nil))
		}
	}

	// Full traversal.
	var errs []error
	for _, p := range d.Providers {
		creds, err := p.Retrieve()
		if err == nil {
			if d.reuseEnabled {
				d.lastProvider.Store(p)
			}
			return creds, nil
		}
		errs = append(errs, err)
	}

	return Value{},
		volcengineerr.NewBatchError("NoCredentialProviders",
			"no valid providers in default credential chain", errs)
}

// IsExpired checks if the last successful provider's credentials are expired.
// If no provider has been cached yet, returns true to force a Retrieve.
func (d *DefaultCredentialProvider) IsExpired() bool {
	if d.reuseEnabled {
		if last, ok := d.lastProvider.Load().(Provider); ok && last != nil {
			return last.IsExpired()
		}
	}
	return true
}

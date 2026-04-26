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
	// from env var VOLCENGINE_ECS_METADATA.
	RoleName string

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

// providerHolder wraps a Provider so that atomic.Value can store a nil Provider
// without triggering "store of nil value into Value" panic.
type providerHolder struct{ p Provider }

// DefaultCredentialProvider implements Provider and walks through a chain of
// credential providers. When reuseLastProviderEnabled is true (default), after
// the first successful resolution the chain remembers the provider and reuses
// it for subsequent calls, falling back to full traversal if that provider
// later fails.
type DefaultCredentialProvider struct {
	Providers     []Provider
	lastProvider  atomic.Value // stores providerHolder
	reuseEnabled  bool
	verboseErrors bool
}

// NewDefaultCredentialProviderFromProviders creates a DefaultCredentialProvider
// from an explicit list of providers. This is intended to be called from the
// defaults package which assembles the full default chain.
//
// verboseErrors mirrors volcengine.Config.CredentialsChainVerboseErrors: when
// true the failure error includes per-provider details, otherwise the stable
// ErrNoValidProvidersFoundInChain is returned (compatible with the legacy
// ChainProvider behavior).
func NewDefaultCredentialProviderFromProviders(providers []Provider, reuseEnabled, verboseErrors bool) *Credentials {
	p := &DefaultCredentialProvider{
		Providers:     append([]Provider{}, providers...),
		reuseEnabled:  reuseEnabled,
		verboseErrors: verboseErrors,
	}
	return NewCredentials(p)
}

// Retrieve walks the chain and returns the first successful credentials.
func (d *DefaultCredentialProvider) Retrieve() (Value, error) {
	// Fast path: reuse last successful provider.
	if d.reuseEnabled {
		if h, ok := d.lastProvider.Load().(providerHolder); ok && h.p != nil {
			creds, err := h.p.Retrieve()
			if err == nil {
				return creds, nil
			}
			// Clear cached provider and fall through to full traversal.
			d.lastProvider.Store(providerHolder{})
		}
	}

	// Full traversal.
	var errs []error
	for _, p := range d.Providers {
		creds, err := p.Retrieve()
		if err == nil {
			if d.reuseEnabled {
				d.lastProvider.Store(providerHolder{p})
			}
			return creds, nil
		}
		errs = append(errs, err)
	}

	if d.verboseErrors {
		return Value{}, volcengineerr.NewBatchError("NoCredentialProviders",
			"no valid providers in chain", errs)
	}
	return Value{}, ErrNoValidProvidersFoundInChain
}

// IsExpired checks if the last successful provider's credentials are expired.
// If no provider has been cached yet, returns true to force a Retrieve.
func (d *DefaultCredentialProvider) IsExpired() bool {
	if d.reuseEnabled {
		if h, ok := d.lastProvider.Load().(providerHolder); ok && h.p != nil {
			return h.p.IsExpired()
		}
	}
	return true
}

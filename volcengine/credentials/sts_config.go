package credentials

// STSConfig holds client defaults for built-in STS providers. Values are
// snapshotted when a session or service client is created. Explicit provider
// settings take precedence; environment proxy values must not be stored here.
type STSConfig struct {
	Region     string
	Host       string
	HTTPProxy  string
	HTTPSProxy string
}

// WithSTSConfig binds built-in STS credentials to independent client state.
// It does not change the source provider or copy a live credential mutex/cache.
// Custom providers retain their existing sharing and retrieval contract.
func (c *Credentials) WithSTSConfig(cfg STSConfig) *Credentials {
	if c == nil {
		return nil
	}
	c.m.RLock()
	defer c.m.RUnlock()

	switch p := c.provider.(type) {
	case *StsProvider:
		if p == nil {
			return c
		}
		bound := &StsProvider{
			StsValue:  p.StsValue,
			Expiry:    Expiry{CurrentTime: p.CurrentTime},
			stsConfig: cfg,
		}
		return NewCredentials(bound)
	case *OIDCCredentialsProvider:
		if p == nil {
			return c
		}
		bound := *p
		bound.stsConfig = cfg
		bound.lastUpdateTimestamp = 0
		bound.expirationTimestamp = 0
		bound.sessionValue = nil
		bound.MaxRetries = copySTSMaxRetries(p.MaxRetries)
		bound.httpOptions = copySTSHttpOptions(p.httpOptions)
		return NewCredentials(&bound)
	case *SAMLCredentialsProvider:
		if p == nil {
			return c
		}
		bound := *p
		bound.stsConfig = cfg
		bound.lastUpdateTimestamp = 0
		bound.expirationTimestamp = 0
		bound.sessionValue = nil
		bound.MaxRetries = copySTSMaxRetries(p.MaxRetries)
		bound.httpOptions = copySTSHttpOptions(p.httpOptions)
		return NewCredentials(&bound)
	default:
		return c
	}
}

func copySTSMaxRetries(value *int) *int {
	if value == nil {
		return nil
	}
	copy := *value
	return &copy
}

func copySTSHttpOptions(value *HttpOptions) *HttpOptions {
	if value == nil {
		return nil
	}
	copy := *value
	return &copy
}

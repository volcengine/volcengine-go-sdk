package credentials

import (
	"testing"
	"time"
)

func TestNewEcsRoleProvider_DefaultRetryConfig(t *testing.T) {
	p := NewEcsRoleProvider("role")

	if p.RetryInterval != imdsDefaultRetryInterval {
		t.Fatalf("expected RetryInterval=%v, got %v", imdsDefaultRetryInterval, p.RetryInterval)
	}
}

func TestEcsRoleProvider_RetryConfig_DefaultIntervalWhenUnset(t *testing.T) {
	p := &EcsRoleProvider{
		MaxRetries:    2,
		RetryInterval: 0,
	}

	maxRetries, retryInterval := p.retryConfig()
	if maxRetries != 2 {
		t.Fatalf("expected MaxRetries=2, got %d", maxRetries)
	}
	if retryInterval != imdsDefaultRetryInterval {
		t.Fatalf("expected default RetryInterval=%v, got %v", imdsDefaultRetryInterval, retryInterval)
	}
}

func TestEcsRoleProvider_RetryConfig_NegativeRetriesBecomeZero(t *testing.T) {
	p := &EcsRoleProvider{
		MaxRetries:    -1,
		RetryInterval: 10 * time.Millisecond,
	}

	maxRetries, retryInterval := p.retryConfig()
	if maxRetries != 0 {
		t.Fatalf("expected MaxRetries=0, got %d", maxRetries)
	}
	if retryInterval != 10*time.Millisecond {
		t.Fatalf("expected RetryInterval=10ms, got %v", retryInterval)
	}
}

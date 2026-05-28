package credentials

import (
	"testing"
	"time"
)

func intPtr(v int) *int {
	return &v
}

func TestNewEcsRoleProvider_DefaultRetryConfig(t *testing.T) {
	p := NewEcsRoleProvider("role")

	maxRetries, retryInterval := p.retryConfig()
	if maxRetries != DefaultRetryerMaxNumRetries {
		t.Fatalf("expected default MaxRetries=%d, got %d", DefaultRetryerMaxNumRetries, maxRetries)
	}
	if retryInterval != DefaultRetryerRetryDelay {
		t.Fatalf("expected default RetryInterval=%v, got %v", DefaultRetryerRetryDelay, retryInterval)
	}
}

func TestEcsRoleProvider_RetryConfig_DefaultIntervalWhenUnset(t *testing.T) {
	p := &EcsRoleProvider{
		MaxRetries:    intPtr(2),
		RetryInterval: 0,
	}

	maxRetries, retryInterval := p.retryConfig()
	if maxRetries != 2 {
		t.Fatalf("expected MaxRetries=2, got %d", maxRetries)
	}
	if retryInterval != DefaultRetryerRetryDelay {
		t.Fatalf("expected default RetryInterval=%v, got %v", DefaultRetryerRetryDelay, retryInterval)
	}
}

func TestEcsRoleProvider_RetryConfig_ZeroRetriesDisableRetry(t *testing.T) {
	p := &EcsRoleProvider{
		MaxRetries:    intPtr(0),
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

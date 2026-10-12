package session

import (
	"testing"

	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/defaults"
	"github.com/volcengine/volcengine-go-sdk/volcengine/endpoints"
)

func TestClientConfigWithErrRejectsInvalidRegionAndService(t *testing.T) {
	invalid := []string{
		"cn-beijing@127.0.0.1", "cn-beijing:443", "cn/beijing", "cn?beijing",
		"cn#beijing", "cn beijing", "cn\tbeijing", "cn\nbeijing", "cn\x00beijing",
	}

	newSess := func(region string) *Session {
		return &Session{
			Config: defaults.Config().
				WithRegion(region).
				WithCredentials(credentials.NewStaticCredentials("ak", "sk", "")),
			Handlers: defaults.Handlers(),
		}
	}

	if _, err := newSess("cn-beijing").clientConfigWithErr("ecs"); err != nil {
		t.Fatalf("valid region should pass: %v", err)
	}

	for _, region := range invalid {
		_, err := newSess(region).clientConfigWithErr("ecs")
		if err == nil {
			t.Fatalf("region %q should be rejected", region)
		}
	}

	for _, svc := range invalid {
		if _, err := newSess("cn-beijing").clientConfigWithErr(svc); err == nil {
			t.Fatalf("service %q should be rejected", svc)
		}
	}

	// A user supplied endpoint is not built from region/service, keep it working.
	s := newSess("cn-beijing")
	s.Config.WithEndpoint("custom.example.com")
	if _, err := s.clientConfigWithErr("ecs"); err != nil {
		t.Fatalf("custom endpoint should pass: %v", err)
	}

	// Also covers the standard resolver path.
	s = newSess("cn/beijing")
	s.Config.WithEndpointResolver(endpoints.NewStandardEndpointResolver())
	if _, err := s.clientConfigWithErr("ecs"); err == nil {
		t.Fatalf("standard resolver should reject invalid region")
	}
}

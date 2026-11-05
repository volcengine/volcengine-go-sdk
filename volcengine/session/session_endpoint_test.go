package session

import (
	"testing"

	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/defaults"
	"github.com/volcengine/volcengine-go-sdk/volcengine/endpoints"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
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
		want := endpoints.ValidateRegion(region)
		cfg, err := newSess(region).clientConfigWithErr("ecs")
		assertSameError(t, "region "+region, err, want)
		assertFailedClientConfig(t, "region "+region, cfg, want)
	}

	for _, svc := range invalid {
		want := endpoints.ValidateService(svc)
		cfg, err := newSess("cn-beijing").clientConfigWithErr(svc)
		assertSameError(t, "service "+svc, err, want)
		assertFailedClientConfig(t, "service "+svc, cfg, want)
	}

	// A user supplied endpoint is not built from region/service, keep it working.
	s := newSess("cn-beijing")
	s.Config.WithEndpoint("custom.example.com")
	if _, err := s.clientConfigWithErr("ecs"); err != nil {
		t.Fatalf("custom endpoint should pass: %v", err)
	}

	// With a custom EndpointResolver the region is rejected by the session
	// level check before the resolver is called, so a resolver that does not
	// validate its input can not be used to inject a host either. The
	// validation inside StandardEndpointResolver itself is covered by the
	// endpoints package tests.
	s = newSess("cn/beijing")
	s.Config.WithEndpointResolver(endpoints.ResolverFunc(
		func(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
			t.Fatalf("resolver must not be called for invalid region %q", region)
			return endpoints.ResolvedEndpoint{}, nil
		}))
	_, err := s.clientConfigWithErr("ecs")
	assertSameError(t, "custom resolver", err, endpoints.ValidateRegion("cn/beijing"))
}

// assertSameError checks that got is the error produced by the label check.
func assertSameError(t *testing.T, name string, got, want error) {
	t.Helper()
	if want == nil {
		t.Fatalf("%s: test case is expected to be invalid", name)
	}
	if got == nil {
		t.Fatalf("%s: expect error %q, got nil", name, want)
	}
	if got.Error() != want.Error() {
		t.Fatalf("%s: expect error %q, got %q", name, want, got)
	}
}

// assertFailedClientConfig checks that the config returned on failure can be
// dereferenced and that its Validate handlers fail requests with want.
func assertFailedClientConfig(t *testing.T, name string, cfg client.Config, want error) {
	t.Helper()
	if cfg.Config == nil {
		t.Fatalf("%s: returned client.Config.Config must not be nil", name)
	}
	r := &request.Request{}
	cfg.Handlers.Validate.Run(r)
	assertSameError(t, name+": validate handler", r.Error, want)
}

package session_test

import (
	"net/http"
	"sync/atomic"
	"testing"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/endpoints"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/universal"
)

// countingTransport fails the test if any request is sent over the network.
type countingTransport struct {
	calls int32
}

func (c *countingTransport) RoundTrip(*http.Request) (*http.Response, error) {
	atomic.AddInt32(&c.calls, 1)
	return nil, http.ErrUseLastResponse
}

var invalidRegions = []string{
	"cn-beijing@127.0.0.1", "cn-beijing:443", "cn/beijing", "cn?beijing",
	"cn#beijing", "cn beijing", "cn\tbeijing", "cn\nbeijing", "cn\x00beijing",
	"evil.com",
}

func newTestSession(t *testing.T, region string, tr http.RoundTripper, resolver endpoints.Resolver) *session.Session {
	t.Helper()
	cfg := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials("ak", "sk", "")).
		WithHTTPClient(&http.Client{Transport: tr})
	if resolver != nil {
		cfg.WithEndpointResolver(resolver)
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		t.Fatalf("new session: %v", err)
	}
	return sess
}

// The public construction chain (service.New -> request -> Send) must not
// panic for an invalid region, must return an error before anything is sent.
func TestServiceClientInvalidRegionReturnsErrorWithoutPanicOrNetwork(t *testing.T) {
	resolvers := map[string]endpoints.Resolver{
		"default":  nil,
		"standard": endpoints.NewStandardEndpointResolver(),
	}

	for name, resolver := range resolvers {
		for _, region := range invalidRegions {
			tr := &countingTransport{}
			sess := newTestSession(t, region, tr, resolver)

			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Fatalf("[%s] region %q: panic: %v", name, region, r)
					}
				}()

				svc := ecs.New(sess)
				_, err := svc.DescribeInstances(&ecs.DescribeInstancesInput{})
				if err == nil {
					t.Fatalf("[%s] region %q: expect error", name, region)
				}
			}()

			if n := atomic.LoadInt32(&tr.calls); n != 0 {
				t.Fatalf("[%s] region %q: %d network request(s) were sent", name, region, n)
			}
		}
	}
}

func TestUniversalClientInvalidRegionReturnsErrorWithoutPanicOrNetwork(t *testing.T) {
	for _, region := range invalidRegions {
		tr := &countingTransport{}
		sess := newTestSession(t, region, tr, nil)

		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("region %q: panic: %v", region, r)
				}
			}()

			_, err := universal.New(sess).DoCall(universal.RequestUniversal{
				ServiceName: "ecs",
				Action:      "DescribeInstances",
				Version:     "2020-04-01",
				HttpMethod:  universal.GET,
				ContentType: universal.Default,
			}, &map[string]interface{}{})
			if err == nil {
				t.Fatalf("region %q: expect error", region)
			}
		}()

		if n := atomic.LoadInt32(&tr.calls); n != 0 {
			t.Fatalf("region %q: %d network request(s) were sent", region, n)
		}
	}
}

func TestInvalidServiceNameReturnsErrorWithoutPanic(t *testing.T) {
	tr := &countingTransport{}
	sess := newTestSession(t, "cn-beijing", tr, nil)

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic: %v", r)
		}
	}()

	_, err := universal.New(sess).DoCall(universal.RequestUniversal{
		ServiceName: "ecs@127.0.0.1",
		Action:      "DescribeInstances",
		Version:     "2020-04-01",
		HttpMethod:  universal.GET,
		ContentType: universal.Default,
	}, &map[string]interface{}{})
	if err == nil {
		t.Fatalf("expect error for invalid service")
	}
	if n := atomic.LoadInt32(&tr.calls); n != 0 {
		t.Fatalf("%d network request(s) were sent", n)
	}
}

func TestValidRegionStillBuildsClient(t *testing.T) {
	tr := &countingTransport{}
	sess := newTestSession(t, "cn-beijing", tr, nil)

	svc := ecs.New(sess)
	if want := "https://open.volcengineapi.com"; svc.Endpoint != want {
		t.Fatalf("want endpoint %s, got %s", want, svc.Endpoint)
	}
}

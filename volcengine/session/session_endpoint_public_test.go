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
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
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

// assertErrCode checks that err carries the error code of the endpoint label
// check. Asserting err != nil is not enough: the failed client config has no
// endpoint, so ValidateEndpointHandler would fail with MissingEndpoint even if
// the label check error was lost.
func assertErrCode(t *testing.T, name string, err error, code string) {
	t.Helper()
	if err == nil {
		t.Fatalf("%s: expect error with code %s, got nil", name, code)
	}
	e, ok := err.(volcengineerr.Error)
	if !ok || e.Code() != code {
		t.Fatalf("%s: expect error with code %s, got %v", name, code, err)
	}
}

// The public construction chain (service.New -> request -> Send) must not
// panic for an invalid region, must return an error before anything is sent.
// The "standard" case configures a custom EndpointResolver: the region is
// rejected by the session level check before the resolver is called.
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
				assertErrCode(t, "["+name+"] region "+region, err, endpoints.ErrCodeInvalidRegion)
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
			assertErrCode(t, "region "+region, err, endpoints.ErrCodeInvalidRegion)
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
	assertErrCode(t, "invalid service", err, endpoints.ErrCodeInvalidService)
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

package session

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
)

func newProxyTestSession(cfg *volcengine.Config) (*Session, error) {
	options := Options{
		Config:            *cfg,
		SharedConfigFiles: []string{},
	}
	return newSession(options, envConfig{}, &options.Config)
}

func TestConcurrentSessionClientConfigDoesNotMutateDefaultHTTPClient(t *testing.T) {
	previousDefaultClient := http.DefaultClient
	defaultClient := &http.Client{Timeout: 3 * time.Second}
	http.DefaultClient = defaultClient
	t.Cleanup(func() { http.DefaultClient = previousDefaultClient })

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	const clients = 24
	configs := make([]*volcengine.Config, clients)
	sessionClients := make([]*http.Client, clients)
	errs := make([]error, clients)
	requestErrs := make([]error, clients)
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(clients * 2)
	proxy := server.URL
	for i := 0; i < clients; i++ {
		go func(i int) {
			defer wg.Done()
			<-start
			sess, err := newProxyTestSession(volcengine.NewConfig().
				WithCredentials(credentials.NewStaticCredentials("synthetic-ak", "synthetic-sk", "")).
				WithRegion("cn-beijing").
				WithHTTPSProxy(proxy))
			if err != nil {
				errs[i] = err
				return
			}
			sessionClients[i] = sess.Config.HTTPClient
			cfg, err := sess.clientConfigWithErr("vpc")
			if err != nil {
				errs[i] = err
				return
			}
			configs[i] = cfg.Config
		}(i)
		go func(i int) {
			defer wg.Done()
			<-start
			resp, err := defaultClient.Get(server.URL)
			if err != nil {
				requestErrs[i] = err
				return
			}
			resp.Body.Close()
		}(i)
	}
	close(start)
	wg.Wait()

	for i, err := range requestErrs {
		if err != nil {
			t.Errorf("synthetic default-client request %d: %v", i, err)
		}
	}
	if defaultClient.Transport != nil {
		t.Fatal("client configuration mutated http.DefaultClient.Transport")
	}

	seenClients := make(map[*http.Client]struct{}, clients)
	seenTransports := make(map[*http.Transport]struct{}, clients)
	for i, cfg := range configs {
		if errs[i] != nil {
			t.Errorf("client config %d: %v", i, errs[i])
			continue
		}
		if sessionClients[i] != defaultClient {
			t.Errorf("session %d did not use the synthetic default HTTP client", i)
		}
		if cfg == nil || cfg.HTTPClient == nil {
			t.Errorf("client config %d has no HTTP client", i)
			continue
		}
		if cfg.HTTPClient == defaultClient {
			t.Errorf("client config %d reuses the process-wide default HTTP client", i)
		}
		if _, exists := seenClients[cfg.HTTPClient]; exists {
			t.Errorf("client config %d shares an HTTP client with another concurrent config", i)
		}
		seenClients[cfg.HTTPClient] = struct{}{}

		transport, ok := cfg.HTTPClient.Transport.(*http.Transport)
		if !ok {
			t.Errorf("client config %d transport has type %T, want *http.Transport", i, cfg.HTTPClient.Transport)
			continue
		}
		if _, exists := seenTransports[transport]; exists {
			t.Errorf("client config %d shares a transport with another concurrent config", i)
		}
		seenTransports[transport] = struct{}{}

		gotProxy, err := transport.Proxy(&http.Request{URL: &url.URL{Scheme: "https", Host: "api.example.test"}})
		if err != nil {
			t.Errorf("client config %d resolve proxy: %v", i, err)
		} else if gotProxy == nil || gotProxy.String() != proxy {
			t.Errorf("client config %d proxy = %v, want %s", i, gotProxy, proxy)
		}
		if cfg.HTTPClient.Timeout != defaultClient.Timeout {
			t.Errorf("client config %d timeout = %s, want %s", i, cfg.HTTPClient.Timeout, defaultClient.Timeout)
		}
	}
}

func TestClientConfigDoesNotMutateCustomHTTPTransport(t *testing.T) {
	const originalProxy = "http://127.0.0.1:8000"
	const configuredProxy = "http://127.0.0.1:9000"
	originalProxyURL, err := url.Parse(originalProxy)
	if err != nil {
		t.Fatal(err)
	}
	originalTransport := &http.Transport{Proxy: http.ProxyURL(originalProxyURL)}
	customClient := &http.Client{Transport: originalTransport, Timeout: 2 * time.Second}
	sess, err := newProxyTestSession(volcengine.NewConfig().
		WithCredentials(credentials.NewStaticCredentials("synthetic-ak", "synthetic-sk", "")).
		WithRegion("cn-beijing").
		WithHTTPSProxy(configuredProxy).
		WithHTTPClient(customClient))
	if err != nil {
		t.Fatalf("create synthetic session: %v", err)
	}

	cfg, err := sess.clientConfigWithErr("vpc")
	if err != nil {
		t.Fatalf("client config: %v", err)
	}
	if customClient.Transport != originalTransport {
		t.Fatal("client configuration changed the custom HTTP client's transport")
	}
	if cfg.Config.HTTPClient == customClient {
		t.Fatal("client config reuses the caller-owned HTTP client")
	}
	transport, ok := cfg.Config.HTTPClient.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("client config transport has type %T, want *http.Transport", cfg.Config.HTTPClient.Transport)
	}
	if transport == originalTransport {
		t.Fatal("client config reuses the caller-owned HTTP transport")
	}
	if transport.Proxy == nil {
		t.Fatal("client config transport has no proxy function")
	}
	gotOriginalProxy, err := originalTransport.Proxy(&http.Request{URL: &url.URL{Scheme: "https", Host: "api.example.test"}})
	if err != nil {
		t.Fatalf("resolve original proxy: %v", err)
	}
	if gotOriginalProxy == nil || gotOriginalProxy.String() != originalProxy {
		t.Fatalf("original transport proxy = %v, want %s", gotOriginalProxy, originalProxy)
	}
	gotConfiguredProxy, err := transport.Proxy(&http.Request{URL: &url.URL{Scheme: "https", Host: "api.example.test"}})
	if err != nil {
		t.Fatalf("resolve configured proxy: %v", err)
	}
	if gotConfiguredProxy == nil || gotConfiguredProxy.String() != configuredProxy {
		t.Fatalf("client config proxy = %v, want %s", gotConfiguredProxy, configuredProxy)
	}
	if cfg.Config.HTTPClient.Timeout != customClient.Timeout {
		t.Fatalf("client timeout = %s, want %s", cfg.Config.HTTPClient.Timeout, customClient.Timeout)
	}
}

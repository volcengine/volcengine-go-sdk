package defaults_test

// Verifies the retry attempt headers added to volcengine/corehandlers:
// X-Sdk-Invocation-Id must stay the same across every retry attempt of a single
// logical request, and X-Sdk-Request must report the current/max attempt count.

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/volcengine/volcengine-go-sdk/volcengine/client"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/defaults"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

func TestAddRetryInfoHeaderAcrossRetries(t *testing.T) {
	var invocationIDs []string
	var attemptHeaders []string
	callCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		invocationIDs = append(invocationIDs, r.Header.Get("X-Sdk-Invocation-Id"))
		attemptHeaders = append(attemptHeaders, r.Header.Get("X-Sdk-Request"))
		if callCount < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := defaults.Config().WithHTTPClient(server.Client()).WithRegion("test-region")
	handlers := defaults.Handlers()

	c := client.New(*cfg, metadata.ClientInfo{Endpoint: server.URL, SigningRegion: "test-region"}, handlers)
	c.Retryer = client.DefaultRetryer{NumMaxRetries: 2}

	req := c.NewRequest(&request.Operation{Name: "Test", HTTPMethod: "POST", HTTPPath: "/"}, &struct{}{}, &struct{}{})
	if err := req.Send(); err != nil {
		t.Fatalf("expected success after retries, got error: %v", err)
	}

	if callCount != 3 {
		t.Fatalf("expected 3 attempts (1 initial + 2 retries), got %d", callCount)
	}

	for i, id := range invocationIDs {
		if id == "" {
			t.Fatalf("attempt %d: X-Sdk-Invocation-Id header missing", i+1)
		}
		if id != invocationIDs[0] {
			t.Fatalf("attempt %d: X-Sdk-Invocation-Id changed across retries: %q vs %q", i+1, id, invocationIDs[0])
		}
	}

	wantAttempts := []string{"attempt=1; max=3", "attempt=2; max=3", "attempt=3; max=3"}
	for i, want := range wantAttempts {
		if attemptHeaders[i] != want {
			t.Errorf("attempt %d: X-Sdk-Request = %q, want %q", i+1, attemptHeaders[i], want)
		}
	}
}

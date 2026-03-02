package request

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
)

func newTestRequest(handlers Handlers) *Request {
	cfg := volcengine.Config{
		Region:      volcengine.String("cn-beijing"),
		Credentials: credentials.NewStaticCredentials("AKID", "SECRET", ""),
	}

	op := &Operation{
		Name:       "TestAction",
		HTTPMethod: http.MethodGet,
		HTTPPath:   "",
	}

	return New(cfg, metadata.ClientInfo{
		ServiceName:   "test_service",
		ServiceID:     "test_service",
		SigningName:   "test_service",
		SigningRegion: "cn-beijing",
		Endpoint:      "https://test.cn-beijing.volcengineapi.com",
		APIVersion:    "2022-01-01",
	}, handlers, nil, op, nil, nil)
}

func TestPresignRequest_ValidExpire(t *testing.T) {
	signCalled := false

	var handlers Handlers
	handlers.Sign.PushBack(func(req *Request) {
		signCalled = true
		if !req.IsPresigned() {
			t.Error("IsPresigned() should return true during PresignRequest")
		}
		if req.ExpireTime != 15*time.Minute {
			t.Errorf("ExpireTime = %v, want %v", req.ExpireTime, 15*time.Minute)
		}
		// Simulate URL signing by appending signature to query
		q := req.HTTPRequest.URL.Query()
		q.Set("X-Signature", "test-signature")
		req.HTTPRequest.URL.RawQuery = q.Encode()
	})

	req := newTestRequest(handlers)

	q := req.HTTPRequest.URL.Query()
	q.Set("Action", "TestAction")
	q.Set("Version", "2022-01-01")
	req.HTTPRequest.URL.RawQuery = q.Encode()

	signedURL, headers, err := req.PresignRequest(15 * time.Minute)
	if err != nil {
		t.Fatalf("PresignRequest() unexpected error: %v", err)
	}

	if !signCalled {
		t.Error("Sign handler was not called")
	}

	if signedURL == "" {
		t.Error("PresignRequest() returned empty URL")
	}

	if !strings.Contains(signedURL, "X-Signature=test-signature") {
		t.Error("PresignRequest() URL should contain signature from Sign handler")
	}

	if !strings.Contains(signedURL, "Action=TestAction") {
		t.Error("PresignRequest() URL should contain original query params")
	}

	if !strings.HasPrefix(signedURL, "https://test.cn-beijing.volcengineapi.com") {
		t.Errorf("PresignRequest() URL has unexpected prefix: %s", signedURL)
	}

	// headers can be nil when no signed headers are set
	_ = headers
}

func TestPresignRequest_ZeroExpire(t *testing.T) {
	var handlers Handlers
	req := newTestRequest(handlers)

	_, _, err := req.PresignRequest(0)
	if err == nil {
		t.Fatal("PresignRequest(0) should return error")
	}

	if !strings.Contains(err.Error(), "expire duration greater than 0") {
		t.Errorf("PresignRequest(0) error = %v, want InvalidPresignExpire", err)
	}
}

func TestPresignRequest_NegativeExpire(t *testing.T) {
	var handlers Handlers
	req := newTestRequest(handlers)

	_, _, err := req.PresignRequest(-1 * time.Second)
	if err == nil {
		t.Fatal("PresignRequest(-1s) should return error")
	}
}

func TestPresignRequest_DoesNotMutateOriginal(t *testing.T) {
	var handlers Handlers
	handlers.Sign.PushBack(func(req *Request) {
		q := req.HTTPRequest.URL.Query()
		q.Set("X-Signature", "sig")
		req.HTTPRequest.URL.RawQuery = q.Encode()
	})

	req := newTestRequest(handlers)

	originalExpireTime := req.ExpireTime

	_, _, err := req.PresignRequest(10 * time.Minute)
	if err != nil {
		t.Fatalf("PresignRequest() unexpected error: %v", err)
	}

	// Original request should not be mutated (PresignRequest copies)
	if req.ExpireTime != originalExpireTime {
		t.Errorf("Original request ExpireTime was mutated: got %v, want %v", req.ExpireTime, originalExpireTime)
	}
}

func TestPresignRequest_SignHandlerError(t *testing.T) {
	var handlers Handlers
	handlers.Sign.PushBack(func(req *Request) {
		req.Error = fmt.Errorf("sign error")
	})

	req := newTestRequest(handlers)

	_, _, err := req.PresignRequest(15 * time.Minute)
	if err == nil {
		t.Fatal("PresignRequest() should return error when Sign handler fails")
	}

	if !strings.Contains(err.Error(), "sign error") {
		t.Errorf("PresignRequest() error = %v, want 'sign error'", err)
	}
}

func TestPresignRequest_BeforePresignFn(t *testing.T) {
	beforeCalled := false

	var handlers Handlers
	handlers.Sign.PushBack(func(req *Request) {})

	req := newTestRequest(handlers)
	req.Operation.BeforePresignFn = func(r *Request) error {
		beforeCalled = true
		q := r.HTTPRequest.URL.Query()
		q.Set("CustomParam", "injected")
		r.HTTPRequest.URL.RawQuery = q.Encode()
		return nil
	}

	signedURL, _, err := req.PresignRequest(15 * time.Minute)
	if err != nil {
		t.Fatalf("PresignRequest() unexpected error: %v", err)
	}

	if !beforeCalled {
		t.Error("BeforePresignFn was not called")
	}

	if !strings.Contains(signedURL, "CustomParam=injected") {
		t.Error("PresignRequest() URL should contain param added by BeforePresignFn")
	}
}

func TestPresignRequest_BeforePresignFnError(t *testing.T) {
	var handlers Handlers
	req := newTestRequest(handlers)
	req.Operation.BeforePresignFn = func(r *Request) error {
		return fmt.Errorf("before presign error")
	}

	_, _, err := req.PresignRequest(15 * time.Minute)
	if err == nil {
		t.Fatal("PresignRequest() should return error when BeforePresignFn fails")
	}

	if !strings.Contains(err.Error(), "before presign error") {
		t.Errorf("PresignRequest() error = %v, want 'before presign error'", err)
	}
}

func TestIsPresigned(t *testing.T) {
	var handlers Handlers
	req := newTestRequest(handlers)

	if req.IsPresigned() {
		t.Error("IsPresigned() should return false for new request")
	}

	req.ExpireTime = 15 * time.Minute
	if !req.IsPresigned() {
		t.Error("IsPresigned() should return true when ExpireTime > 0")
	}
}

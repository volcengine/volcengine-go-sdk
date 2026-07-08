package vefaas_test

// Confirms the retry attempt headers (X-Sdk-Invocation-Id, X-Sdk-Request) are
// present before signing runs, so they end up covered by the request's
// signature (Authorization: ...SignedHeaders=...) instead of being appended
// after signing.

import (
	"strings"
	"testing"

	"github.com/volcengine/volcengine-go-sdk/service/vefaas"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func TestCreateSandboxRequestSignsRetryHeaders(t *testing.T) {
	sess, err := session.NewSession(volcengine.NewConfig().
		WithRegion("cn-beijing").
		WithEndpoint("https://example.com").
		WithCredentials(credentials.NewStaticCredentials("test-ak", "test-sk", "")))
	if err != nil {
		t.Fatalf("failed to create session: %v", err)
	}

	client := vefaas.New(sess)
	req, _ := client.CreateSandboxRequest(&vefaas.CreateSandboxInput{
		FunctionId: volcengine.String("fn-test"),
	})

	if err := req.Sign(); err != nil {
		t.Fatalf("failed to sign request: %v", err)
	}

	if req.HTTPRequest.Header.Get("X-Sdk-Invocation-Id") == "" {
		t.Fatal("X-Sdk-Invocation-Id header missing before signing")
	}
	if req.HTTPRequest.Header.Get("X-Sdk-Request") == "" {
		t.Fatal("X-Sdk-Request header missing before signing")
	}

	auth := req.HTTPRequest.Header.Get("Authorization")
	if auth == "" {
		t.Fatal("Authorization header missing after Sign()")
	}

	idx := strings.Index(auth, "SignedHeaders=")
	if idx == -1 {
		t.Fatalf("Authorization header has no SignedHeaders: %s", auth)
	}
	signedHeaders := auth[idx:]
	if end := strings.Index(signedHeaders, ","); end != -1 {
		signedHeaders = signedHeaders[:end]
	}

	for _, want := range []string{"x-sdk-invocation-id", "x-sdk-request"} {
		if !strings.Contains(signedHeaders, want) {
			t.Errorf("SignedHeaders = %q, want it to contain %q", signedHeaders, want)
		}
	}
}

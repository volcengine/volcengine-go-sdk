package credentials

import (
	"encoding/pem"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// Clone initializes the base transport's h2 ALPN configuration. The cloned
// transport must install a matching HTTP/2 handler even with TLSClientConfig.
func TestSTSProxyTransportNegotiatesHTTP2(t *testing.T) {
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor != 2 {
			t.Errorf("protocol = %s, want HTTP/2", r.Proto)
		}
		_, _ = io.WriteString(w, "ok")
	}))
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()
	ca, err := ioutil.TempFile("", "sts-http2-ca-*.pem")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(ca.Name())
	if err = pem.Encode(ca, &pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}); err != nil {
		ca.Close()
		t.Fatal(err)
	}
	if err = ca.Close(); err != nil {
		t.Fatal(err)
	}
	transport, err := buildProxyTransport(ProxyConfig{CAFile: ca.Name(), base: &http.Transport{}})
	if err != nil {
		t.Fatal(err)
	}
	defer transport.CloseIdleConnections()
	client := &http.Client{Transport: transport, Timeout: 5 * time.Second}
	response, err := client.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if response.ProtoMajor != 2 {
		t.Fatalf("response protocol = %s, want HTTP/2", response.Proto)
	}
}

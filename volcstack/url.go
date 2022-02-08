package volcstack

// This File is modify from https://github.com/aws/aws-sdk-go/blob/main/aws/url.go

import "net/url"

// URLHostname will extract the Hostname without port from the URL value.
//
// Wrapper of net/url#URL.Hostname for backwards Go version compatibility.
func URLHostname(url *url.URL) string {
	return url.Hostname()
}

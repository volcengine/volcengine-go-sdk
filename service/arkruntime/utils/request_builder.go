package utils

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/apiform"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/apiquery"
)

type RequestBuilder interface {
	Build(ctx context.Context, method, url string, body interface{}, header http.Header) (*http.Request, error)
}

type HTTPRequestBuilder struct {
	marshaller Marshaller
}

func NewRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		marshaller: &JSONMarshaller{},
	}
}

func (b *HTTPRequestBuilder) Build(
	ctx context.Context,
	method string,
	url string,
	body interface{},
	header http.Header,
) (req *http.Request, err error) {
	var bodyReader io.Reader
	contentType := "application/json"

	if body != nil {
		if v, ok := body.(io.Reader); ok { // already marshalled
			bodyReader = v
		} else if v, ok := body.([]byte); ok { // already marshalled
			bodyReader = bytes.NewBuffer(v)
		} else if v, ok := body.(apiform.Marshaler); ok { // multipart form
			var (
				content []byte
				err     error
			)
			content, contentType, err = v.MarshalMultipart()
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewBuffer(content)
		} else if v, ok := body.(apiquery.Queryer); ok { // url query
			q, err := v.URLQuery()
			if err != nil {
				return nil, err
			}
			params := q.Encode()
			if params != "" {
				url = url + "?" + params
			}
		} else { // json
			var reqBytes []byte
			reqBytes, err = b.marshaller.Marshal(body)
			if err != nil {
				return
			}
			bodyReader = bytes.NewBuffer(reqBytes)
		}
	}

	req, err = http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return
	}
	if header != nil {
		req.Header = header
	}
	if bodyReader != nil {
		req.Header.Set("Content-Type", contentType)
	}

	return
}

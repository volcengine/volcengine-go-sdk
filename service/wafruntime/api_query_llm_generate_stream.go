package wafruntime

import (
	"bufio"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/wafruntime/utils"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
	"io"
	"net/http"
	"net/url"
	"time"
)

func copyHTTPRequest(r *http.Request, body io.ReadCloser) *http.Request {
	req := new(http.Request)
	*req = *r
	req.URL = &url.URL{}
	*req.URL = *r.URL
	req.Body = body

	req.Header = http.Header{}
	for k, v := range r.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}
	return req
}

func (c *WAFRuntime) handleStreamResponse(r *request.Request) {
	if r.DataFilled() {
		r.Data.(*QueryLLMGenerateStreamOutput).Reader = bufio.NewReader(r.HTTPResponse.Body)
		r.Data.(*QueryLLMGenerateStreamOutput).IsFinished = false
		r.Data.(*QueryLLMGenerateStreamOutput).Response = r.HTTPResponse
	}
}

func (c *WAFRuntime) sendRequest(r *request.Request) (sendErr error) {
	defer r.Handlers.CompleteAttempt.Run(r)

	r.Retryable = nil
	r.Handlers.Send.Run(r)
	if r.Error != nil {
		return r.Error
	}

	r.Handlers.UnmarshalMeta.Run(r)
	r.Handlers.ValidateResponse.Run(r)
	if r.Error != nil {
		r.Handlers.UnmarshalError.Run(r)
		return r.Error
	}

	c.handleStreamResponse(r)
	if r.Error != nil {
		r.Handlers.UnmarshalError.Run(r)
		return r.Error
	}

	return nil
}

func (c *WAFRuntime) prepareRetry(r *request.Request) error {
	if r.Config.LogLevel.Matches(volcengine.LogDebugWithRequestRetries) {
		r.Config.Logger.Log(fmt.Sprintf("DEBUG: Retrying Request %s/%s, attempt %d",
			r.ClientInfo.ServiceName, r.Operation.Name, r.RetryCount))
	}

	// The previous http.Request will have a reference to the r.Body
	// and the HTTP Client's Transport may still be reading from
	// the request's volcenginebody even though the Client's Do returned.
	r.HTTPRequest = copyHTTPRequest(r.HTTPRequest, nil)
	r.ResetBody()
	if err := r.Error; err != nil {
		return volcengineerr.New(request.ErrCodeSerialization,
			"failed to prepare volcenginebody for retry", err)
	}

	// Closing response volcenginebody to ensure that no response volcenginebody is leaked
	// between retry attempts.
	if r.HTTPResponse != nil && r.HTTPResponse.Body != nil {
		if err := r.HTTPResponse.Body.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (c *WAFRuntime) send(r *request.Request) error {
	interceptorMapping := make(map[int]interface{})
	defer func() {
		// Regardless of success or failure of the request trigger the Complete
		// request handlers.
		r.Handlers.Complete.Run(r)
		for index, interceptor := range r.Config.Interceptors {
			if interceptor.After != nil {
				interceptor.After(r.MergeRequestInfo(), interceptorMapping[index])
			}
		}
	}()

	if err := r.Error; err != nil {
		return err
	}

	for {
		r.Error = nil
		r.AttemptTime = time.Now()

		if err := r.Sign(); err != nil {
			return err
		}

		for index, interceptor := range r.Config.Interceptors {
			if interceptor.Before != nil {
				interceptorMapping[index] = interceptor.Before(r.MergeRequestInfo())
			} else {
				interceptorMapping[index] = nil
			}
		}

		if err := c.sendRequest(r); err == nil {
			return nil
		}
		r.Handlers.Retry.Run(r)
		r.Handlers.AfterRetry.Run(r)

		if r.Error != nil || !volcengine.BoolValue(r.Retryable) {
			return r.Error
		}

		if err := c.prepareRetry(r); err != nil {
			r.Error = err
			return err
		}
	}
}

const opQueryLLMGenerateStream = "QueryLLMGenerate"

// QueryLLMGenerateStreamRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryLLMGenerateStream operation. The "output" return
// value will be populated with the QueryLLMGenerateStreamCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryLLMGenerateStreamCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryLLMGenerateStreamCommon Send returns without error.
//
// See QueryLLMGenerateStream for more information on using the QueryLLMGenerateStream
// API call, and error handling.
//
//	// Example sending a request using the QueryLLMGenerateStreamRequest method.
//	req, resp := client.QueryLLMGenerateStreamRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *WAFRuntime) QueryLLMGenerateStreamRequest(input *QueryLLMGenerateStreamInput) (req *request.Request, output *QueryLLMGenerateStreamOutput) {
	op := &request.Operation{
		Name:       opQueryLLMGenerateStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &QueryLLMGenerateStreamInput{}
	}

	output = &QueryLLMGenerateStreamOutput{
		StreamReader: &utils.StreamReader{},
	}
	req = c.newRequest(op, input, output)

	err := input.Validate()
	if err != nil {
		req.Error = err
		return
	}

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	return
}

// QueryLLMGenerateStream API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation QueryLLMGenerateStream for usage and error information.
func (c *WAFRuntime) QueryLLMGenerateStream(input *QueryLLMGenerateStreamInput) (*QueryLLMGenerateStreamOutput, error) {
	req, out := c.QueryLLMGenerateStreamRequest(input)
	return out, c.send(req)
}

// QueryLLMGenerateStreamWithContext is the same as QueryLLMGenerateStream with the addition of
// the ability to pass a context and additional request options.
//
// See QueryLLMGenerateStream for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAFRuntime) QueryLLMGenerateStreamWithContext(ctx volcengine.Context, input *QueryLLMGenerateStreamInput, opts ...request.Option) (*QueryLLMGenerateStreamOutput, error) {
	req, out := c.QueryLLMGenerateStreamRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, c.send(req)
}

type QueryLLMGenerateStreamInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// MsgID is a required field
	MsgID *string `type:"string" json:",omitempty" required:"true"`

	UseStream *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s QueryLLMGenerateStreamInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryLLMGenerateStreamInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryLLMGenerateStreamInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "QueryLLMGenerateStreamInput"}
	if s.MsgID == nil {
		invalidParams.Add(request.NewErrParamRequired("MsgID"))
	}

	if s.UseStream != nil && !*s.UseStream {
		invalidParams.Add(request.NewErrParamFormat("UseStream", "Stream API Should be set UseStream as true", fmt.Sprintf("%v", *s.UseStream)))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}

	return nil
}

// SetMsgID sets the MsgID field's value.
func (s *QueryLLMGenerateStreamInput) SetMsgID(v string) *QueryLLMGenerateStreamInput {
	s.MsgID = &v
	return s
}

// SetUseStream sets the UseStream field's value.
func (s *QueryLLMGenerateStreamInput) SetUseStream(v bool) *QueryLLMGenerateStreamInput {
	s.UseStream = &v
	return s
}

type QueryLLMGenerateStreamOutput struct {
	*utils.StreamReader
}

package volcstackquery

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackerr"
)

// UnmarshalErrorHandler is a name request handler to unmarshal request errors
var UnmarshalErrorHandler = request.NamedHandler{Name: "kscsdk.volcstackquery.UnmarshalError", Fn: UnmarshalError}

// UnmarshalError unmarshals an error response for an VOLCSTACK Query service.
func UnmarshalError(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	resp := response.VolcstackResponse{}
	if r.DataFilled() {
		body, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			fmt.Printf("read volcstackbody err, %v\n", err)
			r.Error = err
			return
		}

		if err = json.Unmarshal(body, &resp); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			r.Error = err
			return
		}

		if resp.ResponseMetadata == nil {
			simple := response.VolcstackSimpleError{}
			if err = json.Unmarshal(body, &simple); err != nil {
				fmt.Printf("Unmarshal err, %v\n", err)
				r.Error = err
				return
			}
			resp.ResponseMetadata = &response.ResponseMetadata{
				Error: &response.Error{
					Code:    simple.ErrorCode,
					Message: simple.Message,
				},
			}
		}

		r.Error = volcstackerr.NewRequestFailure(
			volcstackerr.New(resp.ResponseMetadata.Error.Code, resp.ResponseMetadata.Error.Message, nil),
			r.HTTPResponse.StatusCode,
			resp.ResponseMetadata.RequestId,
			r.Config.SimpleError,
		)
		return
	} else {
		r.Error = volcstackerr.NewRequestFailure(
			volcstackerr.New("ServiceUnavailableException", "service is unavailable", nil),
			r.HTTPResponse.StatusCode,
			r.RequestID,
		)
		return
	}
}

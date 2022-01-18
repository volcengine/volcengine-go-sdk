package volcstackquery

import (
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/request"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/volcstackerr"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type jsonErrorResponse struct {
	ResponseMetadata Metadata `json:"ResponseMetadata"`
}

type Metadata struct {
	Error     Error  `json:"Error"`
	RequestID string `json:"RequestID"`
}

type Error struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

// UnmarshalErrorHandler is a name request handler to unmarshal request errors
var UnmarshalErrorHandler = request.NamedHandler{Name: "kscsdk.volcstackquery.UnmarshalError", Fn: UnmarshalError}

// UnmarshalError unmarshals an error response for an AWS Query service.
func UnmarshalError(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	resp := jsonErrorResponse{}
	if r.DataFilled() {
		body, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			fmt.Printf("read volcstackbody err, %v\n", err)
			return
		}
		if err = json.Unmarshal(body, &resp); err != nil {
			fmt.Printf("Unmarshal err, %v\n", err)
			return
		}
		r.Error = volcstackerr.NewRequestFailure(
			volcstackerr.New(resp.ResponseMetadata.Error.Code, resp.ResponseMetadata.Error.Message, nil),
			r.HTTPResponse.StatusCode,
			resp.ResponseMetadata.RequestID,
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

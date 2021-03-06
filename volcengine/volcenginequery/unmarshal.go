package volcenginequery

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

// UnmarshalHandler is a named request handler for unmarshaling volcenginequery protocol requests
var UnmarshalHandler = request.NamedHandler{Name: "volcenginesdk.volcenginequery.Unmarshal", Fn: Unmarshal}

// UnmarshalMetaHandler is a named request handler for unmarshaling volcenginequery protocol request metadata
var UnmarshalMetaHandler = request.NamedHandler{Name: "volcenginesdk.volcenginequery.UnmarshalMeta", Fn: UnmarshalMeta}

// Unmarshal unmarshals a response for an VOLCSTACK Query service.
func Unmarshal(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		body, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			fmt.Printf("read volcenginebody err, %v\n", err)
			r.Error = err
			return
		}

		if reflect.TypeOf(r.Data) == reflect.TypeOf(&map[string]interface{}{}) {
			if err = json.Unmarshal(body, &r.Data); err != nil {
				fmt.Printf("Unmarshal err, %v\n", err)
				r.Error = err
				return
			}
			var info interface{}

			ptr := r.Data.(*map[string]interface{})
			info, err = volcengineutil.ObtainSdkValue("ResponseMetadata.Error.Code", *ptr)
			if err != nil {
				r.Error = err
				return
			}
			if info != nil {
				if processBodyError(r, &response.VolcengineResponse{}, body) {
					return
				}
			}

		} else {
			volcengineResponse := response.VolcengineResponse{}
			if processBodyError(r, &volcengineResponse, body) {
				return
			}

			if _, ok := reflect.TypeOf(r.Data).Elem().FieldByName("Metadata"); ok {
				r.Metadata = *(volcengineResponse.ResponseMetadata)
				reflect.ValueOf(r.Data).Elem().FieldByName("Metadata").Set(reflect.ValueOf(volcengineResponse.ResponseMetadata))
			}

			var b []byte
			if b, err = json.Marshal(volcengineResponse.Result); err != nil {
				fmt.Printf("Unmarshal err, %v\n", err)
				r.Error = err
				return
			}
			if err = json.Unmarshal(b, &r.Data); err != nil {
				fmt.Printf("Unmarshal err, %v\n", err)
				r.Error = err
				return
			}
		}

	}
}

// UnmarshalMeta unmarshals header response values for an VOLCSTACK Query service.
func UnmarshalMeta(r *request.Request) {

}

func processBodyError(r *request.Request, volcengineResponse *response.VolcengineResponse, body []byte) bool {
	if err := json.Unmarshal(body, &volcengineResponse); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		r.Error = err
		return true
	}
	if volcengineResponse.ResponseMetadata.Error != nil && volcengineResponse.ResponseMetadata.Error.Code != "" {
		r.Error = volcengineerr.NewRequestFailure(
			volcengineerr.New(volcengineResponse.ResponseMetadata.Error.Code, volcengineResponse.ResponseMetadata.Error.Message, nil),
			http.StatusBadRequest,
			volcengineResponse.ResponseMetadata.RequestId,
		)
		processUnmarshalError(unmarshalErrorInfo{
			Request:  r,
			Response: volcengineResponse,
			Body:     body,
			Err:      r.Error,
		})
		return true
	}
	return false
}

package volcstackquery

import (
	"encoding/json"
	"fmt"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackerr"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
	"io/ioutil"
	"net/http"
	"reflect"
)

type VolcstackResponse struct {
	ResponseMetadata ResponseMetadata
	Result           interface{}
}

type ResponseMetadata struct {
	RequestId string
	Action    string
	Version   string
	Service   string
	Region    string
	Error     *Error
}

type Error struct {
	CodeN   int
	Code    string
	Message string
}

// UnmarshalHandler is a named request handler for unmarshaling volcstackquery protocol requests
var UnmarshalHandler = request.NamedHandler{Name: "volcstacksdk.volcstackquery.Unmarshal", Fn: Unmarshal}

// UnmarshalMetaHandler is a named request handler for unmarshaling volcstackquery protocol request metadata
var UnmarshalMetaHandler = request.NamedHandler{Name: "volcstacksdk.volcstackquery.UnmarshalMeta", Fn: UnmarshalMeta}

// Unmarshal unmarshals a response for an AWS Query service.
func Unmarshal(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		body, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			fmt.Printf("read volcstackbody err, %v\n", err)
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
			info, err = volcstackutil.ObtainSdkValue("ResponseMetadata.Error.Code", *ptr)
			if err != nil {
				r.Error = err
				return
			}
			if info != nil {
				if processBodyError(r, &VolcstackResponse{}, body) {
					return
				}
			}

		} else {
			volcstackResponse := VolcstackResponse{}
			if processBodyError(r, &volcstackResponse, body) {
				return
			}
			var b []byte
			if b, err = json.Marshal(volcstackResponse.Result); err != nil {
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

// UnmarshalMeta unmarshals header response values for an AWS Query service.
func UnmarshalMeta(r *request.Request) {

}

func processBodyError(r *request.Request, volcstackResponse *VolcstackResponse, body []byte) bool {
	if err := json.Unmarshal(body, &volcstackResponse); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		r.Error = err
		return true
	}
	if volcstackResponse.ResponseMetadata.Error != nil && volcstackResponse.ResponseMetadata.Error.Code != "" {
		r.Error = volcstackerr.NewRequestFailure(
			volcstackerr.New(volcstackResponse.ResponseMetadata.Error.Code, volcstackResponse.ResponseMetadata.Error.Message, nil),
			http.StatusBadRequest,
			volcstackResponse.ResponseMetadata.RequestId,
		)
		return true
	}
	return false
}

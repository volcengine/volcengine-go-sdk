package volcstackquery

import (
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/request"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

type VolcstackResponse struct {
	ResponseMetadata interface{}
	Result           interface{}
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
			return
		}

		if reflect.TypeOf(r.Data) == reflect.TypeOf(&map[string]interface{}{}) {
			if err = json.Unmarshal(body, &r.Data); err != nil {
				fmt.Printf("Unmarshal err, %v\n", err)
				return
			}
		} else {
			volcstackResponse := VolcstackResponse{}
			if err = json.Unmarshal(body, &volcstackResponse); err != nil {
				fmt.Printf("Unmarshal err, %v\n", err)
				return
			}
			var b []byte
			if b, err = json.Marshal(volcstackResponse.Result); err != nil {
				fmt.Printf("Unmarshal err, %v\n", err)
				return
			}
			if err = json.Unmarshal(b, &r.Data); err != nil {
				fmt.Printf("Unmarshal err, %v\n", err)
				return
			}
		}

	}
}

// UnmarshalMeta unmarshals header response values for an AWS Query service.
func UnmarshalMeta(r *request.Request) {
	//r.RequestID = r.HTTPResponse.Header.Get("X-Amzn-Requestid")
}

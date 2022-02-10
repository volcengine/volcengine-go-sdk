package query

// This File is modify from https://github.com/aws/aws-sdk-go/blob/main/private/protocol/query/unmarshal.go

import (
	"encoding/xml"

	"github.com/volcengine/volcstack-go-sdk/private/protocol/xml/xmlutil"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackerr"
)

// UnmarshalHandler is a named request handler for unmarshaling volcstackquery protocol requests
var UnmarshalHandler = request.NamedHandler{Name: "awssdk.volcstackquery.Unmarshal", Fn: Unmarshal}

// UnmarshalMetaHandler is a named request handler for unmarshaling volcstackquery protocol request metadata
var UnmarshalMetaHandler = request.NamedHandler{Name: "awssdk.volcstackquery.UnmarshalMeta", Fn: UnmarshalMeta}

// Unmarshal unmarshals a response for an AWS Query service.
func Unmarshal(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		decoder := xml.NewDecoder(r.HTTPResponse.Body)
		err := xmlutil.UnmarshalXML(r.Data, decoder, r.Operation.Name+"Result")
		if err != nil {
			r.Error = volcstackerr.NewRequestFailure(
				volcstackerr.New(request.ErrCodeSerialization, "failed decoding Query response", err),
				r.HTTPResponse.StatusCode,
				r.RequestID,
			)
			return
		}
	}
}

// UnmarshalMeta unmarshals header response values for an AWS Query service.
func UnmarshalMeta(r *request.Request) {
	r.RequestID = r.HTTPResponse.Header.Get("X-Top-Requestid")
}

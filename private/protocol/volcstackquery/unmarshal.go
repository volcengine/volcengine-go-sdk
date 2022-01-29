package volcstackquery

//go:generate go run -tags codegen ../../../private/model/cli/gen-protocol-tests ../../../models/protocol_tests/output/ec2.json unmarshal_test.go

// This File is modify from https://github.com/aws/aws-sdk-go/blob/main/private/protocol/query/unmarshal.go

import (
	"encoding/xml"

	"code.byted.org/iaasng/volcstack-go-sdk/private/protocol/xml/xmlutil"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/request"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/volcstackerr"
)

// UnmarshalHandler is a named request handler for unmarshaling ec2query protocol requests
var UnmarshalHandler = request.NamedHandler{Name: "volcstacksdk.volcstackquery.Unmarshal", Fn: Unmarshal}

// UnmarshalMetaHandler is a named request handler for unmarshaling ec2query protocol request metadata
var UnmarshalMetaHandler = request.NamedHandler{Name: "volcstacksdk.volcstackquery.UnmarshalMeta", Fn: UnmarshalMeta}

// UnmarshalErrorHandler is a named request handler for unmarshaling ec2query protocol request errors
var UnmarshalErrorHandler = request.NamedHandler{Name: "volcstacksdk.volcstackquery.UnmarshalError", Fn: UnmarshalError}

// Unmarshal unmarshals a response body for the EC2 protocol.
func Unmarshal(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		decoder := xml.NewDecoder(r.HTTPResponse.Body)
		err := xmlutil.UnmarshalXML(r.Data, decoder, "")
		if err != nil {
			r.Error = volcstackerr.NewRequestFailure(
				volcstackerr.New(request.ErrCodeSerialization,
					"failed decoding EC2 Query response", err),
				r.HTTPResponse.StatusCode,
				r.RequestID,
			)
			return
		}
	}
}

// UnmarshalMeta unmarshals response headers for the EC2 protocol.
func UnmarshalMeta(r *request.Request) {
	r.RequestID = r.HTTPResponse.Header.Get("X-Amzn-Requestid")
	if r.RequestID == "" {
		// Alternative version of request id in the header
		r.RequestID = r.HTTPResponse.Header.Get("X-Amz-Request-Id")
	}
}

type xmlErrorResponse struct {
	XMLName   xml.Name `xml:"Response"`
	Code      string   `xml:"Errors>Error>Code"`
	Message   string   `xml:"Errors>Error>Message"`
	RequestID string   `xml:"RequestID"`
}

// UnmarshalError unmarshals a response error for the EC2 protocol.
func UnmarshalError(r *request.Request) {
	defer r.HTTPResponse.Body.Close()

	var respErr xmlErrorResponse
	err := xmlutil.UnmarshalXMLError(&respErr, r.HTTPResponse.Body)
	if err != nil {
		r.Error = volcstackerr.NewRequestFailure(
			volcstackerr.New(request.ErrCodeSerialization,
				"failed to unmarshal error message", err),
			r.HTTPResponse.StatusCode,
			r.RequestID,
		)
		return
	}

	r.Error = volcstackerr.NewRequestFailure(
		volcstackerr.New(respErr.Code, respErr.Message, nil),
		r.HTTPResponse.StatusCode,
		respErr.RequestID,
	)
}

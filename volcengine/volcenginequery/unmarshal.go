package volcenginequery

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/special"
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
		body, err := r.ReadResponseBody()
		if err != nil {
			fmt.Printf("read volcenginebody err, %v\n", err)
			r.Error = newResponseError(unmarshalErrorInfo{
				Request: r,
				Body:    body,
			}, request.ErrCodeRead, "failed to read response", err)
			return
		}

		var forceJsonNumberDecoder bool

		if r.Config.ForceJsonNumberDecode != nil {
			forceJsonNumberDecoder = r.Config.ForceJsonNumberDecode(r.Context(), r.MergeRequestInfo())
		}

		if reflect.TypeOf(r.Data) == reflect.TypeOf(&map[string]interface{}{}) {
			originalData := r.Data
			//如果使用map返回 发现精度丢失了 请设置强制JsonNumber 注意返回的整型会动float64->int64
			if err = json.Unmarshal(body, &r.Data); err != nil || forceJsonNumberDecoder {
				//try next
				decoder := json.NewDecoder(bytes.NewReader(body))
				decoder.UseNumber()
				if err = decoder.Decode(&r.Data); err != nil {
					fmt.Printf("Unmarshal err, %v\n", err)
					r.Error = newResponseError(unmarshalErrorInfo{
						Request: r,
						Body:    body,
					}, request.ErrCodeSerialization, "failed to decode response", err)
					return
				}
			}
			var info interface{}

			ptr, ok := r.Data.(*map[string]interface{})
			if !ok || ptr == nil {
				// Keep the output target available for retries and interceptors.
				r.Data = originalData
				r.Error = newResponseError(unmarshalErrorInfo{
					Request: r,
					Body:    body,
				}, "UnknownError", "non-standard response body: null response", nil)
				return
			}
			info, err = volcengineutil.ObtainSdkValue("ResponseMetadata.Error.Code", *ptr)
			if err != nil {
				r.Error = newResponseError(unmarshalErrorInfo{
					Request: r,
					Body:    body,
				}, request.ErrCodeSerialization, "failed to access response error code", err)
				return
			}
			if info != nil {
				if processBodyError(r, &response.VolcengineResponse{}, body, forceJsonNumberDecoder) {
					return
				}
			}

		} else {
			volcengineResponse := response.VolcengineResponse{}
			if processBodyError(r, &volcengineResponse, body, forceJsonNumberDecoder) {
				return
			}

			if _, ok := reflect.TypeOf(r.Data).Elem().FieldByName("Metadata"); ok {
				if volcengineResponse.ResponseMetadata != nil {
					volcengineResponse.ResponseMetadata.HTTPCode = r.HTTPResponse.StatusCode
					r.Metadata = *(volcengineResponse.ResponseMetadata)
				}
				reflect.ValueOf(r.Data).Elem().FieldByName("Metadata").Set(reflect.ValueOf(volcengineResponse.ResponseMetadata))
			}

			var (
				b      []byte
				source interface{}
			)

			if r.Config.CustomerUnmarshalData != nil {
				source = r.Config.CustomerUnmarshalData(r.Context(), r.MergeRequestInfo(), volcengineResponse)
			} else {
				if sp, ok := special.ResponseSpecialMapping()[r.ClientInfo.ServiceName]; ok {
					source = sp(volcengineResponse, r.Data)
				} else {
					source = volcengineResponse.Result
				}
			}

			if b, err = json.Marshal(source); err != nil {
				fmt.Printf("Unmarshal err, %v\n", err)
				r.Error = newResponseError(unmarshalErrorInfo{
					Request:  r,
					Response: &volcengineResponse,
					Body:     body,
				}, request.ErrCodeSerialization, "failed to encode response result", err)
				return
			}
			if err = json.Unmarshal(b, &r.Data); err != nil || forceJsonNumberDecoder {
				decoder := json.NewDecoder(bytes.NewReader(b))
				decoder.UseNumber()
				if err = decoder.Decode(&r.Data); err != nil {
					fmt.Printf("Unmarshal err, %v\n", err)
					r.Error = newResponseError(unmarshalErrorInfo{
						Request:  r,
						Response: &volcengineResponse,
						Body:     body,
					}, request.ErrCodeSerialization, "failed to decode response result", err)
					return
				}
			}
		}

	}
}

// UnmarshalMeta unmarshals header response values for an VOLCSTACK Query service.
func UnmarshalMeta(r *request.Request) {

}

func processBodyError(r *request.Request, volcengineResponse *response.VolcengineResponse, body []byte, forceJsonNumberDecoder bool) bool {
	//防止精度问题 第一次转换 无视 保持原body内容不会失去精度
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()
	if err := decoder.Decode(&volcengineResponse); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		r.Error = newResponseError(unmarshalErrorInfo{
			Request:  r,
			Response: volcengineResponse,
			Body:     body,
		}, request.ErrCodeSerialization, "failed to decode response", err)
		return true
	}

	// JSON null clears the decoded pointer; stop before reading its fields.
	if volcengineResponse == nil {
		r.Error = newResponseError(unmarshalErrorInfo{
			Request:  r,
			Response: volcengineResponse,
			Body:     body,
		}, "UnknownError", "non-standard response body: null response", nil)
		return true
	}
	if volcengineResponse.ResponseMetadata == nil {
		return false
	}
	if volcengineResponse.ResponseMetadata.Error != nil && volcengineResponse.ResponseMetadata.Error.Code != "" {
		processUnmarshalError(unmarshalErrorInfo{
			Request:  r,
			Response: volcengineResponse,
			Body:     body,
		})
		return true
	}
	return false
}

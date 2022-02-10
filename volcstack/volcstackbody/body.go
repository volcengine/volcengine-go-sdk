package volcstackbody

import (
	"encoding/json"
	"github.com/volcengine/volcstack-go-sdk/private/protocol/query/queryutil"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackerr"
	"net/url"
	"reflect"
	"strings"
)

func BodyParam(r *request.Request) {
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.ClientInfo.APIVersion},
	}
	method := strings.ToUpper(r.HTTPRequest.Method)
	if reflect.TypeOf(r.Params) == reflect.TypeOf(&map[string]interface{}{}) {
		m := *(r.Params).(*map[string]interface{})
		for k, v := range m {
			body.Add(k, v.(string))
		}
	} else if err := queryutil.Parse(body, r.Params, false); err != nil {
		r.Error = volcstackerr.New("SerializationError", "failed encoding Query request", err)
		return
	}

	v := r.HTTPRequest.Header.Get("Content-Type")
	if method == "POST" {
		if len(v) > 0 && strings.Contains(strings.ToLower(v), "application/json") {
			BodyJson(r)
			return
		}
	}

	r.HTTPRequest.Header.Add("Accept", "application/json")

	if method == "GET" {
		r.HTTPRequest.URL.RawQuery = body.Encode()
	} else if method == "POST" {
		r.HTTPRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		r.SetBufferBody([]byte(body.Encode()))
	} else if method == "PUT" {
		r.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
		BodyJson(r)
		return
	} else {
		r.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
		r.SetBufferBody([]byte(body.Encode()))
	}
}

func BodyJson(r *request.Request) {
	method := strings.ToUpper(r.HTTPRequest.Method)
	if v := r.HTTPRequest.Header.Get("Content-Type"); len(v) == 0 {
		r.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if v := r.HTTPRequest.Header.Get("Content-Type"); !strings.Contains(strings.ToLower(v), "application/json") || method == "GET" {
		BodyParam(r)
		return
	}
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.ClientInfo.APIVersion},
	}

	if method == "GET" {
		if reflect.TypeOf(r.Params) == reflect.TypeOf(&map[string]interface{}{}) {
			m := *(r.Params).(*map[string]interface{})
			for k, v := range m {
				body.Add(k, v.(string))
			}
		} else if err := queryutil.Parse(body, r.Params, false); err != nil {
			r.Error = volcstackerr.New("SerializationError", "failed encoding Query request", err)
			return
		}
	}
	r.HTTPRequest.URL.RawQuery = body.Encode()

	b, _ := json.Marshal(r.Params)
	str := string(b)
	r.SetStringBody(str)
	r.Params = nil
	r.HTTPRequest.Header.Set("Accept", "application/json")
}

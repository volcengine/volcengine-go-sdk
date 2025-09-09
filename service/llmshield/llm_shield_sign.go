/*
Copyright (year) Beijing Volcano Engine Technology Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package llmshield

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	// 请求接口信息
	Service = "llmshield"
	Version = "2025-08-31"
)

func hmacSHA256(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func getSignedKey(secretKey, date, region, service string) []byte {
	kDate := hmacSHA256([]byte(secretKey), date)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "request")

	return kSigning
}

func hashSHA256(data []byte) []byte {
	hash := sha256.New()
	if _, err := hash.Write(data); err != nil {
		log.Printf("input hash err:%s", err.Error())
	}

	return hash.Sum(nil)
}

func (c *Client) doRequestSign(request *http.Request, body []byte) error {
	queries := request.URL.Query()
	// 1. 构建请求
	_ = fmt.Sprintf("%s%s?%s", request.URL.Host, request.URL.Path, queries.Encode())
	//log.Printf("request addr: %s\n", requestAddr)

	//request, err := http.NewRequest(method, requestAddr, bytes.NewBuffer(body))
	//if err != nil {
	//	return nil, fmt.Errorf("bad request: %w", err)
	//}

	// 2. 构建签名材料
	now := time.Now()
	date := now.UTC().Format("20060102T150405Z")
	//date := "20250826T150405Z"
	authDate := date[:8]
	request.Header.Set("X-Date", date)
	request.Header.Set("X-Top-Service", Service)
	request.Header.Set("X-Top-Region", c.region)

	payload := hex.EncodeToString(hashSHA256(body))
	request.Header.Set("X-Content-Sha256", payload)
	//request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	queryString := strings.Replace(queries.Encode(), "+", "%20", -1)
	signedHeaders := []string{"host", "x-date", "x-content-sha256", "content-type"}
	var headerList []string
	for _, header := range signedHeaders {
		if header == "host" {
			headerList = append(headerList, header+":"+request.Host)
		} else {
			v := request.Header.Get(header)
			headerList = append(headerList, header+":"+strings.TrimSpace(v))
		}
	}
	headerString := strings.Join(headerList, "\n")

	canonicalString := strings.Join([]string{
		request.Method,
		request.URL.Path,
		queryString,
		headerString + "\n",
		strings.Join(signedHeaders, ";"),
		payload,
	}, "\n")
	//log.Printf("canonical string:\n%s\n", canonicalString)

	hashedCanonicalString := hex.EncodeToString(hashSHA256([]byte(canonicalString)))
	//log.Printf("hashed canonical string: %s\n", hashedCanonicalString)

	credentialScope := authDate + "/" + c.region + "/" + Service + "/request"
	signString := strings.Join([]string{
		"HMAC-SHA256",
		date,
		credentialScope,
		hashedCanonicalString,
	}, "\n")
	//log.Printf("sign string:\n%s\n", signString)

	// 3. 构建认证请求头
	signedKey := getSignedKey(c.sk, authDate, c.region, Service)
	signature := hex.EncodeToString(hmacSHA256(signedKey, signString))
	//log.Printf("signature: %s\n", signature)

	authorization := "HMAC-SHA256" +
		" Credential=" + c.ak + "/" + credentialScope +
		", SignedHeaders=" + strings.Join(signedHeaders, ";") +
		", Signature=" + signature
	request.Header.Set("Authorization", authorization)
	//log.Printf("authorization: %s\n", authorization)

	return nil
}

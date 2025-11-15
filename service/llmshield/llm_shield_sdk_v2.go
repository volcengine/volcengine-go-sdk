package llmshield

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	purl "net/url"
	"time"
)

// New 创建一个新的客户端实例
func New(url string, ak string, sk string, region string, timeout time.Duration) *Client {
	return &Client{
		url: url,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		ak:     ak,
		sk:     sk,
		region: region,
	}
}

func NewProxy(url string, ak string, sk string, region string, timeout time.Duration, proxyAddr string) *Client {
	proxyHandler, err := purl.Parse(proxyAddr)
	if err != nil {
		fmt.Printf("NewProxy Parse proxy addr error:%v\n", err)
		return nil
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyHandler), // 使用指定的代理
	}

	return &Client{
		url: url,
		httpClient: &http.Client{
			Transport: transport,
			Timeout:   timeout,
		},
		ak:     ak,
		sk:     sk,
		region: region,
	}
}

func SetServiceCode(IsDev bool) {
	if IsDev {
		Service = ServiceCodeDev
	} else {
		Service = ServiceCodeOnline
	}
}

func GetServiceCode() string {
	return Service
}

// Moderate 方法，根据传入的 Request 结构体反序列成 JSON 并发送请求，读取响应并转化为 Response 结构体
func (c *Client) Moderate(request *ModerateV2Request) (*ModerateV2Response, error) {
	if request == nil {
		request = &ModerateV2Request{
			Message: nil,
			Scene:   "",
		}
	}
	// 将请求结构体序列化为 JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", c.url+"/v2/moderate", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	queries := req.URL.Query()
	queries.Set("Action", "Moderate")
	queries.Set("Version", Version)
	// 将修改后的参数重新赋值给URL
	req.URL.RawQuery = queries.Encode()
	// 发送 HTTP 请求
	//req.Header.Set("x-api-key", c.api_key)
	c.doRequestSign(req, requestBody)
	// 4. 打印请求，发起请求
	//requestRaw, err := httputil.DumpRequest(req, true)
	//if err != nil {
	//	return nil, fmt.Errorf("dump request err: %w", err)
	//}
	//log.Printf("request:\n%s\n", string(requestRaw))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	response := &ModerateV2Response{}
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to Unmarshal response body: %w", err)
	}
	response.ResponseMetadata.HTTPCode = resp.StatusCode
	return response, nil
}

// ModerateStream 方法，根据传入的 Request 结构体反序列成 JSON 并发送请求，读取响应并转化为 Response 结构体
func (c *Client) ModerateStream(request *ModerateV2Request, session *ModerateV2StreamSession) (*ModerateV2Response, error) {
	if request == nil {
		request = &ModerateV2Request{}
	}
	if request.UseStream == 0 || session == nil {
		return nil, fmt.Errorf("useStream cannot 0 ,session cannot be null")
	}
	if request.Message == nil {
		return nil, fmt.Errorf("message cannot be null")
	}

	if session.Request == nil {
		session.Request = &ModerateV2Request{
			Message: &MessageV2{
				Role:        request.Message.Role,
				Content:     request.Message.Content,
				ContentType: request.Message.ContentType,
			},
			Scene:     request.Scene,
			History:   request.History,
			UseStream: request.UseStream,
		}
	} else {
		session.Request.Message.Content += request.Message.Content
		session.Request.UseStream = request.UseStream

	}

	session.StreamSendLen += int64(len(request.Message.Content))
	if session.currentSendWindow == 0 {
		session.currentSendWindow = LLM_STREAM_SEND_BASE_WINDOW_V2
	}

	if session.isNeedSend() != true {
		return session.DefaultBody, nil
	} else {
		session.StreamSendLen = 0
		session.currentSendWindow = session.currentSendWindow * LLM_STREAM_SEND_EXPONENT_V2
	}

	// 将请求结构体序列化为 JSON
	requestBody, err := json.Marshal(session.Request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", c.url+"/v2/moderate", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	queries := req.URL.Query()
	queries.Set("Action", "Moderate")
	queries.Set("Version", Version)
	// 发送 HTTP 请求
	//req.Header.Set("x-api-key", c.api_key)
	req.URL.RawQuery = queries.Encode()
	c.doRequestSign(req, requestBody)

	// 发送 HTTP 请求
	//req.Header.Set("x-api-key", c.api_key)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	response := &ModerateV2Response{}
	err = json.Unmarshal(responseBody, &response)
	response.ResponseMetadata.HTTPCode = resp.StatusCode
	if err != nil {
		return nil, fmt.Errorf("failed to Unmarshal response body: %w", err)
	}

	if session.Request.MsgID == "" {
		session.Request.MsgID = response.Result.MsgID
	}
	session.DefaultBody = response

	return response, nil
}

func (c *Client) GenerateV2Stream(request *GenerateStreamV2Request) (*GenerateStreamV2Response, error) {
	if request == nil {
		request = &GenerateStreamV2Request{}
	}

	// 将请求结构体序列化为 JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", c.url+"/v2/generate", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	queries := req.URL.Query()
	queries.Set("Action", "Generate")
	queries.Set("Version", Version)
	// 发送 HTTP 请求
	//req.Header.Set("x-api-key", c.api_key)
	req.URL.RawQuery = queries.Encode()
	c.doRequestSign(req, requestBody)

	// 发送 HTTP 请求
	//req.Header.Set("x-api-key", c.api_key)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("bad response code: %d", resp.StatusCode)
	} else {
		return &GenerateStreamV2Response{
			Reader: resp.Body,
		}, nil
	}

}

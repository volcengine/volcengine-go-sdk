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

func NewAdvanced(url string, ak string, sk string, region string, timeout time.Duration, proxyAddr string, connMax int) *Client {
	transport := &http.Transport{}

	if len(proxyAddr) > 0 {
		proxyHandler, err := purl.Parse(proxyAddr)
		if err != nil {
			fmt.Printf("NewProxy Parse proxy addr error:%v\n", err)
			return nil
		}

		transport.Proxy = http.ProxyURL(proxyHandler)
	}

	// 设置最大连接数
	if connMax > 0 {
		transport.MaxIdleConns = connMax
		transport.MaxConnsPerHost = connMax
		transport.MaxIdleConnsPerHost = connMax
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

func SetServiceDev(IsDev bool) {
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
	c.DoRequestSign(req, requestBody)
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

// ModerateStreamOld 方法，根据传入的 Request 结构体反序列成 JSON 并发送请求，读取响应并转化为 Response 结构体
func (c *Client) ModerateStreamOld(request *ModerateV2Request, session *ModerateV2StreamSession) (*ModerateV2Response, error) {
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
	c.DoRequestSign(req, requestBody)

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

func StreamSessionInit(streamId int64, chanSize int, isSync bool, role string) (stream *StreamSession) {
	if chanSize <= 0 {
		chanSize = RespChanSize
	}

	stream = &StreamSession{
		StreamId: streamId,
		ChanSize: chanSize,
		ReqDataChan: &streamReader{
			dataChan: make(chan []byte, chanSize),
			closed:   false,
		},
		Role:        role,
		RspDataChan: make(chan *ModerateV2Response, chanSize),
		IsSync:      isSync,
	}

	return
}

// ModerateStream 使用 HTTP POST 建立流式审核连接，并在同步模式下获取结果
func (c *Client) ModerateStream(request *ModerateV2Request, ssn interface{}) (*ModerateV2Response, error) {
	var session *StreamSession

	switch newssn := ssn.(type) {
	case *ModerateV2StreamSession:
		return c.ModerateStreamOld(request, newssn)
	case *StreamSession:
		session = newssn
	default:
		return nil, fmt.Errorf("ssn is error")
	}

	if session == nil {
		return nil, fmt.Errorf("session is error")
	}

	var startErr error
	isFirstCall := false
	session.once.Do(func() {
		isFirstCall = true
		session.Started = true

		if request == nil {
			startErr = fmt.Errorf("ModerateStream request cannot be nil")
			return
		}

		// 准备初始请求体并塞入通道
		requestBody, err := json.Marshal(request)
		if err != nil {
			startErr = fmt.Errorf("ModerateStream failed to marshal request: %v", err)
			return
		}
		session.ReqDataChan.Send(requestBody)

		// 创建 HTTP 请求，直接使用 ReqDataChan
		targetURL := c.url + "/v2/moderatestream"
		req, err := http.NewRequest("POST", targetURL, session.ReqDataChan)
		if err != nil {
			startErr = fmt.Errorf("ModerateStream failed to create request: %v", err)
			return
		}

		req.Header.Set("X-Scene", request.Scene)
		req.Header.Set("Content-Type", "application/json")

		queries := req.URL.Query()
		queries.Set("Action", "ModerateStream")
		queries.Set("Version", Version)
		queries.Set("X-NotSignBody", "stream")
		queries.Set("isSyncCheck", fmt.Sprintf("%v", session.IsSync))
		queries.Set("X-Role", session.Role)

		// 将修改后的参数重新赋值给URL
		req.URL.RawQuery = queries.Encode()

		// 签名 (流式请求不进行 body 鉴权，传 nil 以触发 UNSIGNED-PAYLOAD)
		if err := c.DoRequestSign(req, nil); err != nil {
			startErr = fmt.Errorf("ModerateStream failed to sign request: %v", err)
			return
		}

		// 将发送请求的操作放在独立协程中，确保不阻塞主流程
		go func() {
			if c.httpClient == nil {
				startErr = fmt.Errorf("ModerateStream httpClient is nil")
				return
			}
			resp, err := c.httpClient.Do(req)
			if err != nil {
				startErr = fmt.Errorf("ModerateStream httpClient.Do: %v\n", err)
				session.RspDataChan <- &ModerateV2Response{
					ResponseMetadata: ResponseMetadata{
						Error: &Error{
							Code:    "SDK_HTTP_ERROR",
							Message: err.Error(),
						},
					},
				}
				close(session.RspDataChan)
				return
			}
			defer resp.Body.Close()
			defer close(session.RspDataChan)

			if resp.StatusCode != http.StatusOK {
				body, _ := io.ReadAll(resp.Body)
				startErr = fmt.Errorf("ModerateStream Server RespCode:%d,error:%v\n", resp.StatusCode, string(body))
				return
			}

			// 解析 chunked JSON 响应
			decoder := json.NewDecoder(resp.Body)
			for {
				response := &ModerateV2Response{}
				if err := decoder.Decode(response); err != nil {
					if err == io.EOF {
						fmt.Printf("ModerateStream response stream parsing finished (EOF)\n")
						break
					}
					fmt.Printf("ModerateStream JSON parsing exception: %v\n", err)
					break
				}

				if response.ResponseMetadata.HTTPCode == 0 {
					response.ResponseMetadata.HTTPCode = resp.StatusCode
				}
				session.RspDataChan <- response
			}
		}()
	})

	if startErr != nil {
		return nil, startErr
	}

	// 2. 如果不是第一次进入此函数的调用，且有新数据，则写入数据通道发送给服务端
	if !isFirstCall && request != nil && request.Message != nil {
		if request.Message.Content != "" || len(request.Message.MultiPart) > 0 {
			requestBody, err := json.Marshal(request)
			if err == nil {
				session.ReqDataChan.Send(requestBody)
			}
		}
	}

	// 3. 处理流结束标识
	if request != nil && request.UseStream == 2 {
		session.ReqDataChan.Close()
	}

	// 4. 读取当前已有的响应并整合
	var mergedResponse *ModerateV2Response

	mergeFunc := func(base *ModerateV2Response, next *ModerateV2Response) *ModerateV2Response {
		if base == nil {
			return next
		}
		if next == nil {
			return base
		}

		// 整合 Result 中的字段
		// 字符串直接拼接
		base.Result.ContentInfo += next.Result.ContentInfo
		base.Result.DegradeReason += next.Result.DegradeReason

		// 列表 append
		if next.Result.RiskInfo != nil {
			if base.Result.RiskInfo == nil {
				base.Result.RiskInfo = next.Result.RiskInfo
			} else {
				base.Result.RiskInfo.Risks = append(base.Result.RiskInfo.Risks, next.Result.RiskInfo.Risks...)
			}
		}

		if next.Result.PermitInfo != nil {
			if base.Result.PermitInfo == nil {
				base.Result.PermitInfo = next.Result.PermitInfo
			} else {
				base.Result.PermitInfo.Permits = append(base.Result.PermitInfo.Permits, next.Result.PermitInfo.Permits...)
			}
		}

		if next.Result.Decision != nil {
			if base.Result.Decision == nil {
				base.Result.Decision = next.Result.Decision
			} else {
				// 列表 append
				base.Result.Decision.HitStrategyIDs = append(base.Result.Decision.HitStrategyIDs, next.Result.Decision.HitStrategyIDs...)

				// 决策类型以最新的为准
				if next.Result.Decision.DecisionType != 0 {
					base.Result.Decision.DecisionType = next.Result.Decision.DecisionType
				}

				// 整合替换详情
				if next.Result.Decision.Detail != nil && next.Result.Decision.Detail.ReplaceDetail != nil {
					if base.Result.Decision.Detail == nil {
						base.Result.Decision.Detail = next.Result.Decision.Detail
					} else if base.Result.Decision.Detail.ReplaceDetail == nil {
						base.Result.Decision.Detail.ReplaceDetail = next.Result.Decision.Detail.ReplaceDetail
					} else if next.Result.Decision.Detail.ReplaceDetail.Replacement != nil {
						if base.Result.Decision.Detail.ReplaceDetail.Replacement == nil {
							base.Result.Decision.Detail.ReplaceDetail.Replacement = next.Result.Decision.Detail.ReplaceDetail.Replacement
						} else {
							// 字符串拼接
							base.Result.Decision.Detail.ReplaceDetail.Replacement.Content += next.Result.Decision.Detail.ReplaceDetail.Replacement.Content
							// 列表 append
							base.Result.Decision.Detail.ReplaceDetail.Replacement.MultiPart = append(base.Result.Decision.Detail.ReplaceDetail.Replacement.MultiPart, next.Result.Decision.Detail.ReplaceDetail.Replacement.MultiPart...)
						}
					}
				}
			}
		}

		base.Result.Degraded = next.Result.Degraded
		// 元数据以最新的为准
		base.ResponseMetadata = next.ResponseMetadata

		return base
	}

	if request != nil && request.UseStream == 2 {
		// 只有最后一个请求需要一直持续等待（阻塞直到通道关闭）
		for resp := range session.RspDataChan {
			mergedResponse = mergeFunc(mergedResponse, resp)
		}
	} else {
		// 非最后一个请求，尝试非阻塞地读取当前所有可用响应
		for {
			select {
			case resp, ok := <-session.RspDataChan:
				if !ok {
					return mergedResponse, nil
				}
				mergedResponse = mergeFunc(mergedResponse, resp)
			default:
				// 无更多可用数据，立即返回
				return mergedResponse, nil
			}
		}
	}

	return mergedResponse, nil
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
	c.DoRequestSign(req, requestBody)

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

// Read 实现 io.Reader 接口（核心：每次读取通道中的数据，实现流式发送）
func (sr *streamReader) Read(p []byte) (n int, err error) {
	sr.mu.Lock()
	if len(sr.buffer) > 0 {
		n = copy(p, sr.buffer)
		sr.buffer = sr.buffer[n:]
		sr.mu.Unlock()
		return n, nil
	}

	if sr.closed {
		sr.mu.Unlock()
		return 0, io.EOF
	}
	sr.mu.Unlock()

	for {
		// 从通道中获取数据
		data, ok := <-sr.dataChan
		if !ok {
			sr.mu.Lock()
			sr.closed = true
			sr.mu.Unlock()
			return 0, io.EOF
		}

		// 只有拿到有效数据才返回，否则继续等待
		if len(data) == 0 {
			continue
		}

		sr.mu.Lock()
		n = copy(p, data)
		if n < len(data) {
			sr.buffer = data[n:]
		}
		sr.mu.Unlock()
		return n, nil
	}
}

// Send 向通道中写入数据（供外部调用，发送流式数据）
func (sr *streamReader) Send(data []byte) {
	sr.mu.Lock()
	if sr.closed {
		sr.mu.Unlock()
		return
	}
	ch := sr.dataChan
	sr.mu.Unlock()
	ch <- data
}

// Close 关闭通道（结束流式发送）
func (sr *streamReader) Close() {
	sr.mu.Lock()
	if sr.closed {
		sr.mu.Unlock()
		return
	}
	sr.closed = true
	ch := sr.dataChan
	sr.mu.Unlock()
	close(ch)
}

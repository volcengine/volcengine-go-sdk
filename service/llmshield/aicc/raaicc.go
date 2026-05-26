package aicc

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// RaPod 定义需要验证的Pod信息结构
type RaPod struct {
	ClusterID  string `json:"cluster_id"`    // 集群ID
	Namespace  string `json:"namespace"`     // 命名空间
	Deployment string `json:"deployment"`    // 部署名称
	Pod        string `json:"pod,omitempty"` // 可选，默认"random"，可选值"random"或"all"
}

// RaConfig 远程证明功能的配置结构
type RaConfig struct {
	RaURL            string   `json:"ra_url"`             // 远程证明端点的HTTP地址，默认空字符串
	RaType           string   `json:"ra_type"`            // RA类型，"local"表示本地RA，"tca"表示TCA的RAS_adapter接口（默认）
	RaServiceName    string   `json:"ra_service_name"`    // 数据接收端（服务端）的服务名，用于验证通过机密容器部署的服务
	RaPodsInfo       string   `json:"ra_pods_info"`       // 待验证节点的pod信息（非机密容器部署的服务），格式为JSON字符串：{"cluster_id":"***", "namespace":"***", "deployment":"***"}
	RaUID            string   `json:"ra_uid"`             // 用户的UID
	RaKeyNegotiation bool     `json:"ra_key_negotiation"` // 是否在验证时同步协商密钥，默认true
	RaNeedToken      bool     `json:"ra_need_token"`      // 是否返回验证token，默认true（RA功能稳定后需改为false）
	RaPolicyID       string   `json:"ra_policy_id"`       // 验证使用的策略ID，当RaNeedToken为true时必填
	BytedanceTopInfo string   `json:"bytedance_top_info"` // 访问字节跳动TOP网关的信息（JSON格式字符串）
	RaAttestedPods   []RaPod  `json:"ra_attested_pods"`   // 旧字段：需要验证的节点信息列表
	RaPolicyIDs      []string `json:"ra_policy_ids"`      // 旧字段：验证使用的策略ID列表，当RaNeedToken为true时必填
}

// NewRaConfig 创建默认的RaConfig实例
func NewRaConfig() *RaConfig {
	return &RaConfig{
		RaType:           "tca", // 默认TCA类型
		RaKeyNegotiation: true,  // 默认需要密钥协商
		RaNeedToken:      true,  // 默认需要返回token（后续稳定后需改为false）
	}
}

type ServiceError struct {
	Service string
	Url     string
	Code    interface{}
	Message string
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("ServiceError: service=%s, url=%s, code=%v, message=%s",
		e.Service, e.Url, e.Code, e.Message)
}

// 辅助方法：将RaPodsInfo字符串解析为RaPod结构（用于处理非机密容器服务的验证信息）
func (c *RaConfig) ParseRaPodsInfo() (*RaPod, error) {
	if c.RaPodsInfo == "" {
		return nil, nil
	}
	var pod RaPod
	if err := json.Unmarshal([]byte(c.RaPodsInfo), &pod); err != nil {
		return nil, err
	}
	return &pod, nil
}

// 辅助方法：将BytedanceTopInfo解析为map（用于访问TOP网关的配置）
func (c *RaConfig) ParseBytedanceTopInfo() (map[string]interface{}, error) {
	if c.BytedanceTopInfo == "" {
		return nil, nil
	}
	var info map[string]interface{}
	if err := json.Unmarshal([]byte(c.BytedanceTopInfo), &info); err != nil {
		return nil, err
	}
	return info, nil
}

// AttestServer 对数据接收方进行远程证明
// 返回 (ra_status, pub_key, error)
func AttestServer(config RaConfig, nonce string) (bool, string, error) {
	var response *http.Response // 假设已定义http.Response结构
	var err error

	// 根据配置决定调用远程还是本地证明
	if config.BytedanceTopInfo != "" {
		response, err = requestTCA(config, nonce)
		if err != nil {
			log.Printf("RA Service error: url=%s request failed: %v", config.RaURL, err)
			return false, "", &ServiceError{
				Service: "RA",
				Url:     config.RaURL,
				Message: "request failed",
			}
		}
	} else {
		// 调用本地证明函数，返回值需要适配
		status, pubKey := attestServerLocal(config)
		return status, pubKey, nil
	}

	// 检查响应是否为空
	if response == nil {
		log.Printf("RA Service error: url=%s response=NULL", config.RaURL)
		return false, "", &ServiceError{
			Service: "RA",
			Url:     config.RaURL,
			Message: "NULL response",
		}
	}
	defer response.Body.Close()

	// 检查响应状态码
	if response.StatusCode != 200 {
		// 读取响应内容
		body, _ := io.ReadAll(response.Body)
		log.Printf("RA Service error: return code=%d, return content=%s",
			response.StatusCode, string(body))
		return false, "", &ServiceError{
			Service: "RA",
			Url:     config.RaURL,
			Code:    response.StatusCode,
			Message: fmt.Sprintf("code=%d", response.StatusCode),
		}
	}

	// 解析JSON响应
	var responseJSON map[string]interface{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&responseJSON); err != nil {
		log.Printf("Response is not JSON: service=%s, error=%v", config.RaURL, err)
		return false, "", &ServiceError{
			Service: "RA",
			Url:     config.RaURL,
			Message: "response is not JSON",
		}
	}

	raRes, ok := responseJSON["Result"].(map[string]interface{})
	if !ok || raRes == nil {
		log.Printf("RA response has no valid Result: %v", responseJSON)
		return false, "", nil
	}

	attestationStatus := false
	publicKeyInfo := ""

	// 取第一个节点验证
	for clusterInfo, resInfo := range raRes {
		resInfoMap, ok := resInfo.(map[string]interface{})
		if !ok {
			log.Printf("Invalid resInfo type for cluster %s", clusterInfo)
			continue
		}

		// 获取token
		rawToken, ok := resInfoMap["token"].(string)
		if !ok || rawToken == "" {
			log.Printf("token is None for cluster %s", clusterInfo)
			return false, "", nil
		}

		// 解析证据信息
		evidenceInfo, ok := resInfoMap["evidence"].(string)
		if !ok {
			log.Printf("evidence info is invalid for cluster %s", clusterInfo)
			return false, "", nil
		}

		var evidenceJSON map[string]interface{}
		if err := json.Unmarshal([]byte(evidenceInfo), &evidenceJSON); err != nil {
			log.Printf("Failed to parse evidence info: %v", err)
			return false, "", nil
		}

		proof, _ := evidenceJSON["proof"].(string)
		log.Printf("proof=%s", proof)

		// 解析token获取report_data
		tokenParts := strings.Split(rawToken, ".")
		if len(tokenParts) < 2 {
			log.Printf("Invalid token format")
			return false, "", nil
		}

		tokenData := tokenParts[1]
		// 处理Base64填充
		padding := 4 - (len(tokenData) % 4)
		if padding < 4 {
			tokenData += strings.Repeat("=", padding)
		}

		decodedData, err := base64.URLEncoding.DecodeString(tokenData)
		if err != nil {
			log.Printf("Failed to decode token data: %v", err)
			return false, "", errors.New("decode token error")
		}

		var tokenDict map[string]interface{}
		if err := json.Unmarshal(decodedData, &tokenDict); err != nil {
			log.Printf("Failed to parse token data: %v", err)
			return false, "", errors.New("decode token error")
		}

		// 获取report_data
		tdx, ok := tokenDict["tdx"].(map[string]interface{})
		if !ok {
			log.Printf("token has no tdx field")
			return false, "", nil
		}
		reportData, _ := tdx["report_data"].(string)
		log.Printf("report_data=%s", reportData)

		// 处理公钥信息
		if config.RaKeyNegotiation {
			keyInfo, ok := resInfoMap["key_info"].(map[string]interface{})
			if ok {
				pubKeyInfo, _ := keyInfo["pub_key_info"].(string)
				publicKeyInfo = pubKeyInfo
			}
		} else {
			publicKeyInfo = ""
		}

		// 验证nonce
		if !verifyNonce(proof, reportData, nonce, publicKeyInfo) {
			log.Printf("verify nonce failed")
			return false, "", nil
		}

		// 验证JWT token
		if !verifyJwtToken(rawToken) {
			log.Printf("token verify failed")
			return false, "", nil
		}

		// 再次解析token检查策略匹配
		var tokenDataDict map[string]interface{}
		if err := json.Unmarshal(decodedData, &tokenDataDict); err != nil {
			log.Printf("token decode error: %v", err)
			return false, "", nil
		}

		// 检查策略匹配
		policiesMatched, ok := tokenDataDict["policies_matched"].([]interface{})
		if ok {
			attestationStatus = len(policiesMatched) > 0
		}

		// 只处理第一个节点
		break
	}

	return attestationStatus, publicKeyInfo, nil
}

// requestTCA 向TCA服务发送请求
func requestTCA(config RaConfig, nonce string) (*http.Response, error) {
	// 从配置中获取参数
	raURL := config.RaURL
	raServiceName := config.RaServiceName
	raPodsInfo := config.RaPodsInfo
	raKeyNegotiation := config.RaKeyNegotiation
	raPolicyID := config.RaPolicyID
	raUID := config.RaUID

	// 处理top_info
	topInfo := config.BytedanceTopInfo
	var topInfoMap map[string]interface{}
	err := json.Unmarshal([]byte(topInfo), &topInfoMap)
	if err != nil {
		fmt.Printf("Failed to parse top_info: %v", err)
		return nil, &ServiceError{
			Service: "RA",
			Url:     raURL,
			Message: "invalid top_info format",
		}
	}

	// 设置请求头
	headers := map[string]string{
		"UID": raUID,
	}

	// 构建请求体
	var body map[string]interface{}
	switch {
	case raServiceName != "":
		body = map[string]interface{}{
			"Nonce":          nonce,
			"PolicyID":       raPolicyID,
			"ServiceName":    raServiceName,
			"KeyNegotiation": raKeyNegotiation,
			"Token":          true,
		}
	case raPodsInfo != "":
		// 解析raPodsInfo
		var podsInfo interface{}
		if err := json.Unmarshal([]byte(raPodsInfo), &podsInfo); err != nil {
			log.Printf("Failed to parse ra_pods_info: %v", err)
			return nil, &ServiceError{
				Service: "RA",
				Url:     raURL,
				Message: "invalid ra_pods_info format",
			}
		}
		body = map[string]interface{}{
			"Nonce":          nonce,
			"PolicyID":       raPolicyID,
			"AttestedPods":   []interface{}{podsInfo},
			"KeyNegotiation": raKeyNegotiation,
			"Token":          true,
		}
	default:
		return nil, &ServiceError{
			Service: "RA",
			Url:     raURL,
			Message: "call TCA error: missing service name or pods info",
		}
	}

	// 调用字节跳动网关发送请求
	return RequestBytedanceGateway(raURL, body, headers, topInfoMap)
}

// RequestBytedanceGateway 调用字节跳动网关
func RequestBytedanceGateway(url string, body map[string]interface{},
	headers map[string]string, topInfo map[string]interface{}) (*http.Response, error) {
	// 实现网关请求逻辑
	// 1. 序列化body为JSON
	// 2. 创建HTTP请求
	// 3. 设置请求头
	// 4. 处理topInfo（如果需要）
	// 5. 发送请求并返回响应
	return nil, nil
}

func attestServerLocal(config RaConfig) (bool, string) {
	// 实现本地证明逻辑
	return false, ""
}

func verifyNonce(proof, reportData, nonce, pubKeyInfo string) bool {
	// 实现nonce验证逻辑
	return true
}

func verifyJwtToken(token string) bool {
	// 实现JWT验证逻辑
	return true
}

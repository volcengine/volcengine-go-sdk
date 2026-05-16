package aicc

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// ClientConfig 安全通信数据发送方配置（继承自RaConfig）
type ClientConfig struct {
	RaConfig
	AttestInterval *float64 `json:"attest_interval"` // 自动远程证明的时间间隔（秒），nil表示不自动
	RootKeyConfig  *string  `json:"root_key_config"` // root_key配置（JSON格式字符串）
	LogConfig      *string  `json:"log_config"`      // 日志配置
	PubKeyPath     *string  `json:"pub_key_path"`    // 公钥文件路径
}

// NewClientConfig 创建默认的ClientConfig实例
func NewClientConfig() *ClientConfig {
	return &ClientConfig{
		RaConfig:       *NewRaConfig(), // 继承RaConfig的默认配置
		AttestInterval: nil,            // 默认不自动进行远程证明
		// 其他字段默认值为nil（可选字段）
	}
}

// FromFile 从文件读取配置并创建ClientConfig实例
func (c *ClientConfig) FromFile(path string) (*ClientConfig, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	var configMap map[string]interface{}
	if err := json.Unmarshal(data, &configMap); err != nil {
		return nil, err
	}

	// 移除$schema字段（如果存在）
	delete(configMap, "$schema")

	// 转换为JSON字节流，再反序列化为ClientConfig
	configJSON, err := json.Marshal(configMap)
	if err != nil {
		return nil, err
	}

	var clientConfig ClientConfig
	if err := json.Unmarshal(configJSON, &clientConfig); err != nil {
		return nil, err
	}

	// 确保继承的RaConfig字段有默认值
	if clientConfig.RaType == "" {
		clientConfig.RaType = "tca"
	}
	if clientConfig.RaKeyNegotiation == false {
		clientConfig.RaKeyNegotiation = true // 保持默认值为true
	}
	if clientConfig.RaNeedToken == false {
		clientConfig.RaNeedToken = true // 保持默认值为true
	}

	return &clientConfig, nil
}

// FromDict 从map创建ClientConfig实例
func (c *ClientConfig) FromDict(obj map[string]interface{}) (*ClientConfig, error) {
	// 移除$schema字段
	delete(obj, "$schema")

	// 转换为JSON字节流，再反序列化为ClientConfig
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var config ClientConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	// 补全默认值
	if config.RaType == "" {
		config.RaType = "tca"
	}
	if config.RaKeyNegotiation == false {
		config.RaKeyNegotiation = true
	}
	if config.RaNeedToken == false {
		config.RaNeedToken = true
	}

	return &config, nil
}

// FromEnv 从环境变量读取配置并创建ClientConfig实例
func (c *ClientConfig) FromEnv(envKey string) (*ClientConfig, error) {
	configStr := os.Getenv(envKey)
	if configStr == "" {
		configStr = "{}" // 默认空配置
	}

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(configStr), &obj); err != nil {
		return nil, err
	}

	return c.FromDict(obj)
}

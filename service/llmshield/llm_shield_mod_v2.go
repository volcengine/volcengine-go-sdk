package llmshield

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ContentTypeV2 int64

const (
	// 文本
	ContentTypeV2_TEXT ContentTypeV2 = 1
	// 音频
	ContentTypeV2_AUDIO ContentTypeV2 = 2
	// 图片
	ContentTypeV2_IMAGE ContentTypeV2 = 3
	// 视频
	ContentTypeV2_VIDEO ContentTypeV2 = 4
)

type DecisionTypeV2 int64

const (
	// 通过
	DecisionTypeV2_PASS DecisionTypeV2 = 1
	// 拦截
	DecisionTypeV2_BLOCK DecisionTypeV2 = 2
	// 标记
	DecisionTypeV2_MARK DecisionTypeV2 = 3
	// 替换
	DecisionTypeV2_REPLACE DecisionTypeV2 = 4
	// 优化回答
	DecisionTypeV2_OPTIMIZE DecisionTypeV2 = 5
)

const (
	LLM_STREAM_SEND_BASE_WINDOW_V2 int64 = 10
	LLM_STREAM_SEND_EXPONENT_V2    int64 = 2
)

type MessageV2 struct {
	// 消息内容ID @tpl=select @jsonCEnums=["user","assistant","system","rag"]
	Role string `thrift:"role,1" form:"Role" json:"Role"`
	// 内容文本或链接
	Content string `thrift:"content,2" form:"Content" json:"Content"`
	// 内容类型
	ContentType ContentTypeV2 `thrift:"contentType,3" form:"ContentType" json:"ContentType"`
}

type ModerateV2Request struct {
	// 审核内容
	Message *MessageV2 `thrift:"message,1,required" form:"Message,required" json:"Message,required"`
	// 绑定流式送检会话的id
	MsgID string `thrift:"MsgID,3" form:"MsgID" json:"MsgID"`
	// 0:不使用流式一次性审核，1:使用流式，
	UseStream int64 `thrift:"UseStream,4" form:"UseStream" json:"UseStream"`
	// 场景 @tpl=select @enums=SCENE
	Scene string `thrift:"scene,5,optional" form:"Scene" json:"Scene,omitempty"`
	// 历史消息
	History []*MessageV2 `thrift:"history,6,optional" form:"History" json:"History,omitempty"`
}
type Error struct {
	CodeN   int
	Code    string
	Message string
}

type ResponseMetadata struct {
	RequestId string
	Action    string
	Version   string
	Service   string
	Region    string
	HTTPCode  int
	Error     *Error
}

type ModerateV2Response struct {
	ResponseMetadata ResponseMetadata `json:"ResponseMetadata"`
	Result           ModerateV2Result `json:"Result"`
}

type ModerateV2Result struct {
	// 消息ID，表示请求的唯一标识
	MsgID string `thrift:"msgID,1" form:"MsgID" json:"MsgID"`
	// 风险信息
	RiskInfo *RiskInfoV2 `thrift:"riskInfo,2" form:"RiskInfo" json:"RiskInfo"`
	// 决策
	Decision *DecisionV2 `thrift:"decision,3" form:"Decision" json:"Decision"`
	// 放行信息
	PermitInfo *PermitInfoV2 `thrift:"permitInfo,4" form:"PermitInfo" json:"PermitInfo"`
	// 是否降级
	Degraded bool `thrift:"degraded,101" form:"Degraded" json:"Degraded"`
	// 降级原因
	DegradeReason string `thrift:"degradeReason,102" form:"DegradeReason" json:"DegradeReason"`
	// 内部调试信息，仅在debug模式下返回
	//DebugInfo map[string]string `thrift:"debugInfo,201,optional" form:"DebugInfo" json:"DebugInfo,omitempty"`
}

func (p *ModerateV2Response) String() string {
	if p == nil {
		return "<nil>"
	}
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(b)

}

type UserAction int64

const (
	// 放行
	UserAction_PASS UserAction = 1
	// 拦截
	UserAction_BLOCK UserAction = 2
	// 标记
	UserAction_MARK UserAction = 3
	// 替换
	UserAction_REPLACE UserAction = 4
)

type MatchSource int64

const (
	// 未知，或检出模块未显式指定来源
	MatchSource_UNKNOWN MatchSource = 0
	// 全局内容库
	MatchSource_GLOBAL_CONTENTLIB MatchSource = 1
	// 租户黑白名单库
	MatchSource_ADMIN_CONTENTLIB MatchSource = 2
	// 用户自定义词库
	MatchSource_USER_CONTENTLIB MatchSource = 3
)

type RiskMatchV2 struct {
	// 命中词
	Word string `thrift:"word,1" form:"Word" json:"Word"`
	// 命中动作
	Action *UserAction `thrift:"action,2,optional" form:"Action" json:"Action,omitempty"`
	// 命中词来源
	Source MatchSource `thrift:"source,3" form:"Source" json:"Source"`
	// 规则ID
	RuleID string `thrift:"ruleID,4" form:"RuleID" json:"RuleID,omitempty"`
}
type PermitMatchV2 struct {
	// 命中词
	Word string `thrift:"word,1" form:"Word" json:"Word"`
	// 命中动作
	Action *UserAction `thrift:"action,2,optional" form:"Action" json:"Action,omitempty"`
	// 命中词来源
	Source MatchSource `thrift:"source,3" form:"Source" json:"Source"`
	// 规则ID
	RuleID string `thrift:"ruleID,2" form:"RuleID" json:"RuleID,omitempty"`
}

type RiskV2 struct {
	// 风险分类
	Category string `thrift:"category,1" form:"Category" json:"Category"`
	// 风险标签
	Label string `thrift:"label,2" form:"Label" json:"Label"`
	// 风险概率
	Prob *float64 `thrift:"prob,3,optional" form:"Prob" json:"Prob,omitempty"`
	// 命中词表信息
	Matches []*RiskMatchV2 `thrift:"matches,4,optional" form:"Matches" json:"Matches,omitempty"`
}
type RiskInfoV2 struct {
	// 风险信息
	Risks []*RiskV2 `thrift:"risks,1" form:"Risks" json:"Risks"`
}

type PermitV2 struct {
	// 放行分类，保留字段
	Category string `thrift:"category,1" form:"Category" json:"Category"`
	// 放行标签
	Label string `thrift:"label,2" form:"Label" json:"Label"`
	// 命中放行概率
	Prob *float64 `thrift:"prob,3,optional" form:"Prob" json:"Prob,omitempty"`
	// 命中白词表信息
	Matches []*PermitMatchV2 `thrift:"matches,4,optional" form:"Matches" json:"Matches,omitempty"`
}

type PermitInfoV2 struct {
	// 放行命中信息
	Permits []*PermitV2 `thrift:"permits,1" form:"Permits" json:"Permits"`
}

type BlockDetailV2 struct {
}

type ReplaceDetailV2 struct {
	// 替换内容
	Replacement *MessageV2 `thrift:"replacement,1" form:"Replacement" json:"Replacement"`
}

type DecisionDetailV2 struct {
	// 拦截详情
	BlockDetail *BlockDetailV2 `thrift:"blockDetail,1,optional" form:"BlockDetail" json:"BlockDetail,omitempty"`
	// 替换详情
	ReplaceDetail *ReplaceDetailV2 `thrift:"replaceDetail,2,optional" form:"ReplaceDetail" json:"ReplaceDetail,omitempty"`
}

type DecisionV2 struct {
	// 决策类型
	DecisionType DecisionTypeV2 `thrift:"decisionType,1" form:"DecisionType" json:"DecisionType"`
	// 决策详情
	Detail *DecisionDetailV2 `thrift:"detail,2" form:"DecisionDetail" json:"DecisionDetail"`
	// 决策策略ID
	DecisionStrategyID *string `thrift:"decisionStrategyID,3,optional" form:"DecisionStrategyID" json:"DecisionStrategyID,omitempty"`
	// 命中策略ID列表
	HitStrategyIDs []string `thrift:"hitStrategyIDs,4,optional" form:"HitStrategyIDs" json:"HitStrategyIDs,omitempty"`
}

type ModerateV2StreamSession struct {
	// 用于积累流式的请求
	Request *ModerateV2Request
	// 未发送的长度
	StreamSendLen int64
	// 存储默认响应体
	DefaultBody       *ModerateV2Response
	currentSendWindow int64
}

// ModerateV2StreamSession
func (s *ModerateV2StreamSession) isNeedSend() bool {
	if s.StreamSendLen >= s.currentSendWindow || s.Request.MsgID == "" || s.Request.UseStream == 2 {
		return true
	}
	return false
}

// GenerateStreamV2Request 表示生成流V2版本的请求结构体
type GenerateStreamV2Request struct {
	// MsgID 消息ID，表示请求的唯一标识
	MsgID string `json:"MsgID" validate:"required"`
}

type GenerateStreamV2Response struct {
	Reader io.ReadCloser
}

type GenerateStreamV2ResponseData struct {
	ResponseMetadata ResponseMetadata     `json:"ResponseMetadata"`
	Result           GenerateStreamResult `json:"Result"`
}

// GenerateStreamV2Response 表示生成流V2版本的响应结构体
type GenerateStreamResult struct {
	// Message 优化内容，isFinished为true时为空
	Message MessageV2 `json:"Message"`
	// IsFinished 标识是否结束
	IsFinished bool `json:"IsFinished"`
	// Summarize 总结信息，isFinished为true时有值
	//Summarize *GenerateSummarizeV2 `json:"Summarize,omitempty"`
}

// GenerateSummarizeV2 表示生成过程的总结信息结构体
type GenerateSummarizeV2 struct {
	// TokenCost 消耗的token数量
	TokenCost int64 `json:"TokenCost"`
	// TimeCostMS 消耗的时长（毫秒）
	TimeCostMS int64 `json:"TimeCostMS"`
}

type Client struct {
	url        string
	region     string
	ak         string
	sk         string
	httpClient *http.Client
}

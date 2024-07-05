package model

type TokenizationRequestConverter interface {
	Convert() TokenizationRequest
}

type TokenizationRequest struct {
	Text  interface{} `json:"text"`
	Model string      `json:"model"`
	User  string      `json:"user"`
}

func (r TokenizationRequest) Convert() TokenizationRequest {
	return r
}

// TokenizationRequestStrings is the input to a create tokenization request with a slice of strings.
type TokenizationRequestStrings struct {
	Text  []string `json:"text"`
	Model string   `json:"model"`
	User  string   `json:"user"`
}

func (r TokenizationRequestStrings) Convert() TokenizationRequest {
	return TokenizationRequest{
		Text:  r.Text,
		Model: r.Model,
		User:  r.User,
	}
}

// TokenizationRequestString is the input to a create tokenization request with a slice of strings.
type TokenizationRequestString struct {
	Text  string `json:"text"`
	Model string `json:"model"`
	User  string `json:"user"`
}

func (r TokenizationRequestString) Convert() TokenizationRequest {
	return TokenizationRequest{
		Text:  r.Text,
		Model: r.Model,
		User:  r.User,
	}
}

// TokenizationResponse is the response from a Create tokenization request.
type TokenizationResponse struct {
	ID      string          `json:"id"`
	Created int             `json:"created"`
	Model   string          `json:"model"`
	Object  string          `json:"object"`
	Data    []*Tokenization `json:"data"`

	HttpHeader
}

type Tokenization struct {
	Index         int     `json:"index"`
	Object        string  `json:"object"`
	TotalTokens   int     `json:"total_tokens"`
	TokenIDs      []int   `json:"token_ids"`
	OffsetMapping [][]int `json:"offset_mapping"`
}

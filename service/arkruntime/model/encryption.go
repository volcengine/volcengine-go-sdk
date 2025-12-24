package model

type EncryptInfo struct {
	Version    string `thrift:"Version,1,optional" header:"Version" json:"Version,omitempty"`
	RingID     string `thrift:"RingID,2,optional" header:"RingID" json:"RingID,omitempty"`
	KeyID      string `thrift:"KeyID,3,optional" header:"KeyID" json:"KeyID,omitempty"`
	ExpireTime int64  `thrift:"ExpireTime,4,optional" header:"ExpireTime" json:"ExpireTime,omitempty"`
}

type CertificateResponse struct {
	Error       map[string]string `json:"error,omitempty"`
	Certificate string            `json:"Certificate"`
}

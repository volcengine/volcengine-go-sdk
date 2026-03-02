package base

const (
	timeFormatV4 = "20060102T150405Z"
)

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
	Service         string
	Region          string
	SessionToken    string
}

type metadata struct {
	algorithm       string
	credentialScope string
	signedHeaders   string
	date            string
	region          string
	service         string
}

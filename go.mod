module github.com/volcengine/volcengine-go-sdk

go 1.14

require (
	github.com/google/uuid v1.3.0
	github.com/jmespath/go-jmespath v0.4.0
	github.com/klauspost/compress v1.13.5
	github.com/volcengine/volc-sdk-golang v1.0.23
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v2 v2.2.8
)

// Responses API decoding rejects unknown response fields and can fail on
// otherwise valid responses.
retract [v1.1.51, v1.1.55]

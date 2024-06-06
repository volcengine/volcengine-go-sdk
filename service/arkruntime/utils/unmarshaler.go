package utils

import (
	"encoding/json"
)

type Unmarshaler interface {
	Unmarshal(data []byte, v interface{}) error
}

type JSONUnmarshaler struct{}

func (jm *JSONUnmarshaler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

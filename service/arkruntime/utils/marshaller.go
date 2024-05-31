package utils

import (
	"encoding/json"
)

type Marshaller interface {
	Marshal(value interface{}) ([]byte, error)
}

type JSONMarshaller struct{}

func (jm *JSONMarshaller) Marshal(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

// Package contextmanagement is the contextmanagement service.
package contextmanagement

import (
	"encoding/json"
	reflect "reflect"
)

// MarshalJSON marshals the ToolUsesKeepType enum to JSON.
func (r ToolUsesKeepTypeEnum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the field pointed to by r.
func (r *ToolUsesKeepTypeEnum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ToolUsesKeepTypeEnum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ToolUsesKeepTypeEnum(enumValue)
	return nil
}

// MarshalJSON marshals the StringAllEnum enum to JSON.
func (r StringAllEnum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the field pointed to by r.
func (r *StringAllEnum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := StringAllEnum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = StringAllEnum(enumValue)
	return nil
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the field pointed to by r.
func (r *ThinkingTurnsParamTypeEnum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ThinkingTurnsParamTypeEnum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ThinkingTurnsParamTypeEnum(enumValue)
	return nil
}

// MarshalJSON marshals the ThinkingTurnsParamType enum to JSON.
func (r ThinkingTurnsParamTypeEnum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the field pointed to by r.
func (r *ToolUsesTriggerTypeEnum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ToolUsesTriggerTypeEnum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ToolUsesTriggerTypeEnum(enumValue)
	return nil
}

// MarshalJSON marshals the ToolUsesTriggerType enum to JSON.
func (r ToolUsesTriggerTypeEnum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// MarshalJSON marshals the EditTypeEnum enum to JSON.
func (r EditTypeEnum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the field pointed to by r.
func (r *EditTypeEnum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := EditTypeEnum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = EditTypeEnum(enumValue)
	return nil
}

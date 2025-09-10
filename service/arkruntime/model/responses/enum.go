// Package responses provides the responses model of the ark runtime service.
package responses

import (
	"encoding/json"
	"reflect"
)

// UnmarshalJSON ...
func (r *ResponsesServiceTier_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ResponsesServiceTier_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ResponsesServiceTier_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ResponsesServiceTier_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *ResponsesTruncation_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ResponsesTruncation_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ResponsesTruncation_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ResponsesTruncation_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *ItemType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ItemType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ItemType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ItemType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *MessageRole_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := MessageRole_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = MessageRole_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r MessageRole_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *ContentItemImageDetail_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ContentItemImageDetail_Enum_value[value]
	if !ok {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ContentItemImageDetail_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ContentItemImageDetail_Enum) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *ContentItemType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ContentItemType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ContentItemType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ContentItemType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *ItemStatus_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ItemStatus_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ItemStatus_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ItemStatus_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *ThinkingType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ThinkingType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ThinkingType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ThinkingType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *TextType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := TextType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = TextType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r IncludeType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return []byte(IncludeType_Enum_name[int32(r)]), nil
}

// UnmarshalJSON ...
func (r *IncludeType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := IncludeType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = IncludeType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r CacheType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *CacheType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := CacheType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = CacheType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r TextType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *ToolChoiceMode_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ToolChoiceMode_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ToolChoiceMode_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ToolChoiceMode_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *ToolType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ToolType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ToolType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ToolType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *UserLocationType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := UserLocationType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = UserLocationType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r UserLocationType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// MarshalJSON ...
func (r SourceType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *SourceType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := SourceType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = SourceType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ObjectType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

func (r *ObjectType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ObjectType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ObjectType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ResponseStatus_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

func (r *ResponseStatus_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ResponseStatus_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ResponseStatus_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r EventType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

func (r *EventType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := EventType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = EventType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ActionType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *ActionType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ActionType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ActionType_Enum(enumValue)
	return nil
}

// UnmarshalJSON ...
func (r *ChunkingStrategyType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ChunkingStrategyType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ChunkingStrategyType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ChunkingStrategyType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *AnnotationType_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := AnnotationType_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = AnnotationType_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r AnnotationType_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *ReasoningEffort_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ReasoningEffort_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ReasoningEffort_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ReasoningEffort_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

// UnmarshalJSON ...
func (r *ApprovalMode_Enum) UnmarshalJSON(bytes []byte) error {
	var value string
	if err := json.Unmarshal(bytes, &value); err != nil {
		return err
	}
	enumValue, ok := ApprovalMode_Enum_value[value]
	if !ok || enumValue == 0 {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(r)}
	}
	*r = ApprovalMode_Enum(enumValue)
	return nil
}

// MarshalJSON ...
func (r ApprovalMode_Enum) MarshalJSON() ([]byte, error) {
	if r == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(r.String())
}

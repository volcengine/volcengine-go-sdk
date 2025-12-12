package contextmanagement

import "encoding/json"

func unmarshal(message []byte, val interface{}) error {
	return json.Unmarshal(message, val)
}

// MarshalJSON marshals the Edit union to JSON.
func (r *Edit) MarshalJSON() ([]byte, error) {
	if v := r.GetClearToolUses(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetClearThinking(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the field pointed to by r.
func (r *Edit) UnmarshalJSON(bytes []byte) error {
	var err error
	var typeOnly struct {
		Type *EditTypeEnum `json:"type,omitempty"`
	}
	if err = json.Unmarshal(bytes, &typeOnly); err != nil {
		return err
	}
	if typeOnly.Type == nil {
		return err
	}
	switch *typeOnly.Type {
	case EditType_clear_tool_uses:
		clearToolUses := Edit_ClearToolUses{}
		if err = unmarshal(bytes, &clearToolUses.ClearToolUses); err != nil {
			return err
		}
		r.Union = &clearToolUses
	case EditType_clear_thinking:
		clearThinking := Edit_ClearThinking{}
		if err = unmarshal(bytes, &clearThinking.ClearThinking); err != nil {
			return err
		}
		r.Union = &clearThinking
	default:
		return err
	}
	return nil
}

// MarshalJSON marshals the ClearThinkingKeep union to JSON.
func (r *ClearThinkingKeep) MarshalJSON() ([]byte, error) {
	if v := r.GetAll(); v != StringAll_unspecified {
		return json.Marshal(v)
	}
	if v := r.GetThinkingTurnParams(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the field pointed to by r.
func (r *ClearThinkingKeep) UnmarshalJSON(bytes []byte) error {
	var err error
	oneof1 := ClearThinkingKeep_All{}
	if err = unmarshal(bytes, &oneof1.All); err == nil {
		r.Union = &oneof1
		return nil
	}

	oneof2 := ClearThinkingKeep_ThinkingTurnParams{}
	if err = unmarshal(bytes, &oneof2.ThinkingTurnParams); err == nil {
		r.Union = &oneof2
		return nil
	}
	return err
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the field pointed to by r.
func (r *ClearToolUsesInputs) UnmarshalJSON(bytes []byte) error {
	var err error
	oneof1 := ClearToolUsesInputs_ClearAll{}
	if err = unmarshal(bytes, &oneof1.ClearAll); err == nil {
		r.Union = &oneof1
		return nil
	}
	oneof2 := ClearToolUsesInputs_ToolNameList{}
	if err = unmarshal(bytes, &oneof2.ToolNameList); err == nil {
		r.Union = &oneof2
		return nil
	}
	return err
}

// MarshalJSON marshals the ClearToolUsesInputs union to JSON.
func (r *ClearToolUsesInputs) MarshalJSON() ([]byte, error) {
	if v := r.GetToolNameList(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(r.GetClearAll())
}

// MarshalJSON marshals the AppliedEdit union to JSON.
func (r *AppliedEdit) MarshalJSON() ([]byte, error) {
	if v := r.GetClearToolUsesResponse(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetClearThinkingResponse(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the field pointed to by r.
func (r *AppliedEdit) UnmarshalJSON(bytes []byte) error {
	var err error
	oneof1 := AppliedEdit_ClearToolUsesResponse{}
	if err = unmarshal(bytes, &oneof1.ClearToolUsesResponse); err == nil {
		r.Union = &oneof1
		return nil
	}
	oneof2 := AppliedEdit_ClearThinkingResponse{}
	if err = unmarshal(bytes, &oneof2.ClearThinkingResponse); err == nil {
		r.Union = &oneof2
		return nil
	}
	return err
}

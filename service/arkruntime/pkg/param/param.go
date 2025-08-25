package param

import "encoding/json"

// ParamNullable encapsulates all structs in parameters,
// and all [Opt] types in parameters.
type ParamNullable interface {
	null() bool
}

// ParamStruct represents the set of all structs that are
// used in API parameters, by convention these usually end in
// "Params" or "Param".
type ParamStruct interface {
	Overrides() (interface{}, bool)
	null() bool
	extraFields() map[string]interface{}
}

// APIObject should be embedded in api object fields, preferably using an alias to make private
type APIObject struct{ metadata }

// APIUnion should be embedded in all api unions fields, preferably using an alias to make private
type APIUnion struct{ metadata }

// Overrides returns the value of the struct when it is created with
// [Override], the second argument helps differentiate an explicit null.
func (m metadata) Overrides() (interface{}, bool) {
	if _, ok := m.any.(metadataExtraFields); ok {
		return nil, false
	}
	return m.any, m.any != nil
}

// ExtraFields returns the extra fields added to the JSON object.
func (m metadata) ExtraFields() map[string]interface{} {
	if extras, ok := m.any.(metadataExtraFields); ok {
		return extras
	}
	return nil
}

type metadata struct{ any interface{} }
type metadataNull struct{}
type metadataExtraFields map[string]interface{}

func (m *metadata) setMetadata(override interface{}) {
	if override == nil {
		m.any = metadataNull{}
		return
	}
	m.any = override
}

// SetExtraFields adds extra fields to the JSON object.
//
// SetExtraFields will override any existing fields with the same key.
// For security reasons, ensure this is only used with trusted input data.
//
// To intentionally omit a required field, use [Omit].
//
//	foo.SetExtraFields(map[string]any{"bar": Omit})
//
// If the struct already contains the field ExtraFields, then this
// method will have no effect.
func (m *metadata) SetExtraFields(extraFields map[string]interface{}) {
	m.any = metadataExtraFields(extraFields)
}

// extraFields aliases [metadata.ExtraFields] to avoid name collisions.
func (m metadata) extraFields() map[string]interface{} { return m.ExtraFields() }

func (m metadata) null() bool {
	if _, ok := m.any.(metadataNull); ok {
		return true
	}

	if msg, ok := m.any.(json.RawMessage); ok {
		return string(msg) == "null"
	}

	return false
}

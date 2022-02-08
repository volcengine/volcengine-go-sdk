package volcstack

// This File is modify from https://github.com/aws/aws-sdk-go/blob/main/aws/jsonvalue.go

// JSONValue is a representation of a grab bag type that will be marshaled
// into a json string. This type can be used just like any other map.
//
//	Example:
//
//	values := volcstack.JSONValue{
//		"Foo": "Bar",
//	}
//	values["Baz"] = "Qux"
type JSONValue map[string]interface{}

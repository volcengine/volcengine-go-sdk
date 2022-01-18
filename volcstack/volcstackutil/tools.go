package volcstackutil

import (
	"reflect"
	"strconv"
	"strings"
)

func GetSdkValue(keyPattern string, obj interface{}) (interface{}, error) {
	keys := strings.Split(keyPattern, ".")
	root := obj
	for _, k := range keys {
		if reflect.ValueOf(root).Kind() == reflect.Map {
			root = root.(map[string]interface{})[k]
			if root == nil {
				return root, nil
			}

		} else if reflect.ValueOf(root).Kind() == reflect.Slice {
			i, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			if len(root.([]interface{})) < i {
				return nil, nil
			}
			root = root.([]interface{})[i]
		}
	}
	return root, nil
}

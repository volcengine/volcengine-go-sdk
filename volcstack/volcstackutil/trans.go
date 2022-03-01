package volcstackutil

import "strings"

func BodyToMap(body string, sensitive []string, enable bool) map[string]interface{} {
	if !enable {
		return nil
	}
	result := make(map[string]interface{})
	params := strings.Split(body, "&")
	for _, param := range params {
		values := strings.Split(param, "=")
		if values[0] == "Action" || values[0] == "Version" {
			continue
		}
		v := values[1]
		if sensitive != nil && len(sensitive) > 0 {
			for _, s := range sensitive {
				if strings.Contains(values[0], s) {
					v = "****"
					break
				}
			}
		}
		result[values[0]] = v
	}
	return result
}

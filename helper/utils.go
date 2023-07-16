package helper

import "encoding/json"

func PrettyPrint(data interface{}) string {
	jsonBytes, _ := json.MarshalIndent(data, "", "  ")
	return string(jsonBytes)
}

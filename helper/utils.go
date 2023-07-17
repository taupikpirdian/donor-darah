package helper

import "encoding/json"

func PrettyPrint(data interface{}) string {
	jsonBytes, _ := json.MarshalIndent(data, "", "  ")
	return string(jsonBytes)
}

func StructToString(data interface{}) string {
	json, _ := json.Marshal(data)
	return string(json)
}

package helper

import (
	"encoding/json"
	"time"
)

func PrettyPrint(data interface{}) string {
	jsonBytes, _ := json.MarshalIndent(data, "", "  ")
	return string(jsonBytes)
}

func StructToString(data interface{}) string {
	json, _ := json.Marshal(data)
	return string(json)
}

func DateStringFormat(date string) string {
	layout := "2006-01-02T15:04:05-07:00"
	t, _ := time.Parse(layout, date)

	// Format the time.Time object to the desired date format
	formattedDate := t.Format("2006-01-02")
	return formattedDate
}

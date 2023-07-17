package cfg

import (
	"log"
	"os"
	"time"
)

type Logger struct {
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
}

type ContentLogger struct {
	Url      string `json:"endpoin"`
	Payload  string `json:"payload"`
	Response string `json:"response"`
}

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func NewLoger(path string) *Logger {
	// for logs
	PATH_LOGS := path
	t := time.Now()
	timeString := t.Format("2006-01-02")
	name := "logs" + "_" + timeString + ".txt"
	logsPath := PATH_LOGS + name

	file, err := os.OpenFile(logsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if os.IsNotExist(err) {
		// make a file
		f, err := os.Create(logsPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "\nINFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "\nWARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "\nERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		WarningLogger: WarningLogger,
		InfoLogger:    InfoLogger,
		ErrorLogger:   ErrorLogger,
	}
}

func MessageLog(url string, payload string, data string) string {
	urlApi := "/api/v1/" + url
	return "\n+" + urlApi + "\n++" + payload + "\n+++" + data
}

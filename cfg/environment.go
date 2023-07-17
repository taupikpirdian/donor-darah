package cfg

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func NewEnvironment() Config {
	// godotenv package
	PATH_UPLOAD := goDotEnvVariable("PATH_IMAGE_UPLOAD")
	PATH_UPLOAD_META := goDotEnvVariable("PATH_IMAGE_UPLOAD_META")
	CONFIG_SMTP_HOST := goDotEnvVariable("CONFIG_SMTP_HOST")
	CONFIG_SMTP_PORT := goDotEnvVariable("CONFIG_SMTP_PORT")
	CONFIG_SENDER_NAME := goDotEnvVariable("CONFIG_SENDER_NAME")
	CONFIG_AUTH_EMAIL := goDotEnvVariable("CONFIG_AUTH_EMAIL")
	CONFIG_AUTH_PASSWORD := goDotEnvVariable("CONFIG_AUTH_PASSWORD")
	DOMAIN := goDotEnvVariable("DOMAIN")
	ADDRESS := goDotEnvVariable("ADDRESS")
	PATH_LOGS := goDotEnvVariable("PATH_LOGS")
	contextTimeOut := 5

	DB_HOST := goDotEnvVariable("DB_HOST")
	DB_PORT := goDotEnvVariable("DB_PORT")
	DB_USER := goDotEnvVariable("DB_USER")
	DB_PASS := goDotEnvVariable("DB_PASS")
	DB_NAME := goDotEnvVariable("DB_NAME")

	portMail, errPortConvert := strconv.Atoi(CONFIG_SMTP_PORT)
	if errPortConvert != nil {
		panic(errPortConvert)
	}
	loggerCustom := NewLoger(PATH_LOGS)

	return Config{
		ADDRESS:              ADDRESS,
		CONTEXT_RTO:          contextTimeOut,
		PATH_UPLOAD:          PATH_UPLOAD,
		PATH_UPLOAD_META:     PATH_UPLOAD_META,
		CONFIG_SMTP_HOST:     CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT:     portMail,
		CONFIG_SENDER_NAME:   CONFIG_SENDER_NAME,
		CONFIG_AUTH_EMAIL:    CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD: CONFIG_AUTH_PASSWORD,
		LOGGER:               *loggerCustom,
		DOMAIN:               DOMAIN,
		PATH_LOGS:            PATH_LOGS,
		DB_HOST:              DB_HOST,
		DB_PORT:              DB_PORT,
		DB_USER:              DB_USER,
		DB_PASS:              DB_PASS,
		DB_NAME:              DB_NAME,
	}
}

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	_articleHttpDelivery "github.com/bxcodec/go-clean-arch/article/delivery/http"
	_articleHttpDeliveryMiddleware "github.com/bxcodec/go-clean-arch/article/delivery/http/middleware"
	_articleRepo "github.com/bxcodec/go-clean-arch/article/repository/mysql"
	_articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	_authorRepo "github.com/bxcodec/go-clean-arch/author/repository/mysql"
	"github.com/bxcodec/go-clean-arch/cfg"
	_regionHttpDelivery "github.com/bxcodec/go-clean-arch/region/delivery/http"
	_regionRepo "github.com/bxcodec/go-clean-arch/region/repository/mysql"
	_regionUcase "github.com/bxcodec/go-clean-arch/region/usecase"
	_userHttpDelivery "github.com/bxcodec/go-clean-arch/user/delivery/http"
	_userRepo "github.com/bxcodec/go-clean-arch/user/repository/mysql"
	_userUcase "github.com/bxcodec/go-clean-arch/user/usecase"

	_notificationHttpDelivery "github.com/bxcodec/go-clean-arch/notification/delivery/http"
	_notificationRepo "github.com/bxcodec/go-clean-arch/notification/repository/mysql"
	_notificationUcase "github.com/bxcodec/go-clean-arch/notification/usecase"

	_donorHttpDelivery "github.com/bxcodec/go-clean-arch/donor/delivery/http"
	_donorRepo "github.com/bxcodec/go-clean-arch/donor/repository/mysql"
	_donorUcase "github.com/bxcodec/go-clean-arch/donor/usecase"

	_serviceMailUser "github.com/bxcodec/go-clean-arch/user/service/mail"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	// for logs
	PATH_LOGS := goDotEnvVariable("PATH_LOGS")
	t := time.Now()
	timeString := t.Format("2006-01-02")

	files, _ := ioutil.ReadDir(PATH_LOGS)
	countFile := len(files)
	s2 := strconv.Itoa(countFile)
	name := timeString + "_" + s2 + "_" + "logs.txt"

	logsPath := PATH_LOGS + name
	// make a file
	f, err := os.Create(logsPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	file, err := os.OpenFile(logsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
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
	contextTimeOut := 5

	portMail, errPortConvert := strconv.Atoi(CONFIG_SMTP_PORT)
	if errPortConvert != nil {
		panic(errPortConvert)
	}
	loggerCustom := cfg.Logger{
		InfoLogger:    InfoLogger,
		WarningLogger: WarningLogger,
		ErrorLogger:   ErrorLogger,
	}

	cfg := cfg.Config{
		PATH_UPLOAD:          PATH_UPLOAD,
		PATH_UPLOAD_META:     PATH_UPLOAD_META,
		CONFIG_SMTP_HOST:     CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT:     portMail,
		CONFIG_SENDER_NAME:   CONFIG_SENDER_NAME,
		CONFIG_AUTH_EMAIL:    CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD: CONFIG_AUTH_PASSWORD,
		LOGGER:               loggerCustom,
		DOMAIN:               DOMAIN,
	}

	DB_HOST := goDotEnvVariable("DB_HOST")
	DB_PORT := goDotEnvVariable("DB_PORT")
	DB_USER := goDotEnvVariable("DB_USER")
	DB_PASS := goDotEnvVariable("DB_PASS")
	DB_NAME := goDotEnvVariable("DB_NAME")

	dbHost := DB_HOST
	dbPort := DB_PORT
	dbUser := DB_USER
	dbPass := DB_PASS
	dbName := DB_NAME

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	middL := _articleHttpDeliveryMiddleware.InitMiddleware()
	e.Use(middL.CORS)
	authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	ar := _articleRepo.NewMysqlArticleRepository(dbConn)

	timeoutContext := time.Duration(contextTimeOut) * time.Second
	au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext, cfg)
	_articleHttpDelivery.NewArticleHandler(e, au)

	/*
		service regions
	*/
	repoRegion := _regionRepo.NewMysqlRegionRepository(dbConn)
	uCaseRegion := _regionUcase.NewRegionUsecase(repoRegion, timeoutContext, cfg)
	_regionHttpDelivery.NewRegionHandler(e, uCaseRegion)

	/*
		service notification
	*/
	repoNotification := _notificationRepo.NewMysqlNotificationRepository(dbConn)
	uCaseNotification := _notificationUcase.NewNotificationUsecase(repoNotification, timeoutContext, cfg)
	_notificationHttpDelivery.NewNotificationHandler(e, uCaseNotification)

	/*
		service donor
	*/
	repoDonor := _donorRepo.NewMysqlDonorRepository(dbConn)
	uCaseDonor := _donorUcase.NewDonorUsecase(repoDonor, timeoutContext, cfg)
	_donorHttpDelivery.NewDonorHandler(e, uCaseDonor, cfg)

	/*
		service users
	*/
	repoUser := _userRepo.NewMysqlUserRepository(dbConn)
	serviceMail := _serviceMailUser.NewMailService(cfg)
	uCaseUser := _userUcase.NewUserUsecase(repoUser, serviceMail, timeoutContext, repoDonor, cfg)
	_userHttpDelivery.NewUserHandler(e, uCaseUser, cfg)

	cfg.LOGGER.InfoLogger.Println("Starting the application..." + ADDRESS)
	log.Fatal(e.Start(ADDRESS)) //nolint
}

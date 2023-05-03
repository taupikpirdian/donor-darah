package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_articleHttpDelivery "github.com/bxcodec/go-clean-arch/article/delivery/http"
	_articleHttpDeliveryMiddleware "github.com/bxcodec/go-clean-arch/article/delivery/http/middleware"
	_articleRepo "github.com/bxcodec/go-clean-arch/article/repository/mysql"
	_articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	_authorRepo "github.com/bxcodec/go-clean-arch/author/repository/mysql"
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

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

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

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	_articleHttpDelivery.NewArticleHandler(e, au)

	/*
		service users
	*/
	repoUser := _userRepo.NewMysqlUserRepository(dbConn)
	serviceMail := _serviceMailUser.NewMailService()
	uCaseUser := _userUcase.NewUserUsecase(repoUser, serviceMail, timeoutContext)
	_userHttpDelivery.NewUserHandler(e, uCaseUser)

	/*
		service regions
	*/
	repoRegion := _regionRepo.NewMysqlRegionRepository(dbConn)
	uCaseRegion := _regionUcase.NewRegionUsecase(repoRegion, timeoutContext)
	_regionHttpDelivery.NewRegionHandler(e, uCaseRegion)

	/*
		service notification
	*/
	repoNotification := _notificationRepo.NewMysqlNotificationRepository(dbConn)
	uCaseNotification := _notificationUcase.NewNotificationUsecase(repoNotification, timeoutContext)
	_notificationHttpDelivery.NewNotificationHandler(e, uCaseNotification)

	/*
		service donor
	*/
	repoDonor := _donorRepo.NewMysqlDonorRepository(dbConn)
	uCaseDonor := _donorUcase.NewDonorUsecase(repoDonor, timeoutContext)
	_donorHttpDelivery.NewDonorHandler(e, uCaseDonor)

	log.Fatal(e.Start(viper.GetString("server.address"))) //nolint
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

func main() {
	config := cfg.NewEnvironment()
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)
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

	timeoutContext := time.Duration(config.CONTEXT_RTO) * time.Second
	au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext, config)
	_articleHttpDelivery.NewArticleHandler(e, au)

	/*
		service regions
	*/
	repoRegion := _regionRepo.NewMysqlRegionRepository(dbConn)
	uCaseRegion := _regionUcase.NewRegionUsecase(repoRegion, timeoutContext, config)
	_regionHttpDelivery.NewRegionHandler(e, uCaseRegion)

	/*
		service notification
	*/
	repoNotification := _notificationRepo.NewMysqlNotificationRepository(dbConn)
	uCaseNotification := _notificationUcase.NewNotificationUsecase(repoNotification, timeoutContext, config)
	_notificationHttpDelivery.NewNotificationHandler(e, uCaseNotification)

	/*
		service donor
	*/
	repoDonor := _donorRepo.NewMysqlDonorRepository(dbConn)
	uCaseDonor := _donorUcase.NewDonorUsecase(repoDonor, timeoutContext, config)
	_donorHttpDelivery.NewDonorHandler(e, uCaseDonor, config)

	/*
		service users
	*/
	repoUser := _userRepo.NewMysqlUserRepository(dbConn)
	serviceMail := _serviceMailUser.NewMailService(config)
	uCaseUser := _userUcase.NewUserUsecase(repoUser, serviceMail, timeoutContext, repoDonor, config)
	_userHttpDelivery.NewUserHandler(e, uCaseUser, config)

	config.LOGGER.InfoLogger.Println("Starting the application..." + config.ADDRESS)
	log.Fatal(e.Start(config.ADDRESS))
}

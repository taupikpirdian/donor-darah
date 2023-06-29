package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"

	_donorRepo "github.com/bxcodec/go-clean-arch/donor/repository/mysql"
	_notificationRepo "github.com/bxcodec/go-clean-arch/notification/repository/mysql"
	_userRepo "github.com/bxcodec/go-clean-arch/user/repository/mysql"
)

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

	/*
		cron
	*/
	// set scheduler berdasarkan zona waktu sesuai kebutuhan
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	// stop scheduler tepat sebelum fungsi berakhir
	defer scheduler.Stop()

	// set task yang akan dijalankan scheduler
	// gunakan crontab string untuk mengatur jadwal
	scheduler.AddFunc("0 8 * * *", func() { SendNotification(dbConn) })

	// start scheduler
	go scheduler.Start()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func SendNotification(dbConn *sql.DB) {
	// ... instruksi untuk mengirim automail berdasarkan automailType
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " Cron SendNotification telah dijalankan.\n")
	ctx := context.Background()
	t := time.Now().Format("2006-01-02")
	tString := time.Now().Format("02 January 2006")
	/*
		service users
	*/
	repoUser := _userRepo.NewMysqlUserRepository(dbConn)
	/*
		service donor
	*/
	repoDonor := _donorRepo.NewMysqlDonorRepository(dbConn)
	/*
		service notification
	*/
	repoNotification := _notificationRepo.NewMysqlNotificationRepository(dbConn)

	// cari data user aktif
	listUser, errUser := repoUser.GetListUser(ctx)
	if errUser != nil {
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " Error GetListUser.\n")
	}
	for _, user := range listUser {
		idString := user.Id
		id, _ := strconv.ParseInt(idString, 10, 64)
		// // loop user dan cari data register donor terakhir user tersebut
		lastDonor, _ := repoDonor.LastDonorByStatus(ctx, id, "DONE")
		if lastDonor != nil {
			// check tanggalnya sudah lewat 1 bulan 28 hari atau belum
			tFrom := lastDonor.CreatedAt.Format("2006-01-02")
			dur := duration(t, tFrom)
			if dur == -58 {
				// jika sudah kirim notif
				title := "Jadwal Donor Darah"
				msg := "Anda memiliki jadwal donor darah selanjutnya pada " + tString + ". Mohon untuk segera mengajukan donor darah dan berpartisipasi dalam kegiatan donor darah tersebut. Ajukan Donor Darah Sekarang"
				repoNotification.CreateNotification(ctx, title, msg, id)
				fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " - Success Create Notif.\n")
			}
		}
	}
}

func duration(dateFrom string, dateTo string) int {
	layout := "2006-01-02"

	date1, err := time.Parse(layout, dateFrom)
	if err != nil {
		fmt.Println("Error parsing date string 1:", err)
	}

	date2, err := time.Parse(layout, dateTo)
	if err != nil {
		fmt.Println("Error parsing date string 2:", err)
	}

	duration := date2.Sub(date1)
	days := int(duration.Hours() / 24)

	return days
}

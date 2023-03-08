package mysql

import (
	"database/sql"

	"github.com/bxcodec/go-clean-arch/domain"
)

type mysqlNotificationRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlNotificationRepository(conn *sql.DB) domain.NotificationRepository {
	return &mysqlNotificationRepository{conn}
}

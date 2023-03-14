package mysql

import (
	"database/sql"

	"github.com/bxcodec/go-clean-arch/domain"
)

type mysqlDonorRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlDonorRepository(conn *sql.DB) domain.DonorRepository {
	return &mysqlDonorRepository{conn}
}

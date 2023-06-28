package http_request

import (
	"time"
)

type BodyBloodStock struct {
	Title                string                  `json:"title" validate:"required"`
	BodyBloodStockDetail []*BodyBloodStockDetail `json:"detail_stock" validate:"required"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type BodyBloodStockDetail struct {
	Title     string `json:"title" validate:"required"`
	Stock     int64  `json:"stock" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
package domain

import (
	"context"
	"time"
)

type District struct {
	Id        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type DistrictData struct {
	id        int64
	code      string
	name      string
	updatedAt time.Time
	createdAt time.Time
}

// UserUsecase represent the user's usecases
type RegionUsecase interface {
	GetDistrict(ctx context.Context) ([]*DistrictData, error)
}

// UserRepository represent the user's repository contract
type RegionRepository interface {
	GetDistrict(ctx context.Context) ([]*DistrictData, error)
}

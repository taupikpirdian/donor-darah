package domain

import "time"

type UnitDTO struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type Unit struct {
	id        int64
	name      string
	updatedAt time.Time
	createdAt time.Time
}

package domain

import (
	"time"
)

type DonorSchedulle struct {
	id            int64
	unitId        int64
	placeName     string
	address       string
	date          time.Time
	timeStart     time.Time
	timeEnd       time.Time
	typeSchedulle string
	updatedAt     time.Time
	createdAt     time.Time
}

type DonorSchedulleDTO struct {
	Id            int64
	UnitId        int64
	PlaceName     string
	Address       string
	Date          time.Time
	TimeStart     string
	TimeEnd       string
	TypeSchedulle string
	UpdatedAt     time.Time
	CreatedAt     time.Time
}

func NewDonorSchedule(req DonorSchedulleDTO) (*DonorSchedulle, error) {

	schedule := &DonorSchedulle{
		id:        req.Id,
		unitId:    req.UnitId,
		placeName: req.PlaceName,
		address:   req.Address,
		date:      req.Date,
		updatedAt: time.Now(),
		createdAt: time.Now(),
	}

	return schedule, nil
}

func (drs *DonorSchedulle) GetId_DonorSchedule() int64 {
	return drs.id
}

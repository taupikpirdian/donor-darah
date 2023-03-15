package domain

import "time"

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

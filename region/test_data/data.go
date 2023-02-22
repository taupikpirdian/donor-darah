package testdata

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/bxcodec/go-clean-arch/domain"
)

func MultipleDistrict() []*domain.DistrictData {
	d := make([]*domain.DistrictData, 0)
	data := &domain.DistrictData{
		Id:        1,
		Code:      faker.Word(),
		Name:      faker.Word(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	d = append(d, data)
	return d
}

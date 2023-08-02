package model

import (
	"time"

	modelUser "github.com/bxcodec/go-clean-arch/user/repository/model"
)

type DonorRegisterModel struct {
	Id                         int64
	Code                       string
	UserId                     int64
	DonorSchedulleId           int64
	IsApprove                  bool
	StatusApprove              string
	DonorProof                 string
	Status                     string
	DonorSchedulle             DonorSchedulleModel
	User                       modelUser.UserModel
	Unit                       UnitModel
	DonorRegisterQuestionerDTO []*DonorRegisterQuestionerModel
	UpdatedAt                  time.Time
	CreatedAt                  time.Time
}

type DonorSchedulleModel struct {
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

type UnitModel struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type DonorRegisterQuestionerModel struct {
	Id           int64
	CodeQuestion string
	Answer       string
}

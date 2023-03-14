package domain

import (
	"context"
	"time"
)

type DonorRegisterDTO struct {
	Id               int64                       `json:"id"`
	Code             string                      `json:"code"`
	UserId           int64                       `json:"userId"`
	DonorSchedulleId int64                       `json:"donorSchedulleId"`
	IsApprove        bool                        `json:"isApprove"`
	DonorProof       bool                        `json:"donorProof"`
	Questioner       *DonorRegisterQuestionerDTO `json:"questioner"`
	UpdatedAt        time.Time                   `json:"updatedAt"`
	CreatedAt        time.Time                   `json:"createdAt"`
}

type DonorRegisterQuestionerDTO struct {
	Id              int64     `json:"id"`
	DonorRegisterId int64     `json:"donorRegisterId"`
	CodeQuestion    string    `json:"codeQuestion"`
	Title           string    `json:"title"`
	Answer          string    `json:"answer"`
	UpdatedAt       time.Time `json:"updatedAt"`
	CreatedAt       time.Time `json:"createdAt"`
}

type DonorRegister struct {
	id               int64
	code             string
	userId           int64
	donorSchedulleId int64
	isApprove        bool
	donorProof       bool
	Questioner       *DonorRegisterQuestioner
	updatedAt        time.Time
	createdAt        time.Time
}

type DonorRegisterQuestioner struct {
	id              int64
	donorRegisterId int64
	codeQuestion    string
	title           string
	answer          string
	updatedAt       time.Time
	createdAt       time.Time
}

type DonorUsecase interface {
	DonorRegister(ctx context.Context, userId int64, dto *DonorRegisterDTO) error
}

// UserRepository represent the user's repository contract
type DonorRepository interface {
}

package domain

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type RequestRegisterDonor struct {
	IdSchedulle int64                `json:"idSchedulle" validate:"required"`
	ListAnswer  []*RequestListAnswer `json:"listAnswer" validate:"required"`
}

type RequestListAnswer struct {
	QuestionCode string `json:"questionCode"`
	Title        string `json:"title"`
	Answer       string `json:"answer"`
}

type DonorRegisterDTO struct {
	Id               int64
	Code             string
	UserId           int64
	DonorSchedulleId int64
	IsApprove        bool
	DonorProof       string
	Status           string
	DonorSchedulle   DonorSchedulleDTO
	UpdatedAt        time.Time
	CreatedAt        time.Time
}

type DonorRegister struct {
	id               int64
	code             string
	userId           int64
	donorSchedulleId int64
	isApprove        bool
	donorProof       string
	status           string
	Questioner       []*DonorRegisterQuestioner
	DonorSchedulle   *DonorSchedulle
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
	DonorRegister(ctx context.Context, userId int64, req *RequestRegisterDonor) error
	ListAgenda(ctx context.Context, userId int64) ([]*DonorRegisterDTO, error)
}

// UserRepository represent the user's repository contract
type DonorRepository interface {
	DonorRegister(ctx context.Context, donor *DonorRegister) (int64, error)
	DonorRegisterQuestioner(ctx context.Context, donor *DonorRegisterQuestioner, donorRegisterId int64) error
	ListAgenda(ctx context.Context, userId int64) ([]*DonorRegisterDTO, error)
}

func NewDonorRegister(userId int64, req RequestRegisterDonor) (*DonorRegister, error) {
	if req.IdSchedulle == 0 {
		return nil, errors.New("ID SCHEDULLE NOT SET")
	}

	questioner := make([]*DonorRegisterQuestioner, 0)
	for _, answer := range req.ListAnswer {
		addQuestioner := &DonorRegisterQuestioner{
			codeQuestion: answer.QuestionCode,
			title:        answer.Title,
			answer:       answer.Answer,
			updatedAt:    time.Now(),
			createdAt:    time.Now(),
		}
		questioner = append(questioner, addQuestioner)
	}

	timeNow := time.Now().Format("20060102150405")
	randString := rand.Intn(900) + 100
	randCode := timeNow + strconv.Itoa(randString)

	donorRegister := &DonorRegister{
		code:             randCode,
		userId:           userId,
		donorSchedulleId: req.IdSchedulle,
		isApprove:        true,
		donorProof:       "",
		status:           "OPEN",
		Questioner:       questioner,
		updatedAt:        time.Now(),
		createdAt:        time.Now(),
	}

	return donorRegister, nil
}

func (dr *DonorRegister) GetId_DonorRegister() int64 {
	return dr.id
}

func (dr *DonorRegister) GetCode_DonorRegister() string {
	return dr.code
}

func (dr *DonorRegister) GetUserId_DonorRegister() int64 {
	return dr.userId
}

func (dr *DonorRegister) GetDonorSchedulleId_DonorRegister() int64 {
	return dr.donorSchedulleId
}

func (dr *DonorRegister) GetIsApprove_DonorRegister() bool {
	return dr.isApprove
}

func (dr *DonorRegister) GetDonorProof_DonorRegister() string {
	return dr.donorProof
}

func (dr *DonorRegister) GetQuestion_DonorRegister() []*DonorRegisterQuestioner {
	return dr.Questioner
}

func (dr *DonorRegister) GetStatus_DonorRegister() string {
	return dr.status
}

func (dr *DonorRegister) GetUpdateAt_DonorRegister() time.Time {
	return dr.updatedAt
}

func (dr *DonorRegister) GetCreatedAt_DonorRegister() time.Time {
	return dr.createdAt
}

func (drq *DonorRegisterQuestioner) GetDonorRegisterId_DonorRegister() int64 {
	return drq.donorRegisterId
}

func (drq *DonorRegisterQuestioner) GetCodeQuestion_DonorRegister() string {
	return drq.codeQuestion
}

func (drq *DonorRegisterQuestioner) GetTitle_DonorRegister() string {
	return drq.title
}

func (drq *DonorRegisterQuestioner) GetAnswer_DonorRegister() string {
	return drq.answer
}

func (drq *DonorRegisterQuestioner) GetUpdateAt_DonorRegisterQ() time.Time {
	return drq.updatedAt
}

func (drq *DonorRegisterQuestioner) GetCreatedAt_DonorRegisterQ() time.Time {
	return drq.createdAt
}

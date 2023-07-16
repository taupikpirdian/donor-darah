package domain

import (
	"context"
	"errors"
	"math/rand"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/bxcodec/go-clean-arch/donor/delivery/http_request"
)

type RequestRegisterDonor struct {
	IdSchedule int64                `json:"idSchedule" validate:"required"`
	ListAnswer []*RequestListAnswer `json:"listAnswer" validate:"required"`
}

type RequestListAnswer struct {
	QuestionCode string `json:"questionCode"`
	Title        string `json:"title"`
	Answer       string `json:"answer"`
}

type DonorRegisterDTO struct {
	Id                         int64
	Code                       string
	UserId                     int64
	DonorSchedulleId           int64
	IsApprove                  bool
	StatusApprove              string
	DonorProof                 string
	Status                     string
	DonorSchedulle             DonorSchedulleDTO
	User                       User
	Unit                       UnitDTO
	DonorRegisterQuestionerDTO []*DonorRegisterQuestionerDTO
	UpdatedAt                  time.Time
	CreatedAt                  time.Time
}

type DonorRegisterQuestionerDTO struct {
	Id           int64
	CodeQuestion string
	Answer       string
}

type DonorRegister struct {
	id              int64
	code            string
	userId          int64
	donorScheduleId int64
	isApprove       bool
	donorProof      string
	status          string
	Questioner      []*DonorRegisterQuestioner
	DonorSchedulle  *DonorSchedulle
	updatedAt       time.Time
	createdAt       time.Time
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

type DonorStock struct {
	Id               int64            `json:"id"`
	Title            string           `json:"title"`
	DonorDetailStock DonorDetailStock `json:"-"`
	UpdatedAt        time.Time        `json:"updatedAt"`
	CreatedAt        time.Time        `json:"createdAt"`
}

type DonorDetailStock struct {
	Id        int64
	Title     string
	Stock     int64
	UpdatedAt time.Time
	CreatedAt time.Time
}

type UploadedFile struct {
	FileHeader *multipart.FileHeader `json:"fileHeader" form:"file"`
}

type DonorUsecase interface {
	DonorRegister(ctx context.Context, userId int64, req *RequestRegisterDonor) error
	ListAgenda(ctx context.Context, userId int64) ([]*DonorRegisterDTO, error)
	SingleAgenda(ctx context.Context, id int64) (*DonorRegisterDTO, error)
	ListSchedulle(ctx context.Context, unitId int64) ([]*DonorSchedulleDTO, error)
	ListRiwayat(ctx context.Context, userId int64) ([]*DonorRegisterDTO, error)
	UploadBukti(ctx context.Context, id int64, file *multipart.FileHeader) error
	UploadBuktiView(ctx context.Context, id int64) (*DonorRegisterDTO, error)
	CancelDonor(ctx context.Context, id int64) error
	Reschedule(ctx context.Context, id int64, dto *DonorSchedulleDTO) error
	Card(ctx context.Context, userId int64) (*Card, error)
	StoreStock(ctx context.Context, unitId int64, req *http_request.BodyBloodStock) error
	ListStock(ctx context.Context, unitId int64) ([]*DonorStock, error)
	ListDetailStock(ctx context.Context, stockId int64) ([]*DonorDetailStock, error)
	SchedulleStore(ctx context.Context, req *http_request.SchedulleStore) error
	SchedulleDelete(ctx context.Context, id int64) error
	StockUpdateDonor(ctx context.Context, req *http_request.BodyBloodStock) error
	StockDelete(ctx context.Context, id int64) error
	ListDonorRegister(ctx context.Context) ([]*DonorRegisterDTO, error)
}

// UserRepository represent the user's repository contract
type DonorRepository interface {
	DonorRegister(ctx context.Context, donor *DonorRegister) (int64, error)
	DonorRegisterQuestioner(ctx context.Context, donor *DonorRegisterQuestioner, donorRegisterId int64) error
	ListAgenda(ctx context.Context, userId int64) ([]*DonorRegisterDTO, error)
	SingleAgenda(ctx context.Context, id int64) (*DonorRegisterDTO, error)
	ListSchedulle(ctx context.Context, unitId int64) ([]*DonorSchedulleDTO, error)
	ListRiwayat(ctx context.Context, userId int64) ([]*DonorRegisterDTO, error)
	UploadBukti(ctx context.Context, id int64, path string) error
	UploadBuktiView(ctx context.Context, id int64) (*DonorRegisterDTO, error)
	CancelDonor(ctx context.Context, id int64) error
	FindSchedule(ctx context.Context, dto *DonorSchedulleDTO) (*DonorSchedulleDTO, error)
	Reschedule(ctx context.Context, id int64, dto *DonorSchedulle) error
	GetCard(ctx context.Context, userId int64) (*Card, error)
	ListRiwayatByStatus(ctx context.Context, userId int64, status string) ([]*DonorRegisterDTO, error)
	NextDonorByStatus(ctx context.Context, userId int64, status string) (*DonorRegisterDTO, error)
	LastDonorByStatus(ctx context.Context, userId int64, status string) (*DonorRegisterDTO, error)
	StoreStock(ctx context.Context, unitId int64, title string) (int64, error)
	StoreStockDetail(ctx context.Context, stockId int64, data *http_request.BodyBloodStockDetail) error
	ListStock(ctx context.Context, unitId int64) ([]*DonorStock, error)
	ListDetailStock(ctx context.Context, stockId int64) ([]*DonorDetailStock, error)
	SchedulleStore(ctx context.Context, req *http_request.SchedulleStore) error
	SchedulleDelete(ctx context.Context, id int64) error
	FindDonorRegister(ctx context.Context, id int64) ([]*DonorRegisterDTO, error)
	StockUpdateDonor(ctx context.Context, req *http_request.BodyBloodStock) error
	UpdateStockDetail(ctx context.Context, stockId int64, data *http_request.BodyBloodStockDetail) error
	StockDelete(ctx context.Context, id int64) error
	StockDeleteDetail(ctx context.Context, stockId int64) error
	ListDonorRegister(ctx context.Context) ([]*DonorRegisterDTO, error)
	ListAnswer(ctx context.Context, registerId int64) ([]*DonorRegisterQuestionerDTO, error)
}

func NewDonorRegister(userId int64, req RequestRegisterDonor) (*DonorRegister, error) {
	if req.IdSchedule == 0 {
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
		code:            randCode,
		userId:          userId,
		donorScheduleId: req.IdSchedule,
		isApprove:       true,
		donorProof:      "",
		status:          "OPEN",
		Questioner:      questioner,
		updatedAt:       time.Now(),
		createdAt:       time.Now(),
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
	return dr.donorScheduleId
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

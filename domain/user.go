package domain

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/bxcodec/go-clean-arch/user/delivery/http_request"
	"github.com/bxcodec/go-clean-arch/user/repository/model"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// User is representing the User data struct
type User struct {
	Id                   string       `json:"id"`
	Name                 string       `json:"name" validate:"required"`
	Email                string       `json:"email" validate:"required"`
	Phone                string       `json:"phone" validate:"required"`
	Password             string       `json:"password" validate:"required"`
	PasswordConfirmation string       `json:"passwordConfirmation" validate:"required"`
	JobId                string       `json:"jobId"`
	UnitId               string       `json:"unitId"`
	PlaceOfBirth         string       `json:"placeOfBirth"`
	DateOfBirth          string       `json:"dateOfBirth"`
	Gender               string       `json:"gender"`
	SubDistrictId        string       `json:"subDistrictId"`
	VillageId            string       `json:"villageId"`
	Address              string       `json:"address"`
	PostalCode           string       `json:"postalCode"`
	Role                 string       `json:"role"`
	CreatedAt            time.Time    `json:"-"`
	MemberCode           string       `json:"memberCode"`
	Job                  Job          `json:"job"`
	Unit                 UnitDTO      `json:"unit"`
	SubDistrict          DistrictData `json:"subDistrict"`
	Village              VillageData  `json:"village"`
	Histories            Histories    `json:"-"`
}

type UserData struct {
	id          int64
	name        string
	email       string
	phone       string
	role        string
	password    []byte
	updatedAt   time.Time
	createdAt   time.Time
	profileData ProfileData
	profile     Profile
}

type ProfileData struct {
	id            int64
	code          string
	userId        int64
	jobId         string
	unitId        string
	placeOfBirth  string
	dateOfBirth   time.Time
	gender        string
	subDistrictId string
	villageId     string
	address       string
	postalCode    string
	updatedAt     time.Time
	createdAt     time.Time
}

type Histories struct {
	TotalDonor           int64                  `json:"totalDonor"`
	LastDonor            time.Time              `json:"lastDonor"`
	NextDonor            time.Time              `json:"nextDonor"`
	HistoryDonorRegister []HistoryDonorRegister `json:"donorRegister"`
}

type HistoryDonorRegister struct {
	Id         int64     `json:"id"`
	Code       string    `json:"code"`
	PlaceName  string    `json:"placeName"`
	Address    string    `json:"address"`
	Date       time.Time `json:"schedule_date"`
	TimeStart  string    `json:"timeStart"`
	TimeEnd    string    `json:"timeEnd"`
	Type       string    `json:"type"`
	DonorProof string    `json:"donorProof"`
	CreatedAt  time.Time `json:"createdAt"`
}

type Profile struct {
	Id             int64          `json:"id"`
	Name           string         `json:"name"`
	MemberCode     sql.NullString `json:"memberCode"`
	UrlImage       string         `json:"urlImage"`
	UrlImageFromDB sql.NullString `json:"urlImageFromDb"`
	TotalDonor     int64          `json:"totalDonor"`
	LastDonor      time.Time      `json:"lastDonor"`
	NextDonor      time.Time      `json:"nextDonor"`
	User           *User          `json:"user"`
}

type Job struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type JwtCustomClaims struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type DtoRequestLogin struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Auth struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type Token struct {
	Token string `json:"token"`
}

type Card struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	MemberCode  string    `json:"memberCode"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

// UserUsecase represent the user's usecases
type UserUsecase interface {
	Register(ctx context.Context, us *User) error
	GetJob(ctx context.Context) ([]*Job, error)
	Login(ctx context.Context, us *DtoRequestLogin) (*Auth, error)
	GetUnit(ctx context.Context) ([]*UnitDTO, error)
	ChangePassword(ctx context.Context, us *User, userID int64) error
	ForgotPassword(ctx context.Context, us *User) error
	ListUser(ctx context.Context) ([]*User, error)
	CreatetUser(ctx context.Context, us *User) error
	DeleteUser(ctx context.Context, id string) error
	GetProfile(ctx context.Context, userId int64) (*Profile, error)
	UpdateProfile(ctx context.Context, userId int64, req *http_request.BodyUpdateProfile) error
	DetailUser(ctx context.Context, id string) (*Profile, error)
	RefreshToken(ctx context.Context, userId int64) (*Token, error)
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	Register(ctx context.Context, us *UserData) error
	StoreProfile(ctx context.Context, us *UserData) error
	GetJob(ctx context.Context) ([]*Job, error)
	FindUser(ctx context.Context, us *UserData) (*User, error)
	FindUserById(ctx context.Context, us *UserData) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*UserData, error)
	GetUnit(ctx context.Context) ([]*UnitDTO, error)
	ChangePassword(ctx context.Context, us *UserData) error
	GetListUser(ctx context.Context) ([]*User, error)
	DeleteUser(ctx context.Context, id string) error
	DeleteUserProfil(ctx context.Context, id string) error
	GetProfile(ctx context.Context, userId int64) (*Profile, error)
	UpdateProfile(ctx context.Context, userId int64, req *http_request.BodyUpdateProfile) error
	UpdateUser(ctx context.Context, userId int64, req *http_request.BodyUpdateProfile) error
	GetProfileFull(ctx context.Context, userId int64) (*model.UserModel, error)
	GetJobById(ctx context.Context, id string) (*model.JobModel, error)
	GetUnitById(ctx context.Context, id string) (*model.UnitModel, error)
	GetSubDistrictById(ctx context.Context, id string) (*model.SubDistrictModel, error)
	GetVillageById(ctx context.Context, id string) (*model.VillageModel, error)
}

func NewUser(u *User) (*UserData, error) {
	if u.Name == "" {
		return nil, errors.New("NAME NOT SET")
	}

	if u.Password != u.PasswordConfirmation {
		return nil, errors.New("PASSWORD CONFIRMATION NOT SAME")
	}

	// hash password
	resultHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("ERROR HASHING")
	}

	currentTime := time.Now()
	codeTime := ""
	if u.Role != "admin" {
		codeTime = "DN-" + currentTime.Format("20060102150405") + generateCodeString()
	}
	date, err := time.Parse("2006-01-02", u.DateOfBirth)
	if err != nil {
		return nil, errors.New("FORMAT DATE INCORRECT")
	}

	data := &UserData{
		name:      u.Name,
		role:      u.Role,
		email:     u.Email,
		phone:     u.Phone,
		password:  resultHash,
		updatedAt: time.Now(),
		createdAt: time.Now(),
		profileData: ProfileData{
			code:          codeTime,
			jobId:         u.JobId,
			unitId:        u.UnitId,
			placeOfBirth:  u.PlaceOfBirth,
			dateOfBirth:   date,
			gender:        u.Gender,
			subDistrictId: u.SubDistrictId,
			villageId:     u.VillageId,
			address:       u.Address,
			postalCode:    u.PostalCode,
		},
	}

	return data, nil
}

func NewProfile(u *Profile) (*UserData, error) {
	if u.Name == "" {
		return nil, errors.New("NAME NOT SET")
	}

	rand.Seed(time.Now().UnixNano()) // Set seed with current time

	length := 16 // Set length of random string

	// Characters to use in the random string
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	// Generate random string
	result := make([]byte, length)

	for i := range result {
		if i < 3 {
			result[i] = chars[rand.Intn(len(chars))]
		} else {
			result[i] = digits[rand.Intn(len(digits))]
		}
	}

	return &UserData{
		id:   u.Id,
		name: u.Name,
		profile: Profile{
			MemberCode: sql.NullString{String: string(result), Valid: true},
			UrlImage:   u.UrlImage,
			TotalDonor: u.TotalDonor,
			NextDonor:  u.NextDonor,
			LastDonor:  u.LastDonor,
		},
	}, nil
}

func NewProfileV2(u *Profile, p *model.UserModel, len int, nextDonor time.Time, latsDonor time.Time) *Profile {
	if p.DateOfBirth != "" {
		layout := "2006-01-02T15:04:05-07:00"
		// Parse the date string into a time.Time object
		t, err := time.Parse(layout, p.DateOfBirth)
		if err == nil {
			formattedDate := t.Format("2006-01-02")
			p.DateOfBirth = formattedDate
		}
	}
	var profile = &Profile{}

	/*
		model to entity
	*/
	var user = &User{
		Id:            p.Id,
		Name:          p.Name,
		Email:         p.Email,
		Phone:         p.Password,
		JobId:         p.JobId.String,
		UnitId:        p.UnitId.String,
		PlaceOfBirth:  p.PlaceOfBirth.String,
		DateOfBirth:   p.DateOfBirth,
		Gender:        p.Gender,
		SubDistrictId: p.SubDistrictId,
		VillageId:     p.VillageId,
		Address:       p.Address,
		PostalCode:    p.PostalCode,
		Role:          p.Role,
		MemberCode:    p.MemberCode.String,
	}

	profile = &Profile{
		Id:         u.Id,
		MemberCode: u.MemberCode,
		Name:       u.Name,
		UrlImage:   u.UrlImageFromDB.String,
		TotalDonor: int64(len),
		LastDonor:  latsDonor,
		NextDonor:  nextDonor,
		User:       user,
	}

	return profile
}

func NewProfileV3(u *Profile, p *model.UserModel, len int, nextDonor time.Time, latsDonor time.Time, job *model.JobModel, unit *model.UnitModel, sub_district *model.SubDistrictModel, village *model.VillageModel) *Profile {
	if p.DateOfBirth != "" {
		layout := "2006-01-02T15:04:05-07:00"
		// Parse the date string into a time.Time object
		t, err := time.Parse(layout, p.DateOfBirth)
		if err == nil {
			formattedDate := t.Format("2006-01-02")
			p.DateOfBirth = formattedDate
		}
	}
	var profile = &Profile{}
	unit64, _ := strconv.ParseInt(unit.Id, 10, 64)
	job64, _ := strconv.ParseInt(job.Id, 10, 64)
	dataSubDistrict := DistrictData{}
	if sub_district != nil {
		subDistrict64, _ := strconv.ParseInt(sub_district.Id, 10, 64)
		dataSubDistrict = DistrictData{
			Id:   subDistrict64,
			Code: sub_district.Code,
			Name: sub_district.Name,
		}
	}
	dataVillage := VillageData{}
	if village != nil {
		village64, _ := strconv.ParseInt(village.Id, 10, 64)
		villageSubDistrictId64, _ := strconv.ParseInt(village.SubDistrictId, 10, 64)
		dataVillage = VillageData{
			Id:            village64,
			SubDistrictId: villageSubDistrictId64,
			Code:          village.Code,
			Name:          village.Name,
		}
	}
	/*
		model to entity
	*/
	var user = &User{
		Id:            p.Id,
		Name:          p.Name,
		Email:         p.Email,
		Phone:         p.Phone,
		JobId:         p.JobId.String,
		UnitId:        p.UnitId.String,
		PlaceOfBirth:  p.PlaceOfBirth.String,
		DateOfBirth:   p.DateOfBirth,
		Gender:        p.Gender,
		SubDistrictId: p.SubDistrictId,
		VillageId:     p.VillageId,
		Address:       p.Address,
		PostalCode:    p.PostalCode,
		Role:          p.Role,
		MemberCode:    p.MemberCode.String,
		Unit: UnitDTO{
			Id:   unit64,
			Name: unit.Name,
		},
		Job: Job{
			Id:   job64,
			Name: job.Name,
		},
		SubDistrict: dataSubDistrict,
		Village:     dataVillage,
	}

	profile = &Profile{
		Id:         u.Id,
		MemberCode: u.MemberCode,
		Name:       u.Name,
		UrlImage:   u.UrlImageFromDB.String,
		TotalDonor: int64(len),
		LastDonor:  latsDonor,
		NextDonor:  nextDonor,
		User:       user,
	}
	return profile
}

func (u *User) SetUserList(
	p *model.UserModel,
	job *model.JobModel,
	unit *model.UnitModel,
	sub_district *model.SubDistrictModel,
	village *model.VillageModel,
	len int,
	nextDonor time.Time,
	latsDonor time.Time,
	histories []*DonorRegisterDTO,
) {
	if p.DateOfBirth != "" {
		layout := "2006-01-02T15:04:05-07:00"
		// Parse the date string into a time.Time object
		t, err := time.Parse(layout, p.DateOfBirth)
		if err == nil {
			formattedDate := t.Format("2006-01-02")
			p.DateOfBirth = formattedDate
		}
	}
	unit64, _ := strconv.ParseInt(unit.Id, 10, 64)
	job64, _ := strconv.ParseInt(job.Id, 10, 64)

	districtData := DistrictData{}
	if sub_district != nil {
		subDistrict64, _ := strconv.ParseInt(sub_district.Id, 10, 64)
		districtData = DistrictData{
			Id:   subDistrict64,
			Code: sub_district.Code,
			Name: sub_district.Name,
		}
	}

	villageData := VillageData{}
	if village != nil {
		village64, _ := strconv.ParseInt(village.Id, 10, 64)
		villageSubDistrictId64, _ := strconv.ParseInt(village.SubDistrictId, 10, 64)
		villageData = VillageData{
			Id:            village64,
			SubDistrictId: villageSubDistrictId64,
			Code:          village.Code,
			Name:          village.Name,
		}
	}

	/*
		set history
	*/
	historyRegisters := make([]HistoryDonorRegister, 0)
	for _, value := range histories {
		historyRegister := HistoryDonorRegister{
			Id:         value.Id,
			Code:       value.Code,
			PlaceName:  value.DonorSchedulle.PlaceName,
			Address:    value.DonorSchedulle.Address,
			Date:       value.DonorSchedulle.Date,
			TimeStart:  value.DonorSchedulle.TimeStart,
			TimeEnd:    value.DonorSchedulle.TimeEnd,
			Type:       value.DonorSchedulle.TypeSchedulle,
			DonorProof: value.DonorProof,
			CreatedAt:  value.CreatedAt,
		}
		historyRegisters = append(historyRegisters, historyRegister)
	}

	history := Histories{
		TotalDonor:           int64(len),
		LastDonor:            latsDonor,
		NextDonor:            nextDonor,
		HistoryDonorRegister: historyRegisters,
	}
	/*
		set to entity
	*/
	u.Id = p.Id
	u.Name = p.Name
	u.Email = p.Email
	u.Phone = p.Phone
	u.JobId = p.JobId.String
	u.UnitId = p.UnitId.String
	u.PlaceOfBirth = p.PlaceOfBirth.String
	u.DateOfBirth = p.DateOfBirth
	u.Gender = p.Gender
	u.SubDistrictId = p.SubDistrictId
	u.VillageId = p.VillageId
	u.Address = p.Address
	u.PostalCode = p.PostalCode
	u.Role = p.Role
	u.MemberCode = p.MemberCode.String
	u.Unit = UnitDTO{
		Id:   unit64,
		Name: unit.Name,
	}
	u.Job = Job{
		Id:   job64,
		Name: job.Name,
	}
	u.SubDistrict = districtData
	u.Village = villageData
	u.Histories = history
}

func NewUserLogin(u *DtoRequestLogin) (*UserData, error) {
	if u.Phone == "" {
		return nil, errors.New("PHONE IS REQUIRED")
	}

	return &UserData{
		phone: u.Phone,
	}, nil
}

func NewUser2(userId int64, password string) (*UserData, error) {
	resultHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("ERROR HASHING")
	}

	return &UserData{
		id:       userId,
		password: resultHash,
	}, nil
}

func NewUser3(u *User) (*UserData, error) {
	if u.Name == "" {
		return nil, errors.New("NAME NOT SET")
	}

	idInt, err := strconv.ParseInt(u.Id, 10, 64)
	if err != nil {
		return nil, err
	}

	return &UserData{
		id:        idInt,
		name:      u.Name,
		email:     u.Email,
		phone:     u.Phone,
		password:  []byte(u.Password),
		updatedAt: time.Now(),
		createdAt: time.Now(),
	}, nil
}

func NewUser4(userId int64) *UserData {
	return &UserData{
		id: userId,
	}
}

func SetToken(token string, dataUserDb *model.UserModel) (*Auth, error) {
	if token == "" {
		return nil, errors.New("TOKEN IS REQUIRED")
	}

	fmt.Println(dataUserDb)

	return &Auth{
		Token: token,
		User: User{
			Id:            dataUserDb.Id,
			Name:          dataUserDb.Name,
			Email:         dataUserDb.Email,
			Phone:         dataUserDb.Phone,
			Role:          dataUserDb.Role,
			JobId:         dataUserDb.JobId.String,
			UnitId:        dataUserDb.UnitId.String,
			PlaceOfBirth:  dataUserDb.PlaceOfBirth.String,
			DateOfBirth:   dataUserDb.DateOfBirth,
			Gender:        dataUserDb.Gender,
			SubDistrictId: dataUserDb.SubDistrictId,
			VillageId:     dataUserDb.VillageId,
			Address:       dataUserDb.Address,
			PostalCode:    dataUserDb.PostalCode,
			MemberCode:    dataUserDb.MemberCode.String,
		},
	}, nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (cu *UserData) SetIdNewUser(u *UserData, id int64) {
	u.id = id
}

func (cu *UserData) SetPasswordNew(u *UserData, password string) {
	resultHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.password = resultHash
}

func (cu *UserData) GetIdOnUser() int {
	return int(cu.id)
}

func (cu *UserData) GetRoleOnUser() string {
	return cu.role
}

func (cu *UserData) GetNameOnUser() string {
	return cu.name
}

func (cu *UserData) GetPasswordOnUser() []byte {
	return cu.password
}

func (cu *UserData) GetPlaceOfBirthOnProfile() string {
	return cu.profileData.placeOfBirth
}

func (cu *UserData) GetDateOfBirthOnProfile() time.Time {
	return cu.profileData.dateOfBirth
}

func (cu *UserData) GetCodeOnProfile() string {
	return cu.profileData.code
}

func (cu *UserData) GetGenderOnProfile() string {
	return cu.profileData.gender
}

func (cu *UserData) GetJobIdOnProfile() string {
	return cu.profileData.jobId
}

func (cu *UserData) GetUnitIdOnProfile() string {
	return cu.profileData.unitId
}

func (cu *UserData) GetPhoneOnUser() string {
	return cu.phone
}

func (cu *UserData) GetEmailOnUser() string {
	return cu.email
}

func (cu *UserData) GetAddressOnProfile() string {
	return cu.profileData.address
}

func (cu *UserData) GetSubDistrictIdOnProfile() string {
	return cu.profileData.subDistrictId
}

func (cu *UserData) GetVillageIdOnProfile() string {
	return cu.profileData.villageId
}

func (cu *UserData) GetPostalCodeOnProfile() string {
	return cu.profileData.postalCode
}

func (cu *UserData) GetUpdateAtOnUser() time.Time {
	return cu.updatedAt
}

func (cu *UserData) GetCreatedAtOnUser() time.Time {
	return cu.createdAt
}

func generateCodeString() string {
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator with the current time
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomString := make([]byte, 3)
	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(letters))] // Generate a random character from the set of letters
	}

	return string(randomString)
}

func GenerateCodeStringLen(n int) string {
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator with the current time
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomString := make([]byte, n)
	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(letters))] // Generate a random character from the set of letters
	}

	return string(randomString)
}

func (cu *UserData) SetUserId(id int64) {
	cu.id = id
}

func (cu *UserData) SetName(name string) {
	cu.name = name
}

func (cu *UserData) SetEmail(email string) {
	cu.email = email
}

func (cu *UserData) SetPhone(phone string) {
	cu.phone = phone
}

func (cu *UserData) SetJobId(jobId string) {
	cu.profileData.jobId = jobId
}

func (cu *UserData) SetUnitId(unitId string) {
	cu.profileData.unitId = unitId
}

func (cu *UserData) SetPlaceOfBirth(place string) {
	cu.profileData.placeOfBirth = place
}

func (cu *UserData) SetDateOfBirth(date string) {
	dateFormat, _ := time.Parse("2006-01-02", date)
	cu.profileData.dateOfBirth = dateFormat
}

func (cu *UserData) SetGender(gender string) {
	cu.profileData.gender = gender
}

func (cu *UserData) SetSubDistrictId(subDistrictId string) {
	cu.profileData.subDistrictId = subDistrictId
}

func (cu *UserData) SetVillageId(villageId string) {
	cu.profileData.villageId = villageId
}

func (cu *UserData) SetAddress(address string) {
	cu.profileData.address = address
}

func (cu *UserData) SetPostalCode(postalCode string) {
	cu.profileData.postalCode = postalCode
}

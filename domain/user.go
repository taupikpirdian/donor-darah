package domain

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// User is representing the User data struct
type User struct {
	Id                   string    `json:"id"`
	Name                 string    `json:"name" validate:"required"`
	Email                string    `json:"email" validate:"required"`
	Phone                string    `json:"phone" validate:"required"`
	Password             string    `json:"password" validate:"required"`
	PasswordConfirmation string    `json:"passwordConfirmation" validate:"required"`
	JobId                int64     `json:"jobId"`
	UnitId               int64     `json:"unitId"`
	PlaceOfBirth         string    `json:"placeOfBirth"`
	DateOfBirth          string    `json:"dateOfBirth"`
	Gender               string    `json:"gender"`
	SubDistrictId        int64     `json:"subDistrictId"`
	VillageId            int64     `json:"villageId"`
	Address              string    `json:"address"`
	PostalCode           string    `json:"postalCode"`
	Role                 string    `json:"role"`
	CreatedAt            time.Time `json:"created_at"`
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
	jobId         int64
	unitId        int64
	placeOfBirth  string
	dateOfBirth   time.Time
	gender        string
	subDistrictId int64
	villageId     int64
	address       string
	postalCode    string
	updatedAt     time.Time
	createdAt     time.Time
}
type Profile struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	MemberCode string    `json:"memberCode"`
	UrlImage   string    `json:"urlImage"`
	TotalDonor int64     `json:"totalDonor"`
	LastDonor  time.Time `json:"lastDonor"`
	NextDonor  time.Time `json:"nextDonor"`
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

// UserUsecase represent the user's usecases
type UserUsecase interface {
	Register(ctx context.Context, us *User) error
	GetJob(ctx context.Context) ([]*Job, error)
	Login(ctx context.Context, us *DtoRequestLogin) (*Auth, error)
	GetUnit(ctx context.Context) ([]*UnitDTO, error)
	ChangePassword(ctx context.Context, us *User, userID int64) error
	ForgotPassword(ctx context.Context, us *User) error
	ListUser(ctx context.Context) ([]*User, error)
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
	codeTime := "DN-" + currentTime.Format("20060102150405") + generateCodeString()
	date, _ := time.Parse("2006-01-02", u.DateOfBirth)

	return &UserData{
		name:      u.Name,
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
	}, nil
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
			MemberCode: string(result),
			UrlImage:   u.UrlImage,
			TotalDonor: u.TotalDonor,
			NextDonor:  u.NextDonor,
			LastDonor:  u.LastDonor,
		},
	}, nil

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

func SetToken(token string, dataUserDb *User) (*Auth, error) {
	if token == "" {
		return nil, errors.New("TOKEN IS REQUIRED")
	}

	return &Auth{
		Token: token,
		User: User{
			Id:    dataUserDb.Id,
			Name:  dataUserDb.Name,
			Email: dataUserDb.Email,
			Phone: dataUserDb.Phone,
			Role:  dataUserDb.Role,
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

func (cu *UserData) GetJobIdOnProfile() int64 {
	return cu.profileData.jobId
}

func (cu *UserData) GetUnitIdOnProfile() int64 {
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

func (cu *UserData) GetSubDistrictIdOnProfile() int64 {
	return cu.profileData.subDistrictId
}

func (cu *UserData) GetVillageIdOnProfile() int64 {
	return cu.profileData.subDistrictId
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

func (cu *UserData) SetName(name string) {
	cu.name = name
}

func (cu *UserData) SetEmail(email string) {
	cu.email = email
}

func (cu *UserData) SetPhone(phone string) {
	cu.phone = phone
}

func (cu *UserData) SetJobId(jobId int64) {
	cu.profileData.jobId = jobId
}

func (cu *UserData) SetUnitId(unitId int64) {
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

func (cu *UserData) SetSubDistrictId(subDistrictId int64) {
	cu.profileData.subDistrictId = subDistrictId
}

func (cu *UserData) SetVillageId(villageId int64) {
	cu.profileData.villageId = villageId
}

func (cu *UserData) SetAddress(address string) {
	cu.profileData.address = address
}

func (cu *UserData) SetPostalCode(postalCode string) {
	cu.profileData.postalCode = postalCode
}

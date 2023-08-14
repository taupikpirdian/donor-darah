package model

import (
	"database/sql"
	"time"
)

type UserModel struct {
	Id                   string         `json:"id"`
	Name                 string         `json:"name" validate:"required"`
	Email                string         `json:"email" validate:"required"`
	Phone                string         `json:"phone" validate:"required"`
	Password             string         `json:"password" validate:"required"`
	PasswordConfirmation string         `json:"passwordConfirmation" validate:"required"`
	JobId                sql.NullString `json:"jobId"`
	UnitId               sql.NullString `json:"unitId"`
	PlaceOfBirth         string         `json:"placeOfBirth"`
	DateOfBirth          string         `json:"dateOfBirth"`
	Gender               string         `json:"gender"`
	SubDistrictId        string         `json:"subDistrictId"`
	VillageId            string         `json:"villageId"`
	Address              string         `json:"address"`
	PostalCode           string         `json:"postalCode"`
	Role                 string         `json:"role"`
	CreatedAt            time.Time      `json:"-"`
	MemberCode           sql.NullString `json:"memberCode"`
	UrlImage             sql.NullString `json:"urlImage"`
}

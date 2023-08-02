package http_request

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type BodyUpdateProfile struct {
	Name          string
	PlaceOfBirth  string
	DateOfBirth   string
	Gender        string
	JobId         string
	UnitId        string
	Phone         string
	Email         string
	Address       string
	SubDistrictId string
	VillageId     string
	PostalCode    string
	Path          string
	File          *multipart.FileHeader
}

func OrderFilterRequest(c echo.Context) (*BodyUpdateProfile, error) {
	name := c.FormValue("name")
	placeOfBirth := c.FormValue("placeOfBirth")
	dateOfBirth := c.FormValue("dateOfBirth")
	gender := c.FormValue("gender")
	jobId := c.FormValue("jobId")
	unitId := c.FormValue("unitId")
	phone := c.FormValue("phone")
	email := c.FormValue("email")
	address := c.FormValue("address")
	subDistrictId := c.FormValue("subDistrictId")
	villageId := c.FormValue("villageId")
	postalCode := c.FormValue("postalCode")

	file, err := c.FormFile("file")
	if err != nil {
		if err.Error() == "http: no such file" {
			file = nil
		} else {
			return nil, err
		}
	}

	return &BodyUpdateProfile{
		Name:          name,
		PlaceOfBirth:  placeOfBirth,
		DateOfBirth:   dateOfBirth,
		Gender:        gender,
		JobId:         jobId,
		UnitId:        unitId,
		Phone:         phone,
		Email:         email,
		Address:       address,
		SubDistrictId: subDistrictId,
		VillageId:     villageId,
		PostalCode:    postalCode,
		File:          file,
	}, nil
}

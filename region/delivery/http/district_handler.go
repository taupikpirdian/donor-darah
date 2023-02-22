package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/region/delivery/http_response"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// UserHandler  represent the httphandler for user
type DistrictHandler struct {
	AUsecase domain.RegionUsecase
}

// NewUserHandler will initialize the users/ resources endpoint
func NewDistrictHandler(e *echo.Echo, us domain.RegionUsecase) {
	handler := &DistrictHandler{
		AUsecase: us,
	}
	e.GET("/api/v1/district", handler.DistrictGet)
}

func isRequestValid(m *domain.District) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Register will store the user by given request body
func (a *DistrictHandler) DistrictGet(c echo.Context) (err error) {
	var district domain.District
	err = c.Bind(&district)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, http_response.Status{Code: 1, Message: err.Error()})
	}

	var ok bool
	if ok, err = isRequestValid(&district); !ok {
		return c.JSON(http.StatusBadRequest, http_response.Status{Code: 1, Message: err.Error()})
	}

	ctx := c.Request().Context()
	data, errUc := a.AUsecase.GetDistrict(ctx)
	if errUc != nil {
		return c.JSON(getStatusCode(err), http_response.Status{Code: 1, Message: errUc.Error()})
	}

	responseSuccess, _ := http_response.MapResponseDistrict(0, "berhasil get data", data)
	return c.JSON(http.StatusCreated, responseSuccess)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

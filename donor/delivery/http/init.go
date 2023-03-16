package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	echojwt "github.com/labstack/echo-jwt/v4"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// UserHandler  represent the httphandler for user
type DonorHandler struct {
	AUsecase domain.DonorUsecase
}

func isRequestValid_Register(m *domain.RequestRegisterDonor) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// NewUserHandler will initialize the users/ resources endpoint
func NewDonorHandler(e *echo.Echo, us domain.DonorUsecase) {
	handler := &DonorHandler{
		AUsecase: us,
	}
	jwtKey := viper.GetString(`jwt.key`)

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(domain.JwtCustomClaims)
		},
		SigningKey: []byte(jwtKey),
	}
	r := e.Group("/api/v1/donor/")
	r.Use(echojwt.WithConfig(config))
	// list routes
	r.POST("questionnaire", handler.DonorRegister)
	r.GET("agenda", handler.ListAgenda)
	r.GET("agenda/:id", handler.SingleAgenda)
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

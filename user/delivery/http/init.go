package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// UserHandler  represent the httphandler for user
type UserHandler struct {
	AUsecase domain.UserUsecase
	cfg      cfg.Config
}

// NewUserHandler will initialize the users/ resources endpoint
func NewUserHandler(e *echo.Echo, us domain.UserUsecase, cfg cfg.Config) {
	handler := &UserHandler{
		AUsecase: us,
		cfg:      cfg,
	}
	jwtKey := viper.GetString(`jwt.key`)

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(domain.JwtCustomClaims)
		},
		SigningKey: []byte(jwtKey),
	}
	r := e.Group("/api/v1/")
	r.Use(echojwt.WithConfig(config))

	e.POST("/api/v1/register", handler.Register)
	e.GET("/api/v1/job", handler.JobController)
	e.POST("/api/v1/login", handler.LoginController)
	e.GET("/api/v1/unit", handler.UnitController)
	e.POST("/api/v1/forgot-password", handler.ForgotPasswordController)
	// route must auth
	r.POST("change-password", handler.ChangePasswordController)
	r.GET("profile", handler.ProfileController)
	r.POST("profile/update", handler.ProfileUpdateController)
	r.POST("token/refresh", handler.TokenRefreshController)

	g := e.Group("/api/v1/admin/")
	g.Use(echojwt.WithConfig(config))
	g.GET("user/list", handler.AdminUserListController)
	g.POST("user/create", handler.AdminUserCreateController)
	g.DELETE("user/delete/:id", handler.AdminUserDeleteController)
	g.GET("user/detail/:id", handler.AdminUserDetailController)
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

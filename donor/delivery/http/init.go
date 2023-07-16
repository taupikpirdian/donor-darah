package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/cfg"
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
	cfg      cfg.Config
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
func NewDonorHandler(e *echo.Echo, us domain.DonorUsecase, cfg cfg.Config) {
	handler := &DonorHandler{
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
	r := e.Group("/api/v1/donor/")
	r.Use(echojwt.WithConfig(config))
	// list routes
	r.POST("questionnaire", handler.DonorRegister)
	r.GET("agenda", handler.ListAgenda)
	r.GET("agenda/:id", handler.SingleAgenda)
	r.GET("schedule/list/:unitId", handler.ListSchedulle)
	r.GET("riwayat", handler.ListRiwayat)
	r.POST("upload/:donorRegisterId", handler.UploadBukti)
	r.GET("upload-view/:donorRegisterId", handler.UploadBuktiView)
	r.PUT("cancel/:donorRegisterId", handler.CancelDonor)
	r.POST("reschedule/:donorRegisterId", handler.RescheduleDonor)
	r.GET("card", handler.Card)
	r.GET("stock/:unitId", handler.ListStockDonor)
	r.GET("stock-detail/:stockId", handler.ListDetailStockDonor)

	g := e.Group("/api/v1/admin/donor/")
	g.Use(echojwt.WithConfig(config))
	g.POST("stock/:unitId", handler.StockDonor)
	g.POST("stock-update/:unitId", handler.StockUpdateDonor)
	g.DELETE("stock/:id", handler.DeleteStock)
	g.POST("schedulle", handler.SchedulleStore)
	g.DELETE("schedulle/:id", handler.SchedulleDelete)
	g.GET("register", handler.DonorRegisterList)
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

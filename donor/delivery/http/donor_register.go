package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) DonorRegister(c echo.Context) (err error) {
	var donor domain.DonorRegisterDTO
	err = c.Bind(&donor)
	if err != nil {
		responseError, _ := http_response.MapResponseDonorRegister(1, err.Error())
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	var ok bool
	if ok, err = isRequestValid_Register(&donor); !ok {
		responseError2, _ := http_response.MapResponseDonorRegister(1, err.Error())
		return c.JSON(http.StatusBadRequest, responseError2)
	}

	ctx := c.Request().Context()
	// data user by token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	userId := claims.Id

	errUc := d.AUsecase.DonorRegister(ctx, userId, &domain.DonorRegisterDTO{})
	if errUc != nil {
		responseError3, _ := http_response.MapResponseDonorRegister(1, domain.ErrBadBody.Error())
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseDonorRegister(0, "success")
	return c.JSON(http.StatusOK, responseSuccess)
}

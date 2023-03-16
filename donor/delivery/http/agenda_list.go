package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) ListAgenda(c echo.Context) (err error) {
	ctx := c.Request().Context()
	// data user by token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)
	userId := claims.Id

	datas, errUc := d.AUsecase.ListAgenda(ctx, userId)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseListAgenda(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseListAgenda(0, "success", datas)
	return c.JSON(http.StatusOK, responseSuccess)
}

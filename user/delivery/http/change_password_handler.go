package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

)

func (a *UserHandler) ChangePasswordController(c echo.Context) (err error) {
	var user domain.User

	err = c.Bind(&user)
	if err != nil {
		responseError, _ := http_response.MapResponse(1, err.Error())
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	ctx := c.Request().Context()

	// data user by token
	userLogin := c.Get("user").(*jwt.Token)
	claims := userLogin.Claims.(*domain.JwtCustomClaims)
	userId := claims.Id

	err = a.AUsecase.ChangePassword(ctx, &user, userId)

	return
}

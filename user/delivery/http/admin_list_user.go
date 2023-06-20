package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) AdminUserListController(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var user domain.User

	err = c.Bind(&user)
	if err != nil {
		responseError, _ := http_response.MapResponse(1, err.Error())
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	// data user by token
	userLogin := c.Get("user").(*jwt.Token)
	claims := userLogin.Claims.(*domain.JwtCustomClaims)
	role := claims.Role
	if role != "admin" {
		responseError, _ := http_response.MapResponse(1, "only admin can access")
		return c.JSON(http.StatusUnauthorized, responseError)
	}

	datas, errUsecase := a.AUsecase.ListUser(ctx)
	if errUsecase != nil {
		responseErrorUsecase, _ := http_response.MapResponse(1, errUsecase.Error())
		return c.JSON(getStatusCode(err), responseErrorUsecase)
	}

	responseSuccess, _ := http_response.MapResponseUserList(0, "success", datas)
	return c.JSON(http.StatusOK, responseSuccess)
}

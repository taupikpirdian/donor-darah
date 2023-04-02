package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) ForgotPasswordController(c echo.Context) (err error) {
	var user domain.User

	err = c.Bind(&user)
	if err != nil {
		responseError, _ := http_response.MapResponse(1, err.Error())
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	ctx := c.Request().Context()

	err = a.AUsecase.ForgotPassword(ctx, &user)
	if err != nil {
		responseErrorUsecase, _ := http_response.MapResponse(1, domain.ErrBadBody.Error())
		return c.JSON(getStatusCode(err), responseErrorUsecase)
	}

	responseSuccess, _ := http_response.MapResponse(0, "success")
	return c.JSON(http.StatusCreated, responseSuccess)
}

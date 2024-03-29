package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) UnitController(c echo.Context) (err error) {
	ctx := c.Request().Context()
	data, err := a.AUsecase.GetUnit(ctx)
	if err != nil {
		responseErrorUsecase, _ := http_response.MapResponse(1, domain.ErrBadBody.Error())
		return c.JSON(getStatusCode(err), responseErrorUsecase)
	}

	responseSuccess, _ := http_response.MapResponseUnit(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}

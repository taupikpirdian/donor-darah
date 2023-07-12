package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/labstack/echo/v4"
)

func (a *UserHandler) JobController(c echo.Context) (err error) {
	ctx := c.Request().Context()
	data, err := a.AUsecase.GetJob(ctx)
	if err != nil {
		a.cfg.LOGGER.ErrorLogger.Println(err)
		responseErrorUsecase, _ := http_response.MapResponse(1, domain.ErrBadBody.Error())
		return c.JSON(getStatusCode(err), responseErrorUsecase)
	}

	a.cfg.LOGGER.InfoLogger.Println(data)
	responseSuccess, _ := http_response.MapResponseJob(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}

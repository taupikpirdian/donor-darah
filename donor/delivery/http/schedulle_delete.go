package http

import (
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) SchedulleDelete(c echo.Context) (err error) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseErrorConv, _ := http_response.MapResponseGeneral(1, err.Error(), nil)
		return c.JSON(getStatusCode(err), responseErrorConv)
	}
	id := int64(idP)
	ctx := c.Request().Context()

	errUc := d.AUsecase.SchedulleDelete(ctx, id)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseGeneral(1, errUc.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseGeneral(0, "success", nil)
	return c.JSON(http.StatusOK, responseSuccess)
}

package http

import (
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) ListSchedulle(c echo.Context) (err error) {
	idP, err := strconv.Atoi(c.Param("unitId"))
	if err != nil {
		responseErrorConv, _ := http_response.MapResponseSingleAgenda(1, err.Error(), nil)
		return c.JSON(getStatusCode(err), responseErrorConv)
	}
	id := int64(idP)
	ctx := c.Request().Context()

	datas, errUc := d.AUsecase.ListSchedulle(ctx, id)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseListAgenda(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseSchedulleList(0, "success", datas)
	return c.JSON(http.StatusOK, responseSuccess)
}

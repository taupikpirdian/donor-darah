package http

import (
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_request"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) StockDonor(c echo.Context) (err error) {
	var req http_request.BodyBloodStock

	err = c.Bind(&req)
	if err != nil {
		responseError, _ := http_response.MapResponseDonorRegister(1, err.Error())
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	idP, err := strconv.Atoi(c.Param("unitId"))
	if err != nil {
		responseErrorConv, _ := http_response.MapResponseSingleAgenda(1, err.Error(), nil)
		return c.JSON(getStatusCode(err), responseErrorConv)
	}
	id := int64(idP)
	ctx := c.Request().Context()

	errUc := d.AUsecase.StoreStock(ctx, id, &req)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseGeneral(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseGeneral(0, "success", nil)
	return c.JSON(http.StatusOK, responseSuccess)
}

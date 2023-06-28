package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_request"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) StockUpdateDonor(c echo.Context) (err error) {
	var req http_request.BodyBloodStock
	err = c.Bind(&req)
	if err != nil {
		responseError, _ := http_response.MapResponseGeneral(1, err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	ctx := c.Request().Context()

	errUc := d.AUsecase.StockUpdateDonor(ctx, &req)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseGeneral(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseGeneral(0, "success", nil)
	return c.JSON(http.StatusOK, responseSuccess)
}

package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) DonorRegisterList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	datas, errUc := d.AUsecase.ListDonorRegister(ctx)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseGeneral(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseListDonorRegister(0, "success", datas)
	return c.JSON(http.StatusOK, responseSuccess)
}

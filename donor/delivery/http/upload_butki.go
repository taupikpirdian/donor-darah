package http

import (
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) UploadBukti(c echo.Context) (err error) {
	idP, err := strconv.Atoi(c.Param("donorRegisterId"))
	if err != nil {
		responseErrorConv, _ := http_response.MapResponseBuktiDonor(1, err.Error(), nil)
		return c.JSON(getStatusCode(err), responseErrorConv)
	}
	id := int64(idP)
	ctx := c.Request().Context()

	errUc := d.AUsecase.UploadBukti(ctx, id)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseBuktiDonor(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseBuktiDonor(0, "success", nil)
	return c.JSON(http.StatusOK, responseSuccess)
}

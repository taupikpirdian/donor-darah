package http

import (
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/cfg"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/bxcodec/go-clean-arch/helper"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) UploadBuktiView(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(d.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/donor/upload-view/:donorRegisterId",
	}

	idP, err := strconv.Atoi(c.Param("donorRegisterId"))
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseErrorConv, _ := http_response.MapResponseBuktiDonorView(1, err.Error(), nil)
		return c.JSON(getStatusCode(err), responseErrorConv)
	}
	id := int64(idP)
	ctx := c.Request().Context()

	data, errUc := d.AUsecase.UploadBuktiView(ctx, id)
	if errUc != nil {
		loggerFile.ErrorLogger.Println(errUc)
		responseError3, _ := http_response.MapResponseBuktiDonorView(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	data.DonorProof = d.cfg.DOMAIN + data.DonorProof

	contentLog.Payload = helper.StructToString(id)
	contentLog.Response = helper.StructToString(data)
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseBuktiDonorView(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}

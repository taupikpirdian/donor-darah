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

func (d *DonorHandler) UploadBukti(c echo.Context) (err error) {
	loggerFile := cfg.NewLoger(d.cfg.PATH_LOGS)
	contentLog := cfg.ContentLogger{
		Url: "/api/v1/donor/upload/:donorRegisterId",
	}

	idP, err := strconv.Atoi(c.Param("donorRegisterId"))
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		responseErrorConv, _ := http_response.MapResponseBuktiDonor(1, err.Error(), nil)
		return c.JSON(getStatusCode(err), responseErrorConv)
	}
	id := int64(idP)
	ctx := c.Request().Context()
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		loggerFile.ErrorLogger.Println(err)
		return err
	}

	errUc := d.AUsecase.UploadBukti(ctx, id, file)

	if errUc != nil {
		loggerFile.ErrorLogger.Println(errUc)
		responseError3, _ := http_response.MapResponseBuktiDonor(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	contentLog.Payload = helper.StructToString(id)
	contentLog.Response = helper.StructToString("")
	loggerFile.InfoLogger.Println(contentLog)

	responseSuccess, _ := http_response.MapResponseBuktiDonor(0, "success", nil)
	return c.JSON(http.StatusOK, responseSuccess)
}

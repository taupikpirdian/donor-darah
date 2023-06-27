package http

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/delivery/http_response"
	"github.com/labstack/echo/v4"
)

func (d *DonorHandler) RescheduleDonor(c echo.Context) (err error) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)
	var schedule domain.DonorSchedulleDTO
	err = c.Bind(&schedule)
	if err != nil {
		logger.Print(err)
		fmt.Print(&buf)
		responseError, _ := http_response.MapResponseBuktiDonor(1, err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	idP, errAToi := strconv.Atoi(c.Param("scheduleId"))
	if errAToi != nil {
		logger.Print(errAToi)
		fmt.Print(&buf)
		responseErrorConv, _ := http_response.MapResponseBuktiDonor(1, err.Error(), nil)
		return c.JSON(getStatusCode(errAToi), responseErrorConv)
	}
	id := int64(idP)
	ctx := c.Request().Context()

	errUc := d.AUsecase.Reschedule(ctx, id, &schedule)
	if errUc != nil {
		logger.Print(errUc)
		fmt.Print(&buf)
		responseError3, _ := http_response.MapResponseBuktiDonor(1, errUc.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseBuktiDonor(0, "success", nil)
	return c.JSON(http.StatusOK, responseSuccess)
}

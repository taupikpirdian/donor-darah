package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/region/delivery/http_response"
	"github.com/labstack/echo/v4"
)

// Register will store the user by given request body
func (a *RegionHandler) VillageGet(c echo.Context) (err error) {
	subDistrictId := c.QueryParam("subDistrictId")

	ctx := c.Request().Context()
	data, errUc := a.AUsecase.GetVillage(ctx, subDistrictId)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseVillage(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseVillage(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}

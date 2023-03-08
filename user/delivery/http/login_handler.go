package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func isRequestValidLogin(m *domain.DtoRequestLogin) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *UserHandler) LoginController(c echo.Context) (err error) {
	var dto domain.DtoRequestLogin
	err = c.Bind(&dto)
	if err != nil {
		responseError, _ := http_response.MapResponse(1, err.Error())
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	var ok bool
	if ok, err = isRequestValidLogin(&dto); !ok {
		responseErrorRequest, _ := http_response.MapResponse(1, domain.ErrBadParamInput.Error())
		return c.JSON(http.StatusBadRequest, responseErrorRequest)
	}

	ctx := c.Request().Context()
	token, errUc := a.AUsecase.Login(ctx, &dto)
	if errUc != nil {
		responseErrorUsecase, _ := http_response.MapResponse(1, domain.ErrBadBody.Error())
		return c.JSON(getStatusCode(errUc), responseErrorUsecase)
	}

	responseSuccess, _ := http_response.MapResponseLogin(0, "success", token)
	return c.JSON(http.StatusOK, responseSuccess)
}

// func accessible(c echo.Context) error {
// 	return c.String(http.StatusOK, "Accessible")
// }

// func restricted(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(*jwtCustomClaims)
// 	name := claims.Name
// 	return c.String(http.StatusOK, "Welcome "+name+"!")
// }

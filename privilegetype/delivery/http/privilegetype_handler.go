package http

import (
	"context"
	"net/http"
	"github.com/user-management-with-go/models"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/user-management-with-go/privilegetype"
)

type ResponseError struct {
	Code  	int64  `json:"code"`
	Message string `json:"message"`
}

type HttpPrivilegeTypeHandler struct {
	PTUsecase privilegetype.Usecase
}

type GetResponse struct {
	Code  		int64       	`json:"code"`
	Message 	string      	`json:"message"`
	Data		[]*models.PrivilegeType  `json:"data"`
}

func NewPrivilegeTypeHttpHandler(e *echo.Echo, pts privilegetype.Usecase) {
	handler := &HttpPrivilegeTypeHandler{
		PTUsecase: pts,
	}
	e.GET("/api/v1/privilege_type", handler.Fetch)
}

func (a *HttpPrivilegeTypeHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listAr, err := a.PTUsecase.Fetch(ctx)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, GetResponse{Code:200, Message: "Success", Data: listAr})
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

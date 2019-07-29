package http

import (
	//"github.com/user-management-with-go/health"
	"github.com/user-management-with-go/role"
	"github.com/labstack/echo"
	"net/http"
)

type HttpHealthHandler struct {
}

type PostResponse struct {
	Code  		uint16       	`json:"code"`
	Message 	string      	`json:"message"`
}

func NewHealthHttpHandler(e *echo.Echo, rs role.Usecase) {
	handler := &HttpHealthHandler{}
	e.GET("/api/v1", handler.Get)
}

func (a *HttpHealthHandler) Get(c echo.Context) error {
	return c.JSON(http.StatusOK, PostResponse{Code:200, Message: "Success"})
}

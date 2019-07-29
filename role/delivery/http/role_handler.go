package http

import (
	"context"
	"net/http"
	"strconv"
	"github.com/user-management-with-go/models"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/user-management-with-go/role"
	"gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Code  	int64  `json:"code"`
	Message string `json:"message"`
}

type HttpRoleHandler struct {
	RUsecase role.Usecase
}

type DataResponse struct {	
	Page        int           	`json:"page"`
	TotalData   int           	`json:"total_data"`
	Data		[]*models.Role  `json:"data"`
}

type GetResponse struct {
	Code  		int64       	`json:"code"`
	Message 	string      	`json:"message"`
	Data		interface{}  	`json:"data"`
}

type PostResponse struct {
	Code  		int64       	`json:"code"`
	Message 	string      	`json:"message"`
}

func NewRoleHttpHandler(e *echo.Echo, rs role.Usecase) {
	handler := &HttpRoleHandler{
		RUsecase: rs,
	}
	e.GET("/api/v1/role", handler.Fetch)
	e.POST("/api/v1/role", handler.Store)
	e.PUT("/api/v1/role/:id/update", handler.Update)
}

func (a *HttpRoleHandler) Fetch(c echo.Context) error {
	page := c.QueryParam("page")
	pages, _ := strconv.Atoi(page)

	limit := c.QueryParam("limit")
	limits, _ := strconv.Atoi(limit)

	search := c.QueryParam("search")
	searchs := "%"+search+"%"

	status := c.QueryParam("status")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listAr, total, err := a.RUsecase.Fetch(ctx, int64(pages), int64(limits), searchs, status)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}

	resp := DataResponse{Page:pages, TotalData:total, Data: listAr}

	return c.JSON(http.StatusOK, GetResponse{Code:200, Message: "Success", Data: resp})
}

func (a *HttpRoleHandler) Store(c echo.Context) error {
	var role models.Role
	err := c.Bind(&role)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&role); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}

	err = a.RUsecase.Store(ctx, &role)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, PostResponse{Code:200, Message: "Role has been created"})
}

func (a *HttpRoleHandler) Update(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var role models.Role
	err = c.Bind(&role)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&role); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx = c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.RUsecase.Update(ctx, &role, id)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, PostResponse{Code:200, Message: "role has been updated"})
}


func isRequestValid(m *models.Role) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
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

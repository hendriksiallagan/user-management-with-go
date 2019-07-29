package http

import (
	"context"
	"net/http"
	"strconv"
	"github.com/user-management-with-go/models"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/user-management-with-go/page"
	"gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Code  	int64  `json:"code"`
	Message string `json:"message"`
}

type HttpPageHandler struct {
	PUsecase page.Usecase
}

type DataResponse struct {
	Page        int           `json:"page"`
	TotalData   int           `json:"total_data"`
	Data		[]*models.Page  `json:"data"`
}

type GetResponse struct {
	Code  		int64       	`json:"code"`
	Message 	string      	`json:"message"`
	Data		interface{} 	`json:"data"`
}

type PostResponse struct {
	Code  		int64       	`json:"code"`
	Message 	string      	`json:"message"`
}

func NewPageHttpHandler(e *echo.Echo, ps page.Usecase) {
	handler := &HttpPageHandler{
		PUsecase: ps,
	}
	e.GET("/api/v1/page", handler.Fetch)
	e.POST("/api/v1/page", handler.Store)
	e.PUT("/api/v1/page/:id/update", handler.Update)
}

func (a *HttpPageHandler) Fetch(c echo.Context) error {
	page := c.QueryParam("page")
	pages, _ := strconv.Atoi(page)

	limit := c.QueryParam("limit")
	limits, _ := strconv.Atoi(limit)

	search := c.QueryParam("search")
	searchs := "%"+search+"%"

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listAr, total ,err := a.PUsecase.Fetch(ctx, int64(pages), int64(limits), searchs)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}

	resp := DataResponse{Page:pages, TotalData:total, Data: listAr}

	return c.JSON(http.StatusOK, GetResponse{Code:200, Message: "Success", Data: resp })
}

func (a *HttpPageHandler) Store(c echo.Context) error {
	var page models.Page
	err := c.Bind(&page)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&page); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	generateCode, err := a.PUsecase.Generate(ctx)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}

	err = a.PUsecase.Store(ctx, &page, generateCode)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, PostResponse{Code:200, Message: "page has been created"})
}

func (a *HttpPageHandler) Update(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var page models.Page
	err = c.Bind(&page)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&page); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx = c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.PUsecase.Update(ctx, &page, id)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, PostResponse{Code:200, Message: "page has been updated"})
}


func isRequestValid(m *models.Page) (bool, error) {

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

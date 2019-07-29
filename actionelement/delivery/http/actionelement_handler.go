package http

import (
	"context"
	"net/http"
	"strconv"
	"github.com/user-management-with-go/models"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/user-management-with-go/actionelement"
	"gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Code  	int64  `json:"code"`
	Message string `json:"message"`
}

type HttpActionelementHandler struct {
	AUsecase actionelement.Usecase
}

type DataResponse struct {
	Page        int           `json:"page"`
	TotalData   int           `json:"total_data"`
	Data		[]*models.Actionelement  `json:"data"`
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

func NewActionelementHttpHandler(e *echo.Echo, as actionelement.Usecase) {
	handler := &HttpActionelementHandler{
		AUsecase: as,
	}
	e.GET("/api/v1/actionelement", handler.Fetch)
	e.POST("/api/v1/actionelement", handler.Store)
	e.PUT("/api/v1/actionelement/:id/update", handler.Update)
}

func (a *HttpActionelementHandler) Fetch(c echo.Context) error {
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
	listAr, total ,err := a.AUsecase.Fetch(ctx, int64(pages), int64(limits), searchs)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}

	resp := DataResponse{Page:pages, TotalData:total, Data: listAr}

	return c.JSON(http.StatusOK, GetResponse{Code:200, Message: "Success", Data: resp })
}

func (a *HttpActionelementHandler) Store(c echo.Context) error {
	var actionelement models.Actionelement
	err := c.Bind(&actionelement)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&actionelement); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	generateCode, err := a.AUsecase.Generate(ctx)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}

	err = a.AUsecase.Store(ctx, &actionelement, generateCode)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, PostResponse{Code:200, Message: "Action Element has been created"})
}

func (a *HttpActionelementHandler) Update(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var actionelement models.Actionelement
	err = c.Bind(&actionelement)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&actionelement); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx = c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.AUsecase.Update(ctx, &actionelement, id)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, PostResponse{Code:200, Message: "Action Element has been updated"})
}


func isRequestValid(m *models.Actionelement) (bool, error) {

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

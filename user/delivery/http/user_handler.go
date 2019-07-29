package http

import (
	"context"
	"github.com/user-management-with-go/models"
	"github.com/user-management-with-go/user"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type HttpUserHandler struct {
	UUsecase user.Usecase
}


func NewUserHttpHandler(e *echo.Echo, us user.Usecase) {
	handler := &HttpUserHandler{
		UUsecase: us,
	}
	e.GET("/api/v1/user", handler.Fetch)
	e.GET("/api/v1/user/:id/password", handler.FetchPasswordByID)
	e.GET("/api/v1/user/:id/pin", handler.FetchPinByID)
	e.GET("/api/v1/user/:id/role", handler.FetchRoleByID)
	e.GET("/api/v1/user/:id/status-info", handler.FetchStatusInfoByID)
	e.GET("/api/v1/user/:id/otp", handler.FetchOtpByID)
	e.POST("/api/v1/user", handler.Store)
	e.POST("/api/v1/user/:id/role", handler.StoreRole)
	e.GET("/api/v1/user/:id", handler.FetchByID)
	e.PUT("/api/v1/user/:id/update", handler.Update)
	e.PUT("/api/v1/user/:id/role/update", handler.UpdateRole)
	e.PUT("/api/v1/user/:id/password/reset", handler.UpdatePassword)
}

func (a *HttpUserHandler) Fetch(c echo.Context) error {
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
	listAr, total, err := a.UUsecase.Fetch(ctx, uint64(pages), uint64(limits), searchs)

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}

	resp := models.DataResponse{Page:pages, TotalData:int(total), Data: listAr}

	return c.JSON(http.StatusOK, models.GetResponse{Code:200, Message: "Success", Data: resp})
}

func (a *HttpUserHandler) FetchPasswordByID(c echo.Context) error {
	page := c.QueryParam("page")
	pages, _ := strconv.Atoi(page)

	limit := c.QueryParam("limit")
	limits, _ := strconv.Atoi(limit)

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listPass, total, err := a.UUsecase.FetchPasswordByID(ctx, uint64(pages), uint64(limits), uint64(id))

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}

	resp := models.DataResponse{Page:pages, TotalData:int(total), Data: listPass}

	return c.JSON(http.StatusOK, models.GetResponse{Code:200, Message: "Success", Data: resp})
}

func (a *HttpUserHandler) FetchPinByID(c echo.Context) error {
	page := c.QueryParam("page")
	pages, _ := strconv.Atoi(page)

	limit := c.QueryParam("limit")
	limits, _ := strconv.Atoi(limit)

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listPass, total, err := a.UUsecase.FetchPinByID(ctx, uint64(pages), uint64(limits), uint64(id))

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}

	resp := models.DataResponse{Page:pages, TotalData:int(total), Data: listPass}

	return c.JSON(http.StatusOK, models.GetResponse{Code:200, Message: "Success", Data: resp})
}

func (a *HttpUserHandler) FetchRoleByID(c echo.Context) error {
	page := c.QueryParam("page")
	pages, _ := strconv.Atoi(page)

	limit := c.QueryParam("limit")
	limits, _ := strconv.Atoi(limit)

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listPass, total, err := a.UUsecase.FetchRoleByID(ctx, uint64(pages), uint64(limits), uint64(id))

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}

	resp := models.DataResponse{Page:pages, TotalData:int(total), Data: listPass}

	return c.JSON(http.StatusOK, models.GetResponse{Code:200, Message: "Success", Data: resp})
}

func (a *HttpUserHandler) FetchStatusInfoByID(c echo.Context) error {
	page := c.QueryParam("page")
	pages, _ := strconv.Atoi(page)

	limit := c.QueryParam("limit")
	limits, _ := strconv.Atoi(limit)

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listPass, total, err := a.UUsecase.FetchStatusInfoByID(ctx, uint64(pages), uint64(limits), uint64(id))

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}

	resp := models.DataResponse{Page:pages, TotalData:int(total), Data: listPass}

	return c.JSON(http.StatusOK, models.GetResponse{Code:200, Message: "Success", Data: resp})
}

func (a *HttpUserHandler) FetchOtpByID(c echo.Context) error {
	page := c.QueryParam("page")
	pages, _ := strconv.Atoi(page)

	limit := c.QueryParam("limit")
	limits, _ := strconv.Atoi(limit)

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listPass, total, err := a.UUsecase.FetchOtpByID(ctx, uint64(pages), uint64(limits), uint64(id))

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}

	resp := models.DataResponse{Page:pages, TotalData:int(total), Data: listPass}

	return c.JSON(http.StatusOK, models.GetResponse{Code:200, Message: "Success", Data: resp})
}

func (a *HttpUserHandler) FetchByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	listArById, err := a.UUsecase.FetchByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.GetDetailResponse{Code:200, Message: "Success", Data: listArById})
}

func (a *HttpUserHandler) Store(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	generateCode, err := a.UUsecase.Generate(ctx)

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}

	err = a.UUsecase.Store(ctx, &user, generateCode)

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.Response{Code:200, Message: "User has been created"})
}

func (a *HttpUserHandler) StoreRole(c echo.Context) error {
	var userRole models.UserRole
	err := c.Bind(&userRole)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&userRole); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	generateCode, err := a.UUsecase.GenerateRole(ctx)

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}

	err = a.UUsecase.StoreRole(ctx, &userRole, generateCode)

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.Response{Code:200, Message: "User Role has been created"})
}

func (a *HttpUserHandler) Update(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var user models.User
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx = c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.UUsecase.Update(ctx, &user, id)

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.Response{Code:200, Message: "User has been updated"})
}

func (a *HttpUserHandler) UpdateRole(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := uint64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var userRole models.UserRole
	err = c.Bind(&userRole)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&userRole); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx = c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.UUsecase.UpdateRole(ctx, &userRole, id)

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.Response{Code:200, Message: "User Role has been updated"})
}

func (a *HttpUserHandler) UpdatePassword(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := uint64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var userReset models.ResetPassword
	err = c.Bind(&userReset)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&userReset); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx = c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.UUsecase.UpdatePassword(ctx, &userReset, id)

	if err != nil {
		return c.JSON(getStatusCode(err), models.Response{Code:201, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, models.Response{Code:200, Message: "User Password has been updated"})
}


func isRequestValid(m interface{}) (bool, error) {

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

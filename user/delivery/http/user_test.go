package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	//"strconv"
	"strings"
	"testing"
	//"time"

	//"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	userHttp "github.com/user-management-with-go/user/delivery/http"
	"github.com/user-management-with-go/user/mocks"
	"github.com/user-management-with-go/models"
)

func TestStore(t *testing.T) {
	mockUser := models.User{
		MuCode:   "B005",
		MuName:   "Mahmud",
		MuEmail: "mahmud@gmail.com",
		MuDescription: "Product Manager",
	}

	tempmockUser := mockUser
	tempmockUser.MuID = 8
	mockUCase := new(mocks.Usecase)

	j, err := json.Marshal(tempmockUser)
	assert.NoError(t, err)

	mockUCase.On("Store", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/user", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	handler := userHttp.HttpUserHandler{
		UUsecase: mockUCase,
	}
	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}
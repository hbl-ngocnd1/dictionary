package controller

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hbl-ngocnd1/dictionary/usecase"
	"github.com/hbl-ngocnd1/dictionary/usecase/mock_usecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitTest() {
	if os.Getenv("CI") == "true" {
		return
	}
	err := godotenv.Load("../.env_test")
	if err != nil {
		panic(err)
	}
}

func TestNewDictHandler(t *testing.T) {
	patterns := []struct {
		description   string
		urlParam      string
		newMockDictUC func(ctrl *gomock.Controller) usecase.DictUseCase
		statusCode    int
	}{
		{
			description: "400: error",
			urlParam:    "level=n6",
			statusCode:  http.StatusBadRequest,
		},
		{
			description: "500: error",
			urlParam:    "",
			newMockDictUC: func(ctrl *gomock.Controller) usecase.DictUseCase {
				mock := mock_usecase.NewMockDictUseCase(ctrl)
				mock.EXPECT().GetDict(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("unexpected"))
				return mock
			},
			statusCode: http.StatusInternalServerError,
		},
	}
	for i, p := range patterns {
		t.Run(fmt.Sprintf("%d:%s", i, p.description), func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", "/api/dictionary", p.urlParam), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			var mockDictUC usecase.DictUseCase
			if p.newMockDictUC != nil {
				mockDictUC = p.newMockDictUC(ctrl)
			}
			h := &dictHandler{dictUseCase: mockDictUC}
			// Assertions
			if assert.NoError(t, h.ApiDict(c)) {
				assert.Equal(t, p.statusCode, rec.Code)
			}
		})
	}
}

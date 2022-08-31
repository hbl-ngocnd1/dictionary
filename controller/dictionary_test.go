package controller

import (
	"errors"
	"fmt"
	"github.com/hbl-ngocnd1/dictionary/models"
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
			description: "200: success",
			urlParam:    "",
			statusCode:  http.StatusOK,
			newMockDictUC: func(ctrl *gomock.Controller) usecase.DictUseCase {
				mock := mock_usecase.NewMockDictUseCase(ctrl)
				mock.EXPECT().GetDict(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]models.Word{{}, {}, {}}, nil)
				return mock
			},
		},
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
		{
			description: "400: GetDict return InvalidErr",
			urlParam:    "",
			statusCode:  http.StatusBadRequest,
			newMockDictUC: func(ctrl *gomock.Controller) usecase.DictUseCase {
				mock := mock_usecase.NewMockDictUseCase(ctrl)
				mock.EXPECT().GetDict(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]models.Word{{}, {}, {}}, usecase.InvalidErr)
				return mock
			},
		},
		{
			description: "403: GetDict return PermissionDeniedErr",
			urlParam:    "",
			statusCode:  http.StatusForbidden,
			newMockDictUC: func(ctrl *gomock.Controller) usecase.DictUseCase {
				mock := mock_usecase.NewMockDictUseCase(ctrl)
				mock.EXPECT().GetDict(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]models.Word{{}, {}, {}}, usecase.PermissionDeniedErr)
				return mock
			},
		},
		{
			description: "500: GetDict return StatusInternalServerError",
			urlParam:    "",
			statusCode:  http.StatusInternalServerError,
			newMockDictUC: func(ctrl *gomock.Controller) usecase.DictUseCase {
				mock := mock_usecase.NewMockDictUseCase(ctrl)
				mock.EXPECT().GetDict(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]models.Word{{}, {}, {}}, errors.New("another Error"))
				return mock
			},
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

func TestNewGetDetailHandler(t *testing.T) {
	patterns := []struct {
		description   string
		index         map[string]string
		urlParam      string
		newMockDictUC func(ctrl *gomock.Controller) usecase.DictUseCase
		statusCode    int
	}{

		{
			description: "200: successful request",
			index:       map[string]string{"index": "1"},
			urlParam:    "level=n1",
			newMockDictUC: func(ctrl *gomock.Controller) usecase.DictUseCase {
				mock := mock_usecase.NewMockDictUseCase(ctrl)
				mock.EXPECT().GetDetail(gomock.Any(), gomock.Any(), gomock.Any()).Return(makeTestString("success"), nil)
				return mock
			},
			statusCode: http.StatusOK,
		},
		{
			description: "200: successful with nil data",
			index:       map[string]string{"index": "1"},
			urlParam:    "level=n1",
			newMockDictUC: func(ctrl *gomock.Controller) usecase.DictUseCase {
				mock := mock_usecase.NewMockDictUseCase(ctrl)
				mock.EXPECT().GetDetail(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
				return mock
			},
			statusCode: http.StatusOK,
		},
		{
			description: "400: bad request with index",
			urlParam:    "level=n6",
			statusCode:  http.StatusBadRequest,
		},
		{
			description: "400: bad request with level",
			index:       map[string]string{"index": "1"},
			urlParam:    "level=n6",
			statusCode:  http.StatusBadRequest,
		},
		{
			description: "500: error",
			index:       map[string]string{"index": "1"},
			urlParam:    "level=n1",
			newMockDictUC: func(ctrl *gomock.Controller) usecase.DictUseCase {
				mock := mock_usecase.NewMockDictUseCase(ctrl)
				mock.EXPECT().GetDetail(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("unexpected"))
				return mock
			},
			statusCode: http.StatusInternalServerError,
		},
		{
			description: "500: InvalidErr",
			index:       map[string]string{"index": "1"},
			urlParam:    "level=n1",
			newMockDictUC: func(ctrl *gomock.Controller) usecase.DictUseCase {
				mock := mock_usecase.NewMockDictUseCase(ctrl)
				mock.EXPECT().GetDetail(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, usecase.InvalidErr)
				return mock
			},
			statusCode: http.StatusBadRequest,
		},
	}
	for i, p := range patterns {
		t.Run(fmt.Sprintf("%d:%s", i, p.description), func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s?%s", "/api/dictionary/1", p.urlParam), nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			if len(p.index) != 0 {
				for k, v := range p.index {
					c.SetParamNames(k)
					c.SetParamValues(v)
				}
			}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			var mockDictUC usecase.DictUseCase
			if p.newMockDictUC != nil {
				mockDictUC = p.newMockDictUC(ctrl)
			}
			h := &dictHandler{dictUseCase: mockDictUC}
			// Assertions
			if assert.NoError(t, h.ApiGetDetail(c)) {
				assert.Equal(t, p.statusCode, rec.Code)
			}
		})
	}
}

func makeTestString(s string) *string {
	return &s
}

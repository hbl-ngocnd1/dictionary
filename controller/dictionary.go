package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/hbl-ngocnd1/dictionary/models"

	"github.com/hbl-ngocnd1/dictionary/usecase"
	"github.com/labstack/echo/v4"
)

type dictHandler struct {
	dictUseCase usecase.DictUseCase
}

func NewDictHandler() *dictHandler {
	return &dictHandler{
		dictUseCase: usecase.NewDictUseCase(),
	}
}

func (f *dictHandler) Dict(c echo.Context) error {
	return c.Render(http.StatusOK, "dictionary.html", map[string]interface{}{"router": "dictionary"})
}

func (f *dictHandler) ApiDict(c echo.Context) error {
	return getDataJapanese(f, c, models.MakeWord)
}

func (f *dictHandler) ApiGetDetail(c echo.Context) error {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	level := c.QueryParam("level")
	switch level {
	case "n1", "n2", "n3", "n4", "n5":
	default:
		return c.NoContent(http.StatusBadRequest)
	}
	ctx := context.Background()
	data, err := f.dictUseCase.GetDetail(ctx, level, index)
	switch err {
	case nil:
	case usecase.InvalidErr:
		return c.NoContent(http.StatusBadRequest)
	default:
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if data == nil {
		return c.String(http.StatusOK, "")
	}
	return c.String(http.StatusOK, *data)
}

func (f *dictHandler) ITJapanWonderWord(c echo.Context) error {
	return c.Render(http.StatusOK, "wonder-word.html", map[string]interface{}{"router": "wonder-word"})
}

func (f *dictHandler) ApiITJapanWonderWord(c echo.Context) error {
	ctx := context.Background()
	data, err := f.dictUseCase.GetITJapanWonderWork(ctx, models.MakeWonderWork)
	switch err {
	case nil:
	case usecase.InvalidErr:
		return c.NoContent(http.StatusBadRequest)
	default:
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if data == nil {

		return c.String(http.StatusOK, "")
	}
	return c.JSON(http.StatusOK, data)
}

func getDataJapanese(f *dictHandler, c echo.Context, makeData models.MakeData) error {
	notCache := c.QueryParam("not_cache")
	level := c.QueryParam("level")
	start, err := strconv.Atoi(c.QueryParam("start"))
	if err != nil {
		start = 0
	}
	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil {
		pageSize = 20
	}
	if level == "" {
		level = "n1"
	}
	switch level {
	case "n1", "n2", "n3", "n4", "n5":
	default:
		return c.NoContent(http.StatusBadRequest)
	}
	pwd := c.QueryParam("password")
	ctx := context.Background()
	data, err := f.dictUseCase.GetDict(ctx, start, pageSize, notCache, level, pwd, makeData)
	switch err {
	case nil:
	case usecase.InvalidErr:
		return c.NoContent(http.StatusBadRequest)
	case usecase.PermissionDeniedErr:
		return c.NoContent(http.StatusForbidden)
	default:
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

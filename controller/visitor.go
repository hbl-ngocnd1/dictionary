package controller

import (
	"context"
	"net/http"

	"github.com/hbl-ngocnd1/dictionary/models"
	"github.com/hbl-ngocnd1/dictionary/services"
	"github.com/labstack/echo/v4"
	"github.com/timjacobi/go-couchdb"
)

type VisitorHandler struct {
	visitorService services.VisitorService
}

func NewVisitorHandler(db *couchdb.DB) *VisitorHandler {
	return &VisitorHandler{
		visitorService: services.NewVisitor(db),
	}
}
func (h *VisitorHandler) CreateVisitor(c echo.Context) error {
	ctx := context.Background()
	var visitor models.Visitor
	if c.Bind(&visitor) == nil {
		err := h.visitorService.CreateVisitor(ctx, visitor)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.String(http.StatusOK, "Hello "+visitor.Name)
	}
	return nil
}
func (h *VisitorHandler) ListVisitor(c echo.Context) error {
	ctx := context.Background()
	result, err := h.visitorService.GetListVisitor(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "unable to fetch docs"})
	}
	return c.JSON(http.StatusOK, result.Rows)
}

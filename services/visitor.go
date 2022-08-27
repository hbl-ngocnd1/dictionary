package services

import (
	"context"

	"github.com/hbl-ngocnd1/dictionary/models"
	"github.com/timjacobi/go-couchdb"
)

type alldocsResult struct {
	TotalRows int `json:"total_rows"`
	Offset    int
	Rows      []map[string]interface{}
}
type visitorService struct {
	db *couchdb.DB
}
type VisitorService interface {
	CreateVisitor(ctx context.Context, vm models.Visitor) error
	GetListVisitor(ctx context.Context) (*alldocsResult, error)
}

func NewVisitor(db *couchdb.DB) *visitorService {
	return &visitorService{
		db: db,
	}
}
func (v *visitorService) CreateVisitor(ctx context.Context, vm models.Visitor) error {
	_, _, err := v.db.Post(vm)
	if err != nil {
		return err
	}
	return nil
}
func (v *visitorService) GetListVisitor(ctx context.Context) (*alldocsResult, error) {
	var result alldocsResult
	err := v.db.AllDocs(&result, couchdb.Options{"include_docs": true})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

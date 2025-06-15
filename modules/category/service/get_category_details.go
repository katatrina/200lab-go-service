package categoryservice

import (
	"context"
	categorymodel "go12-service/modules/category/model"

	"github.com/google/uuid"
)

type GetDetailQuery struct {
	Id uuid.UUID
}

type IGetDetailRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
}

type GetDetailQueryHandler struct {
	repo IGetDetailRepo
}

func NewGetDetailQueryHandler(repo IGetDetailRepo) *GetDetailQueryHandler {
	return &GetDetailQueryHandler{repo: repo}
}

func (hdl *GetDetailQueryHandler) Execute(ctx context.Context, query *GetDetailQuery) (*categorymodel.Category, error) {
	category, err := hdl.repo.FindByID(ctx, query.Id)

	if err != nil {
		return nil, err
	}

	if category.Status == categorymodel.StatusDeleted {
		return nil, categorymodel.ErrCategoryNotFound
	}

	return category, nil
}

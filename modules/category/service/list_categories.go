package categoryservice

import (
	"context"
	categorymodel "go12-service/modules/category/model"
	sharedmodel "go12-service/shared/model"
)

type ListQuery struct {
	sharedmodel.PagingDTO
	categorymodel.FilterCategoryDTO
}

type IListRepo interface {
	List(ctx context.Context, pagingDTO *sharedmodel.PagingDTO, filterDTO *categorymodel.FilterCategoryDTO) ([]categorymodel.Category, error)
}

type ListQueryHandler struct {
	repo IListRepo
}

func NewListQueryHandler(repo IListRepo) *ListQueryHandler {
	return &ListQueryHandler{repo: repo}
}

func (hdl *ListQueryHandler) Execute(ctx context.Context, query *ListQuery) ([]categorymodel.Category, error) {
	categories, err := hdl.repo.List(ctx, &query.PagingDTO, &query.FilterCategoryDTO)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

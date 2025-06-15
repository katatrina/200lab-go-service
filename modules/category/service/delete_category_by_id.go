package categoryservice

import (
	"context"
	categorymodel "go12-service/modules/category/model"

	"github.com/google/uuid"
)

type DeleteByIdCommand struct {
	Id uuid.UUID
}

type IDeleteByIdRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
	Delete(ctx context.Context, id uuid.UUID, isHard bool) error
}

type DeleteByIdCommandHandler struct {
	repo IDeleteByIdRepo
}

func NewDeleteByIdCommandHandler(repo IDeleteByIdRepo) *DeleteByIdCommandHandler {
	return &DeleteByIdCommandHandler{repo: repo}
}

func (hdl *DeleteByIdCommandHandler) Execute(ctx context.Context, cmd *DeleteByIdCommand) error {
	category, err := hdl.repo.FindByID(ctx, cmd.Id)

	if err != nil {
		return err
	}

	if category.Status == categorymodel.StatusDeleted {
		return categorymodel.ErrCategoryIsDeleted
	}

	if err := hdl.repo.Delete(ctx, cmd.Id, false); err != nil {
		return err
	}

	return nil
}

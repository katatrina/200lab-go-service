package categoryservice

import (
	"context"
	categorymodel "go12-service/modules/category/model"

	"github.com/google/uuid"
)

type UpdateCategoryCommand struct {
	Id uuid.UUID `json:"-"`
	categorymodel.CategoryUpdateDTO
}

type IUpdateByIdRepo interface {
	FindByID(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
	Update(ctx context.Context, id uuid.UUID, dto *categorymodel.CategoryUpdateDTO) error
}

type UpdateByIdCommandHandler struct {
	repo IUpdateByIdRepo
}

func NewUpdateByIdCommandHandler(repo IUpdateByIdRepo) *UpdateByIdCommandHandler {
	return &UpdateByIdCommandHandler{repo: repo}
}

func (hdl *UpdateByIdCommandHandler) Execute(ctx context.Context, cmd *UpdateCategoryCommand) error {
	if err := cmd.Validate(); err != nil {
		return err
	}

	category, err := hdl.repo.FindByID(ctx, cmd.Id)

	if err != nil {
		return err
	}

	if category.Status == categorymodel.StatusDeleted {
		return categorymodel.ErrCategoryIsDeleted
	}

	if err := hdl.repo.Update(ctx, cmd.Id, &cmd.CategoryUpdateDTO); err != nil {
		return err
	}

	return nil
}

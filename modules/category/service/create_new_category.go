package categoryservice

import (
	"context"
	categorymodel "go12-service/modules/category/model"

	"github.com/google/uuid"
)

type CreateCommand struct {
	Dto categorymodel.Category
}

type ICreateRepo interface {
	Insert(ctx context.Context, data *categorymodel.Category) error
}

type CreateCommandHandler struct {
	catRepo ICreateRepo
}

func NewCreateCommandHandler(catRepo ICreateRepo) *CreateCommandHandler {
	return &CreateCommandHandler{catRepo: catRepo}
}

func (hdl *CreateCommandHandler) Execute(ctx context.Context, cmd *CreateCommand) (*uuid.UUID, error) {
	if err := cmd.Dto.Validate(); err != nil {
		return nil, err
	}

	cmd.Dto.Id, _ = uuid.NewV7()

	if err := hdl.catRepo.Insert(ctx, &cmd.Dto); err != nil {
		return nil, err
	}

	return &cmd.Dto.Id, nil
}

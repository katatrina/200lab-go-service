package categorygormmysql

import (
	"context"
	categorymodel "go12-service/modules/category/model"

	"github.com/google/uuid"
)

func (r *CategoryRepository) Delete(ctx context.Context, id uuid.UUID, isHard bool) error {
	db := r.db

	if isHard {
		if err := db.Table(categorymodel.Category{}.TableName()).
			Where("id = ?", id).
			Delete(nil).
			Error; err != nil {
			return err
		}
	}

	if err := db.Table(categorymodel.Category{}.TableName()).
		Where("id = ?", id).
		Update("status", categorymodel.StatusDeleted).
		Error; err != nil {
		return err
	}

	return nil
}

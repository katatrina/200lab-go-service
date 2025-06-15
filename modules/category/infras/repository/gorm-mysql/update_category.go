package categorygormmysql

import (
	"context"
	categorymodel "go12-service/modules/category/model"

	"github.com/google/uuid"
)

func (r *CategoryRepository) Update(ctx context.Context, id uuid.UUID, dto *categorymodel.CategoryUpdateDTO) error {
	db := r.db.Begin()

	if err := db.Table(dto.TableName()).Where("id = ?", id).Updates(dto).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}

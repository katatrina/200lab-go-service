package categorygormmysql

import (
	"context"
	categorymodel "go12-service/modules/category/model"
	sharedmodel "go12-service/shared/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *CategoryRepository) FindByID(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error) {
	var category categorymodel.Category

	if err := repo.db.Where("id = ?", id).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, sharedmodel.ErrRecordNotFound
		}
		return nil, err
	}
	return &category, nil
}

package categorygormmysql

import (
	"context"
	categorymodel "go12-service/modules/category/model"
	sharedmodel "go12-service/shared/model"
)

func (repo *CategoryRepository) List(ctx context.Context, pagingDTO *sharedmodel.PagingDTO, filterDTO *categorymodel.FilterCategoryDTO) ([]categorymodel.Category, error) {

	var data []categorymodel.Category

	ldb := repo.db.Where("status in (?)", []string{categorymodel.StatusActive})

	if err := ldb.Table(categorymodel.Category{}.TableName()).Count(&pagingDTO.Total).Error; err != nil {
		return nil, err
	}

	if err := ldb.Order("id desc").
		Limit(pagingDTO.Limit).
		Offset((pagingDTO.Page - 1) * pagingDTO.Limit).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

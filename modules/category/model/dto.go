package categorymodel

import "strings"

// DTO = Data Transfer Object
type CategoryUpdateDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

func (c *CategoryUpdateDTO) Validate() error {
	if st := c.Status; st != nil && *st != StatusActive && *st != StatusInactive && *st != StatusDeleted {
		return ErrCategoryStatusInvalid
	}

	if str := c.Name; str != nil {
		*c.Name = strings.TrimSpace(*str)

		if len(*c.Name) == 0 {
			return ErrNameIsEmpty
		}
	}

	return nil
}

func (CategoryUpdateDTO) TableName() string {
	return Category{}.TableName()
}

type FilterCategoryDTO struct {
	Status string `json:"status"`
}

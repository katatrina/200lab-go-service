package categorymodel

import "errors"

var (
	ErrNameIsRequired        = errors.New("name is required")
	ErrNameIsEmpty           = errors.New("name is empty")
	ErrCategoryNotFound      = errors.New("category not found")
	ErrCategoryStatusInvalid = errors.New("status must be in (active, inactive, deleted)")
	ErrCategoryIsDeleted     = errors.New("category is deleted")
)

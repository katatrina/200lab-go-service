package categorymodel

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id          uuid.UUID  `json:"id" gorm:"column:id;"`
	Name        string     `json:"name" gorm:"column:name;"`
	Description string     `json:"description" gorm:"column:description;"`
	Status      string     `json:"status" gorm:"column:status;"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updated_at;"`
}

func (Category) TableName() string {
	return "categories"
}

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusDeleted  = "deleted"
)

func (c *Category) Validate() error {
	c.Name = strings.TrimSpace(c.Name)

	if c.Name == "" {
		return ErrNameIsRequired
	}

	if c.Status != StatusActive && c.Status != StatusInactive && c.Status != StatusDeleted {
		return ErrCategoryStatusInvalid
	}

	return nil
}

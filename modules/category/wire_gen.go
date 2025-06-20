// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package categorymodule

import (
	"go12-service/modules/category/infras/controller/http-gin"
	"go12-service/modules/category/infras/repository/gorm-mysql"
	"go12-service/modules/category/service"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeCategoryController(db *gorm.DB) *categoryhttpgin.CategoryHTTPController {
	categoryRepository := categorygormmysql.NewCategoryRepository(db)
	createCommandHandler := categoryservice.NewCreateCommandHandler(categoryRepository)
	getDetailQueryHandler := categoryservice.NewGetDetailQueryHandler(categoryRepository)
	updateByIdCommandHandler := categoryservice.NewUpdateByIdCommandHandler(categoryRepository)
	deleteByIdCommandHandler := categoryservice.NewDeleteByIdCommandHandler(categoryRepository)
	listQueryHandler := categoryservice.NewListQueryHandler(categoryRepository)
	categoryHTTPController := categoryhttpgin.NewCategoryHTTPController(createCommandHandler, getDetailQueryHandler, updateByIdCommandHandler, deleteByIdCommandHandler, listQueryHandler)
	return categoryHTTPController
}

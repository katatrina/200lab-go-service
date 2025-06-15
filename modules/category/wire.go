//go:build wireinject

package categorymodule

import (
	categoryhttpgin "go12-service/modules/category/infras/controller/http-gin"
	categorygormmysql "go12-service/modules/category/infras/repository/gorm-mysql"
	categoryservice "go12-service/modules/category/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeCategoryController(db *gorm.DB) *categoryhttpgin.CategoryHTTPController {
	wire.Build(
		categoryhttpgin.NewCategoryHTTPController,
		categoryservice.NewGetDetailQueryHandler,
		categoryservice.NewCreateCommandHandler,
		categoryservice.NewDeleteByIdCommandHandler,
		categoryservice.NewUpdateByIdCommandHandler,
		categoryservice.NewListQueryHandler,
		categorygormmysql.NewCategoryRepository,

		wire.Bind(new(categoryhttpgin.IDetailQueryHandler), new(*categoryservice.GetDetailQueryHandler)),
		wire.Bind(new(categoryhttpgin.ICreateCommandHandler), new(*categoryservice.CreateCommandHandler)),
		wire.Bind(new(categoryhttpgin.IDeleteByIdCommandHandler), new(*categoryservice.DeleteByIdCommandHandler)),
		wire.Bind(new(categoryhttpgin.IUpdateByIdCommandHandler), new(*categoryservice.UpdateByIdCommandHandler)),
		wire.Bind(new(categoryhttpgin.IListQueryHandler), new(*categoryservice.ListQueryHandler)),

		wire.Bind(new(categoryservice.ICreateRepo), new(*categorygormmysql.CategoryRepository)),
		wire.Bind(new(categoryservice.IDeleteByIdRepo), new(*categorygormmysql.CategoryRepository)),
		wire.Bind(new(categoryservice.IUpdateByIdRepo), new(*categorygormmysql.CategoryRepository)),
		wire.Bind(new(categoryservice.IListRepo), new(*categorygormmysql.CategoryRepository)),
		wire.Bind(new(categoryservice.IGetDetailRepo), new(*categorygormmysql.CategoryRepository)),
	)

	return &categoryhttpgin.CategoryHTTPController{}
}

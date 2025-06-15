package categoryhttpgin

import (
	"context"
	categorymodel "go12-service/modules/category/model"
	categoryservice "go12-service/modules/category/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// type ICategoryService interface {
// 	CreateNewCategory(ctx context.Context, data *categorymodel.Category) error
// 	GetCategoryDetails(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
// 	ListCategories(ctx context.Context, dto *categoryservice.ListCategoriesDTO) ([]categorymodel.Category, error)
// 	UpdateCategoryById(ctx context.Context, cmd *categoryservice.UpdateCategoryCommand) error
// 	DeleteCategoryById(ctx context.Context, id uuid.UUID) error
// }

type ICreateCommandHandler interface {
	Execute(ctx context.Context, cmd *categoryservice.CreateCommand) (*uuid.UUID, error)
}

type IDetailQueryHandler interface {
	Execute(ctx context.Context, query *categoryservice.GetDetailQuery) (*categorymodel.Category, error)
}

type IUpdateByIdCommandHandler interface {
	Execute(ctx context.Context, cmd *categoryservice.UpdateCategoryCommand) error
}

type IDeleteByIdCommandHandler interface {
	Execute(ctx context.Context, cmd *categoryservice.DeleteByIdCommand) error
}

type IListQueryHandler interface {
	Execute(ctx context.Context, query *categoryservice.ListQuery) ([]categorymodel.Category, error)
}

type CategoryHTTPController struct {
	createCmdHdl    ICreateCommandHandler
	getDetailQryHdl IDetailQueryHandler
	updateCmdHdl    IUpdateByIdCommandHandler
	deleteCmdHdl    IDeleteByIdCommandHandler
	listQryHdl      IListQueryHandler
}

func NewCategoryHTTPController(
	createCmdHdl ICreateCommandHandler,
	getDetailQryHdl IDetailQueryHandler,
	updateCmdHdl IUpdateByIdCommandHandler,
	deleteCmdHdl IDeleteByIdCommandHandler,
	listQryHdl IListQueryHandler,
) *CategoryHTTPController {
	return &CategoryHTTPController{
		getDetailQryHdl: getDetailQryHdl,
		createCmdHdl:    createCmdHdl,
		updateCmdHdl:    updateCmdHdl,
		deleteCmdHdl:    deleteCmdHdl,
		listQryHdl:      listQryHdl,
	}
}

func (ctl *CategoryHTTPController) SetupRoutes(g *gin.RouterGroup) {
	g.POST("", ctl.CreateCategoryAPI)
	g.GET("", ctl.ListCategoryAPI)
	g.GET("/:id", ctl.GetCategoryByIdAPI)
	g.PATCH("/:id", ctl.UpdateCategoryByIdAPI)
	g.DELETE("/:id", ctl.DeleteCategoryByIdAPI)
}

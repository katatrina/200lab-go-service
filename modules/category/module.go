package categorymodule

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCategoryModule(db *gorm.DB, g *gin.RouterGroup) {
	catCtl := InitializeCategoryController(db)

	catCtl.SetupRoutes(g)
}

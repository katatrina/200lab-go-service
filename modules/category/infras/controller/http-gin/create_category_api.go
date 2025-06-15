package categoryhttpgin

import (
	categorymodel "go12-service/modules/category/model"
	categoryservice "go12-service/modules/category/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *CategoryHTTPController) CreateCategoryAPI(c *gin.Context) {
	var requestBodyData categorymodel.Category

	if err := c.ShouldBindJSON(&requestBodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call business logic in service
	cmd := categoryservice.CreateCommand{Dto: requestBodyData}

	id, err := ctl.createCmdHdl.Execute(c.Request.Context(), &cmd)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": id})
}

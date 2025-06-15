package categoryhttpgin

import (
	categoryservice "go12-service/modules/category/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctl *CategoryHTTPController) ListCategoryAPI(c *gin.Context) {
	var dto categoryservice.ListQuery

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto.Process()

	data, err := ctl.listQryHdl.Execute(c.Request.Context(), &dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data, "paging": dto.PagingDTO, "filter": dto.FilterCategoryDTO})
}

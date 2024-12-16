package oneiromancy

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/oneironmancy/request"
	"github.com/gin-gonic/gin"
)

type QueryAPI struct {
}

// Query
// @Tags Oneiromancy
// @Summary 查询梦境
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param query query string true "梦境关键词"
// @Success 200 {object} response.Response{data=string} "查询成功"
// @Router /oneiromancy/query [post]
func (a *QueryAPI) Query(c *gin.Context) {
	var reqBody request.QueryRequestBody
	// 将请求体绑定到结构体上
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		// 如果绑定失败，返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := QueryService.Query(reqBody.Query)
	response.OkWithData(result, c)
}

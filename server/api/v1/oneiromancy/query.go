package oneiromancy

import (
	"bufio"
	"fmt"
	"io"
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/oneironmancy/request"
	"github.com/gin-gonic/gin"
)

type QueryAPI struct {
}

// Query
// @Tags		Oneiromancy
// @Summary	查询梦境
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce	application/json
// @Param		query	query		string							true	"梦境关键词"
// @Success	200		{object}	response.Response{data=string}	"查询成功"
// @Router		/oneiromancy/query [post]
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

func (a *QueryAPI) Query_stream(c *gin.Context) {
	query := c.Query("query")
	resp, err := QueryService.Query2(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// 设置流式响应头
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	// 逐行读取并写入客户端
	reader := bufio.NewReader(resp.Body)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err != io.EOF {
				fmt.Printf("error reading stream: %v\n", err)
			}
			break
		}
		var formattedData string
		// 检查是否读取到了换行符，如果是则发送数据
		if r == '\n' {
			nextRune, _, nextErr := reader.ReadRune()
			if nextErr == nil && nextRune == '\n' {
				formattedData = "data: <br>\n\n"
			} else {
				// 若没有连续换行，则回退读取的字符
				_ = reader.UnreadRune()
				formattedData = "data: \n\n"
			}
		} else {
			formattedData = fmt.Sprintf("data: %c\n\n", r)
		}
		// formattedData := fmt.Sprintf("data: %c\n\n", r)
		_, writeErr := c.Writer.WriteString(formattedData)
		if writeErr != nil {
			fmt.Printf("error writing stream: %v\n", writeErr)
			break
		}
		c.Writer.Flush()
	}
}

package oneiromancy

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/request"
)

type QueryService struct {
}

func (s *QueryService) Query(query string) string {
	resp, err := request.HttpRequest(
		"http://127.0.0.1:80/ai_core/api/query",
		"POST",
		nil,
		nil,
		map[string]string{
			"userInput": query,
		},
	)
	if err != nil {
		return fmt.Sprintf("error making request: %v", err)
	}
	var resStruct response.Response
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("error reading response body: %v", err)
	}
	defer resp.Body.Close()
	json.Unmarshal(body, &resStruct)
	return resStruct.Data.(string)
}

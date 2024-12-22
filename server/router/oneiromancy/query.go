package oneiromancy

import "github.com/gin-gonic/gin"

type OneiromancyRouter struct {
}

func (s *OneiromancyRouter) InitOneiromancyRouter(Router *gin.RouterGroup) {
	oneiromancyRouter := Router.Group("oneiromancy")
	{
		oneiromancyRouter.POST("query", QueryAPI.Query)
		oneiromancyRouter.GET("query_stream", QueryAPI.Query_stream)
	}
}

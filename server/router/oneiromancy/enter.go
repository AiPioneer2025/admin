package oneiromancy

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	OneiromancyRouter
}

var (
	QueryAPI = api.ApiGroupApp.OneiromancyApiGroup.QueryAPI
)

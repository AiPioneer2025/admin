package oneiromancy

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	QueryAPI
}

var (
	QueryService = service.ServiceGroupApp.OneironmancyServiceGroup.QueryService
)

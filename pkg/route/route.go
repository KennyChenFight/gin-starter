package route

import (
	"github.com/KennyChenFight/gin-starter/pkg/service"
	"github.com/gin-gonic/gin"
)

func InitRoutingRule(engine *gin.Engine, svc *service.BaseService) *gin.Engine {
	v1Group := engine.Group("/v1")
	{
		v1Group.POST("/members", svc.CreateMember)
		v1Group.GET("/members/:id", svc.GetMember)
		v1Group.PUT("/members/:id", svc.UpdateMember)
		v1Group.DELETE("/members/:id", svc.DeleteMember)
	}
	return engine
}

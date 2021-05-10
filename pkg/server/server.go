package server

import (
	"net/http"

	"github.com/KennyChenFight/gin-starter/pkg/middleware"

	"github.com/KennyChenFight/gin-starter/pkg/service"
	"github.com/gin-gonic/gin"
)

func NewHTTPServer(engine *gin.Engine, port string, mwe *middleware.BaseMiddleware, svc *service.BaseService) *http.Server {
	return &http.Server{
		Addr:    port,
		Handler: registerRoutingRule(engine, mwe, svc),
	}
}

func registerRoutingRule(engine *gin.Engine, mwe *middleware.BaseMiddleware, svc *service.BaseService) *gin.Engine {
	engine.Use(mwe.GlobalErrorHandle())
	engine.NoMethod(svc.HandleMethodNotAllowed)
	engine.NoRoute(svc.HandlePathNotFound)
	v1Group := engine.Group("/v1")
	{
		v1Group.POST("/members", svc.CreateMember)
		v1Group.GET("/members/:id", svc.GetMember)
		v1Group.PUT("/members/:id", svc.UpdateMember)
		v1Group.DELETE("/members/:id", svc.DeleteMember)
	}
	return engine
}

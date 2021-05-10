package service

import (
	"net/http"

	"github.com/KennyChenFight/gin-starter/pkg/business"
	"github.com/KennyChenFight/gin-starter/pkg/dao"
	"github.com/gin-gonic/gin"
)

type BaseService struct {
	memberDAO dao.MemberDAO
}

func NewService(memberDAO dao.MemberDAO) *BaseService {
	return &BaseService{memberDAO: memberDAO}
}

func (s *BaseService) HandleMethodNotAllowed(c *gin.Context) {
	s.responseWithError(c, business.NewError(business.MethodNowAllowed, http.StatusMethodNotAllowed, "http method not allowed", nil))
}

func (s *BaseService) HandlePathNotFound(c *gin.Context) {
	s.responseWithError(c, business.NewError(business.PathNotFound, http.StatusNotFound, "http path not found", nil))
}

func (s *BaseService) responseWithError(c *gin.Context, businessError *business.Error) {
	c.Abort()
	c.Error(businessError)
}

func (s *BaseService) responseWithSuccess(c *gin.Context, businessSuccess *business.Success) {
	c.Set("success", businessSuccess)
}

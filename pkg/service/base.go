package service

import (
	"github.com/KennyChenFight/gin-starter/pkg/dao"
	"github.com/KennyChenFight/gin-starter/pkg/util"
	"github.com/gin-gonic/gin"
)

type BaseService struct {
	memberDAO            dao.MemberDAO
	validationTranslator *util.ValidationTranslator
}

func NewService(memberDAO dao.MemberDAO, validatorTranslator *util.ValidationTranslator) *BaseService {
	return &BaseService{memberDAO: memberDAO, validationTranslator: validatorTranslator}
}

func sendErrorResponse(c *gin.Context, service *BaseService, businessError *util.BusinessError) {
	translated, err := service.validationTranslator.Translate(c.GetHeader("Accept-Language"), businessError.Reason)
	if translated == nil && err == nil {
		c.JSON(businessError.HTTPStatusCode, businessError)
		return
	}

	// translate fail, you can logging.
	if err != nil {
		c.JSON(businessError.HTTPStatusCode, businessError)
	} else {
		businessError.ValidationErrors = translated
		c.JSON(businessError.HTTPStatusCode, businessError)
	}
}

func sendSuccessResponse(c *gin.Context, statusCode int, response interface{}) {
	c.JSON(statusCode, response)
}

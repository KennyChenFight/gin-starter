package middleware

import (
	"github.com/KennyChenFight/gin-starter/pkg/business"
	"github.com/KennyChenFight/gin-starter/pkg/validation"
	"github.com/KennyChenFight/golib/loglib"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewMiddleware(logger *loglib.Logger, validationTranslator *validation.ValidationTranslator) *BaseMiddleware {
	return &BaseMiddleware{logger: logger, validationTranslator: validationTranslator}
}

type BaseMiddleware struct {
	logger               *loglib.Logger
	validationTranslator *validation.ValidationTranslator
}

func (b *BaseMiddleware) sendErrorResponse(c *gin.Context, businessError *business.Error) {
	translated, err := b.validationTranslator.Translate(c.GetHeader("Accept-Language"), businessError.Reason)
	if translated == nil && err == nil {
		c.JSON(businessError.HTTPStatusCode, businessError)
		return
	}

	if err != nil {
		b.logger.Error("fail to translate validation message", zap.Error(err))
		c.JSON(businessError.HTTPStatusCode, businessError)
	} else {
		businessError.ValidationErrors = translated
		c.JSON(businessError.HTTPStatusCode, businessError)
	}
}

func (b *BaseMiddleware) sendSuccessResponse(c *gin.Context, success *business.Success) {
	c.JSON(success.HTTPStatusCode, success.Response)
}

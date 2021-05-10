package middleware

import (
	"net/http"

	"github.com/KennyChenFight/gin-starter/pkg/business"
	"github.com/gin-gonic/gin"
)

func (b *BaseMiddleware) GlobalErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[len(c.Errors)-1].Err
			switch err.(type) {
			case *business.Error:
				businessError := err.(*business.Error)
				b.sendErrorResponse(c, businessError)
				return
			default:
				unknownError := err.(error)
				businessError := business.NewError(business.Unknown, http.StatusInternalServerError, unknownError.Error(), unknownError)
				b.sendErrorResponse(c, businessError)
				return
			}
		}
		success := c.MustGet("success").(*business.Success)
		b.sendSuccessResponse(c, success)
	}
}

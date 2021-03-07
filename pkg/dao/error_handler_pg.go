package dao

import (
	"log"
	"net/http"

	"github.com/KennyChenFight/gin-starter/pkg/util"
)

const (
	PGErrMsgNoRowsFound      = "pg: no rows in result set"
	PGErrMsgNoMultiRowsFound = "pg: multiple rows in result set"
)

func pgErrorHandle(err error) *util.BusinessError {
	switch err.Error() {
	case PGErrMsgNoRowsFound:
		return util.NewBusinessError(util.NotFound, http.StatusNotFound, "record not found", err)
	case PGErrMsgNoMultiRowsFound:
		return util.NewBusinessError(util.NotFound, http.StatusNotFound, "multi records not found", err)
	default:
		log.Printf("[pgErrorHandle]:%v", err)
		return util.NewBusinessError(util.Unknown, http.StatusInternalServerError, "internal error", err)
	}
}

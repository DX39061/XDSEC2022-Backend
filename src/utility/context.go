package utility

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var ErrIDIsBroken = errors.New("id is broken")

func GetUintFromContext(ctx *gin.Context, param string) (uint, error) {
	dataRaw, exists := ctx.Get(param)
	if !exists {
		return 0, ErrIDIsBroken
	}
	data, ok := dataRaw.(uint)
	if !ok {
		return 0, ErrIDIsBroken
	}
	return data, nil
}

func GetUintFromUrlParam(ctx *gin.Context, param string) (uint, error) {
	idStr := ctx.Param(param)
	return AtoU(idStr)
}

func GetIntFromUrlQuery(ctx *gin.Context, param string) (int, error) {
	idStr := ctx.Query(param)
	return strconv.Atoi(idStr)
}

func GetUintFromUrlQuery(ctx *gin.Context, param string) (uint, error) {
	idStr := ctx.Query(param)
	return AtoU(idStr)
}

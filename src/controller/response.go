package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

const (
	SigOk                 = 0x00
	SigInternalFailed     = 0xFF
	SigAuthFailed         = 0xF0
	SigAccessDenied       = 0xF1
	SigRobotTestFailed    = 0xF2
	SigAccountNotVerified = 0xF3
)

func sendResponse(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

//goland:noinspection GoUnusedExportedFunction
func Ok(c *gin.Context) {
	sendResponse(c, SigOk, nil, "succeeded")
}

//goland:noinspection GoUnusedExportedFunction
func OkWithMessage(c *gin.Context, message string) {
	sendResponse(c, SigOk, nil, message)
}

//goland:noinspection GoUnusedExportedFunction
func OkWithData(c *gin.Context, data interface{}) {
	sendResponse(c, SigOk, data, "succeeded")
}

//goland:noinspection GoUnusedExportedFunction
func OkWithDetail(c *gin.Context, data interface{}, message string) {
	sendResponse(c, SigOk, data, message)
}

//goland:noinspection GoUnusedExportedFunction
func InternalFailed(c *gin.Context) {
	sendResponse(c, SigInternalFailed, nil, "failed")
}

//goland:noinspection GoUnusedExportedFunction
func InternalFailedWithMessage(c *gin.Context, message string) {
	sendResponse(c, SigInternalFailed, nil, message)
}

//goland:noinspection GoUnusedExportedFunction
func InternalFailedWithDetail(c *gin.Context, data interface{}, message string) {
	sendResponse(c, SigInternalFailed, data, message)
}

//goland:noinspection GoUnusedExportedFunction
func AuthFailed(c *gin.Context) {
	sendResponse(c, SigAuthFailed, nil, "please login first")
}

//goland:noinspection GoUnusedExportedFunction
func AccessDenied(c *gin.Context) {
	sendResponse(c, SigAccessDenied, nil, "access denied")
}

//goland:noinspection GoUnusedExportedFunction
func RobotTestFailed(c *gin.Context) {
	sendResponse(c, SigRobotTestFailed, nil, "hey robot")
}

//goland:noinspection GoUnusedExportedFunction
func AccountNotVerified(c *gin.Context) {
	sendResponse(c, SigAccountNotVerified, nil, "please verify your account first")
}

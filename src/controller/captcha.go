package controller

import (
	"XDSEC2022-Backend/src/captcha"
	"github.com/gin-gonic/gin"
)

func init() {
	RegisterApiRoute(func(router *gin.RouterGroup) {
		captchaRouter := router.Group("/captcha")
		captchaRouter.POST("", verifyCaptchaHandler)
	})
}

type FrontendVerifyRequest struct {
	Token string `json:"token"`
}

func verifyCaptchaHandler(ctx *gin.Context) {
	var req FrontendVerifyRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	verified, err := captcha.VerifyCaptcha(req.Token)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	OkWithData(ctx, verified)
}

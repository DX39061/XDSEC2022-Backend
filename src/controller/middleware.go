package controller

import "github.com/gin-gonic/gin"

func LoginRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetBool("login") {
			ctx.Next()
			return
		} else {
			AuthFailed(ctx)
			ctx.Abort()
			return
		}
	}
}

func PrepareInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := ExtractToken(ctx)
		if err != nil {
			ctx.Set("login", false)
		} else {
			ctx.Set("login", true)
		}
		ctx.Next()
	}
}

func AdminRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetBool("admin") {
			ctx.Next()
			return
		} else {
			AccessDenied(ctx)
			ctx.Abort()
			return
		}
	}
}

func DontLoginRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !ctx.GetBool("login") {
			ctx.Next()
			return
		} else {
			InternalFailedWithMessage(ctx, "You are already logged in")
			ctx.Abort()
			return
		}
	}
}

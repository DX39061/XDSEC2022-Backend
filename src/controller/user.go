package controller

import (
	"XDSEC2022-Backend/src/game_bind"
	"XDSEC2022-Backend/src/repository"
	"XDSEC2022-Backend/src/utility"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func init() {
	RegisterApiRoute(func(router *gin.RouterGroup) {
		adminUserRouter := router.Group("/users", LoginRequired(), AdminRequired())
		{
			adminUserRouter.GET("/count", UserCountHandler)
			adminUserRouter.GET("/search", UserSearchHandler) // keyword = ?
			adminUserRouter.GET("", UserGetListHandler)
			adminUserRouter.GET("/:user", UserGetDetailHandler)
			adminUserRouter.PATCH("/:user", UserChangeProfileHandler)
			adminUserRouter.PATCH("/:user/reset-password", UserResetPasswordHandler)
			adminUserRouter.GET("/:user/bind-game", UserBindGameHandler)
			adminUserRouter.DELETE("/:user", UserDeleteHandler)
		}
	})
}

func UserCountHandler(ctx *gin.Context) {
	count, err := repository.GetUserCount()
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	OkWithData(ctx, count)
}

func UserSearchHandler(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	usersShort, err := repository.SearchUsers(keyword)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	OkWithData(ctx, usersShort)
}

func UserGetListHandler(ctx *gin.Context) {
	limit, err := utility.GetIntFromUrlQuery(ctx, "limit")
	if err != nil {
		limit = 10
	}
	skip, err := utility.GetIntFromUrlQuery(ctx, "skip")
	if err != nil {
		skip = 0
	}
	usersShort, err := repository.GetUserList(limit, skip)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	OkWithData(ctx, usersShort)
}

func UserGetDetailHandler(ctx *gin.Context) {
	userID, err := utility.GetUintFromUrlParam(ctx, "user")
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	userDetail, err := repository.GetUserDetail(userID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	OkWithData(ctx, userDetail)
}

type ChangeUserProfileRequest struct {
	Admin              string `json:"admin"`
	Name               string `json:"name"`
	Sex                string `json:"sex"`
	Major              string `json:"major"` // 学院/专业
	StudentID          string `json:"student-id"`
	Telephone          string `json:"telephone"`
	Email              string `json:"email"`
	Department         string `json:"department"`          // 意向部门
	Direction          string `json:"direction"`           // 学习方向
	LearnedTechnique   string `json:"learned-technique"`   // 技术基础
	LearningExperience string `json:"learning-experience"` // 学习经历
	HobbyAndAdvantage  string `json:"hobby-and-advantage"` // 爱好特长
}

func UserChangeProfileHandler(ctx *gin.Context) {
	var req ChangeUserProfileRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	userID, err := utility.GetUintFromUrlParam(ctx, "user")
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	user, err := repository.GetUserByID(userID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	err = copier.Copy(&user, req)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	err = repository.UpdateUser(&user)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	Ok(ctx)
}

type ResetUserPasswordRequest struct {
	NewPassword string `json:"new-password"`
}

func UserResetPasswordHandler(ctx *gin.Context) {
	var req ResetUserPasswordRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	password, err := utility.HashPassword(req.NewPassword)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	userID, err := utility.GetUintFromUrlParam(ctx, "user")
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	user, err := repository.GetUserByID(userID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	user.Password = password
	err = repository.UpdateUser(&user)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	Ok(ctx)
}

func UserBindGameHandler(ctx *gin.Context) {
	userID, err := utility.GetUintFromUrlParam(ctx, "user")
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	user, err := repository.GetUserByID(userID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	gameData, err := game_bind.GetUserDataOfGame(user.StudentID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	OkWithData(ctx, gameData)
}

func UserDeleteHandler(ctx *gin.Context) {
	userID, err := utility.GetUintFromUrlParam(ctx, "user")
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	err = repository.DeleteUser(userID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	Ok(ctx)
}

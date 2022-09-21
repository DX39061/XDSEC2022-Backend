package controller

import (
	"XDSEC2022-Backend/src/model"
	"XDSEC2022-Backend/src/repository"
	"XDSEC2022-Backend/src/utility"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func init() {
	RegisterApiRoute(func(router *gin.RouterGroup) {
		adminInterviewRouter := router.Group("/interviews", LoginRequired(), AdminRequired())
		{
			adminInterviewRouter.GET("/count", GetInterviewCountHandler)
			adminInterviewRouter.GET("", GetInterviewListHandler)
			adminInterviewRouter.POST("", CreateInterviewHandler)
			adminInterviewRouter.GET("/:interview", GetInterviewDetailHandler)
			adminInterviewRouter.GET("/of-user", GetInterviewOfUserHandler) // user-id = ?
			adminInterviewRouter.PATCH("/:interview", UpdateInterviewHandler)
			adminInterviewRouter.DELETE("/:interview", DeleteInterviewHandler)
		}
	})
}

func GetInterviewCountHandler(ctx *gin.Context) {
	count, err := repository.GetInterviewCount()
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	OkWithData(ctx, count)
}

func GetInterviewListHandler(ctx *gin.Context) {
	limit, err := utility.GetIntFromUrlQuery(ctx, "limit")
	if err != nil {
		limit = 10
	}
	skip, err := utility.GetIntFromUrlQuery(ctx, "skip")
	if err != nil {
		skip = 0
	}
	interviewsShort, err := repository.GetInterviewList(limit, skip)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	OkWithData(ctx, interviewsShort)
}

type CreateInterviewRequest struct {
	Round         uint   `json:"round"`
	Pass          bool   `json:"pass"`
	Record        string `json:"record"`
	InterviewerID uint   `json:"interviewer-id"`
	IntervieweeID uint   `json:"interviewee-id"`
}

func CreateInterviewHandler(ctx *gin.Context) {
	var req CreateInterviewRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	var interview model.Interview
	err = copier.Copy(&interview, req)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	interviewer, err := repository.GetUserByID(interview.InterviewerID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	interviewee, err := repository.GetUserByID(interview.IntervieweeID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	interview.Interviewer = interviewer.NickName
	interview.Interviewee = interviewee.NickName
	err = repository.CreateInterview(&interview)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	Ok(ctx)
}

func GetInterviewDetailHandler(ctx *gin.Context) {
	interviewID, err := utility.GetUintFromUrlParam(ctx, "interview")
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	interviewDetail, err := repository.GetInterviewDetail(interviewID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	OkWithData(ctx, interviewDetail)
}

func GetInterviewOfUserHandler(ctx *gin.Context) {
	userID, err := utility.GetUintFromUrlQuery(ctx, "user-id")
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	interviewsDetail, err := repository.GetInterviewDetailOfUser(userID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	OkWithData(ctx, interviewsDetail)
}

type UpdateInterviewRequest struct {
	Pass   bool   `json:"pass"`
	Record string `json:"record"`
}

func UpdateInterviewHandler(ctx *gin.Context) {
	var req UpdateInterviewRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	interviewID, err := utility.GetUintFromUrlParam(ctx, "interview")
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	interview, err := repository.GetInterview(interviewID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	err = copier.Copy(&interview, req)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	err = repository.UpdateInterview(&interview)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	Ok(ctx)
}

func DeleteInterviewHandler(ctx *gin.Context) {
	interviewID, err := utility.GetUintFromUrlParam(ctx, "interview")
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	err = repository.DeleteInterview(interviewID)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	Ok(ctx)
}

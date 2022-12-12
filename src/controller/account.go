package controller

import (
    "XDSEC2022-Backend/src/cache"
    "XDSEC2022-Backend/src/captcha"
    "XDSEC2022-Backend/src/email"
    "XDSEC2022-Backend/src/model"
    "XDSEC2022-Backend/src/repository"
    "XDSEC2022-Backend/src/utility"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/copier"
)

func init() {
    RegisterApiRoute(func(router *gin.RouterGroup) {
        accountRouter := router.Group("/account", DontLoginRequired())
        {
            accountRouter.POST("/join", AccountJoinHandler)
            accountRouter.POST("/join/verifyemail", AccountVerifyEmailHandler)
            accountRouter.POST("/login", AccountLoginHandler)
        }
        authorizedAccountRouter := router.Group("/account", LoginRequired())
        {
            authorizedAccountRouter.DELETE("/logout", AccountLogoutHandler)
            authorizedAccountRouter.GET("/profile", AccountGetSelfProfileHandler)
            authorizedAccountRouter.PATCH("/profile", AccountChangeProfileHandler)
            authorizedAccountRouter.PATCH("/change-password", AccountChangePasswordHandler)
        }
    })
}

type LoginRequest struct {
    Account      string `json:"account"`
    Password     string `json:"password"`
    CaptchaToken string `json:"captcha-token"`
}

func AccountLoginHandler(ctx *gin.Context) {
    var req LoginRequest
    err := ctx.ShouldBind(&req)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    captchaVerified, err := captcha.VerifyCaptcha(req.CaptchaToken)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    if !captchaVerified {
        RobotTestFailed(ctx)
        ctx.Abort()
        return
    }
    user, err := repository.ValidateUserPassword(req.Account, req.Password)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    loginUser(ctx, user)
    Ok(ctx)
}

type JoinRequest struct {
    NickName           string `json:"nick-name"`
    Name               string `json:"name"`
    Sex                string `json:"sex"`
    Major              string `json:"major"` // 学院/专业
    StudentID          string `json:"student-id"`
    Telephone          string `json:"telephone"`
    QQ                 string `json:"qq"`
    Email              string `json:"email"`
    Password           string `json:"password"`
    Department         string `json:"department"`          // 意向部门
    Direction          string `json:"direction"`           // 学习方向
    LearnedTechnique   string `json:"learned-technique"`   // 技术基础
    LearningExperience string `json:"learning-experience"` // 学习经历
    HobbyAndAdvantage  string `json:"hobby-and-advantage"` //爱好特长
    VerifyEmail        string `json:"verify-email"`
    CaptchaToken       string `json:"captcha-token"`
}

type EmailRequest struct {
    Email string `json:"verify-email"`
    // CaptchaToken string `json:"captcha-token"`
}

func AccountVerifyEmailHandler(ctx *gin.Context) {
    var req EmailRequest
    if err := ctx.ShouldBind(&req); err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    if !utility.VerifyEmailFormat(req.Email) {
        InternalFailedWithMessage(ctx, "invalid format")
        ctx.Abort()
        return
    }
    code, _ := utility.GetRandString()
    err := email.SendVerifyCode(req.Email, code)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
    }
    err = cache.PermitEmail(req.Email, code)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
    }
    go cache.DeleteEmail(req.Email)
    Ok(ctx)
}

func AccountJoinHandler(ctx *gin.Context) {
    var req JoinRequest
    if err := ctx.ShouldBind(&req); err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    captchaVerified, err := captcha.VerifyCaptcha(req.CaptchaToken)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    if !captchaVerified {
        RobotTestFailed(ctx)
        ctx.Abort()
        return
    }
    if !utility.VerifyEmailFormat(req.Email) ||
        !utility.VerifyTelephoneFormat(req.Telephone) ||
        !utility.VerifyPasswordFormat(req.Password) ||
        !utility.VerifyStudentIDFormat(req.StudentID) {
        InternalFailedWithMessage(ctx, "invalid format")
        ctx.Abort()
        return
    }
    if !cache.ValidateEmail(req.Email, req.VerifyEmail) {
        InternalFailedWithMessage(ctx, "验证邮箱失败")
        ctx.Abort()
        return
    }
    var user model.User
    err = copier.Copy(&user, req)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    user.Admin = false
    hashedPassword, err := utility.HashPassword(user.Password)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    user.Password = hashedPassword
    err = repository.CreateUser(&user)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    loginUser(ctx, user)
    Ok(ctx)
}

func AccountLogoutHandler(ctx *gin.Context) {
    _ = RemoveToken(ctx)
    Ok(ctx)
}

func AccountGetSelfProfileHandler(ctx *gin.Context) {
    userID, err := utility.GetUintFromContext(ctx, "id")
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

type ChangeSelfProfileRequest struct {
    NickName           string `json:"nick-name"`
    Name               string `json:"name"`
    Sex                string `json:"sex"`
    Major              string `json:"major"` // 学院/专业
    StudentID          string `json:"student-id"`
    Telephone          string `json:"telephone"`
    QQ                 string `json:"qq"`
    Email              string `json:"email"`
    Department         string `json:"department"`          // 意向部门
    Direction          string `json:"direction"`           // 学习方向
    LearnedTechnique   string `json:"learned-technique"`   // 技术基础
    LearningExperience string `json:"learning-experience"` // 学习经历
    HobbyAndAdvantage  string `json:"hobby-and-advantage"` // 爱好特长
}

func AccountChangeProfileHandler(ctx *gin.Context) {
    var req ChangeSelfProfileRequest
    err := ctx.ShouldBind(&req)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    if !utility.VerifyEmailFormat(req.Email) ||
        !utility.VerifyTelephoneFormat(req.Telephone) ||
        !utility.VerifyStudentIDFormat(req.StudentID) {
        InternalFailedWithMessage(ctx, "invalid format")
        ctx.Abort()
        return
    }
    userID, err := utility.GetUintFromContext(ctx, "id")
    if err != nil {
        AuthFailed(ctx)
        _ = RemoveToken(ctx)
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

type ChangePasswordRequest struct {
    OldPassword string `json:"old-password"`
    NewPassword string `json:"new-password"`
}

func AccountChangePasswordHandler(ctx *gin.Context) {
    userID, err := utility.GetUintFromContext(ctx, "id")
    if err != nil {
        AuthFailed(ctx)
        _ = RemoveToken(ctx)
        ctx.Abort()
        return
    }
    var req ChangePasswordRequest
    if err := ctx.ShouldBind(&req); err != nil {
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
    if !utility.CheckPasswordHash(req.OldPassword, user.Password) {
        InternalFailedWithMessage(ctx, "old password not match")
        ctx.Abort()
        return
    }
    if !utility.VerifyPasswordFormat(req.NewPassword) {
        InternalFailedWithMessage(ctx, "new password format error")
        ctx.Abort()
        return
    }
    hashedPassword, err := utility.HashPassword(req.NewPassword)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    user.Password = hashedPassword
    err = repository.UpdateUser(&user)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    Ok(ctx)
}

func loginUser(ctx *gin.Context, user model.User) {
    err := DistributeToken(ctx, TokenClaims{
        UserID:   user.ID,
        NickName: user.NickName,
        Admin:    user.Admin,
    })
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
    }
}

package base

import (
    "github.com/astaxie/beego/context"
    "github.com/astaxie/beego/session"
    "strings"
    "github.com/thanzen/eq/utils"
    "github.com/astaxie/beego"
    "github.com/thanzen/eq/models/user"
    "github.com/thanzen/eq/setting"
)

type UserBaseController struct{
    BaseController
}


//todo: move out this to user controller

// get login redirect url from cookie
func (this *UserBaseController) GetLoginRedirect(ctx *context.Context) string {
    loginRedirect := strings.TrimSpace(ctx.GetCookie("login_to"))
    if utils.IsMatchHost(loginRedirect) == false {
        loginRedirect = "/"
    } else {
        ctx.SetCookie("login_to", "", -1, "/")
    }
    return loginRedirect
}

// login user
//todo: added token (oauth2) session if it is rest api call
func (this *UserBaseController) LoginUser(u *user.User, ctx *context.Context, remember bool) {
    // werid way of beego session regenerate id...
    ctx.Input.CruSession.SessionRelease(ctx.ResponseWriter)
    ctx.Input.CruSession = beego.GlobalSessions.SessionRegenerateId(ctx.ResponseWriter, ctx.Request)
    ctx.Input.CruSession.Set("auth_user_id", u.Id)

    if remember {
        this.WriteRememberCookie(u, ctx)
    }
}
func (this *UserBaseController) WriteRememberCookie(u *user.User, ctx *context.Context) {
    secret := utils.EncodeMd5(u.PasswordSalt + u.Password)
    days := 86400 * setting.LoginRememberDays
    ctx.SetCookie(setting.CookieUsername, u.Username, days)
    ctx.SetSecureCookie(secret, setting.CookieRememberName, u.Username, days)
}
func (this *UserBaseController) DeleteRememberCookie(ctx *context.Context) {
    ctx.SetCookie(setting.CookieUsername, "", -1)
    ctx.SetCookie(setting.CookieRememberName, "", -1)
}

func (this *UserBaseController) LoginUserFromRememberCookie(u *user.User, ctx *context.Context) (success bool) {
    userName := ctx.GetCookie(setting.CookieUsername)
    if len(userName) == 0 {
        return false
    }

    defer func() {
        if !success {
            this.DeleteRememberCookie(ctx)
        }
    }()
    u.Username = userName
    if err := this.UserService.Read(u, "Username"); err != nil {
        return false
    }

    secret := utils.EncodeMd5(u.PasswordSalt + u.Password)
    value, _ := ctx.GetSecureCookie(secret, setting.CookieRememberName)
    if value != userName {
        return false
    }

    this.LoginUser(u, ctx, true)

    return true
}

// logout user
func (this *UserBaseController) LogoutUser(ctx *context.Context) {
    this.DeleteRememberCookie(ctx)
    ctx.Input.CruSession.Delete("auth_user_id")
    ctx.Input.CruSession.Flush()
    beego.GlobalSessions.SessionDestroy(ctx.ResponseWriter, ctx.Request)
}

func (this *UserBaseController) GetUserIdFromSession(sess session.SessionStore) int64 {
    if id, ok := sess.Get("auth_user_id").(int64); ok && id > 0 {
        return id
    }
    return 0
}

// get user if key exist in session
func (this *UserBaseController) GetUserFromSession(u *user.User, sess session.SessionStore) bool {
    id := this.GetUserIdFromSession(sess)
    if id > 0 {
        u.Id = id
        if this.UserService.Read(u) == nil {
            return true
        }
    }

    return false
}
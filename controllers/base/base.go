package base

import (
	"fmt"
	"html/template"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/thanzen/eq/models/user"
	userServ "github.com/thanzen/eq/services/user"
	"github.com/thanzen/eq/setting"
	"github.com/thanzen/eq/utils"
    "github.com/astaxie/beego/session"
    "github.com/astaxie/beego/context"
)

type NestPreparer interface {
	NestPrepare()
}

// BaseController implemented global settings for all other routers.
type BaseController struct {
	beego.Controller
	i18n.Locale
	User        user.User
	IsLogin     bool
	UserService *userServ.UserService
}

// Prepare implemented Prepare method for baseRouter.
func (this *BaseController) Prepare() {

	//initialize UserService
	this.UserService = &userServ.UserService{}

	if setting.EnforceRedirect {
		// if the host not matching app settings then redirect to AppUrl
		beego.Info(this.Ctx.Request.Host, " and ", setting.AppHost)
		if this.Ctx.Request.Host != setting.AppHost {
			this.Redirect(setting.AppUrl, 302)
			return
		}
	}

	// page start time
	this.Data["PageStartTime"] = time.Now()

	// start session
	this.StartSession()

	// check flash redirect, if match url then end, else for redirect return
	if match, redir := this.CheckFlashRedirect(this.Ctx.Request.RequestURI); redir {
		return
	} else if match {
		this.EndFlashRedirect()
	}

	switch {
	// save logined user if exist in session
	case this.GetUserFromSession(&this.User, this.CruSession):
		this.IsLogin = true
	// save logined user if exist in remember cookie
	case this.LoginUserFromRememberCookie(&this.User, this.Ctx):
		this.IsLogin = true
		//todo: add token based(oauth2) authentication user retrive
	}


	if this.IsLogin {
		this.IsLogin = true
		this.Data["User"] = &this.User
		this.Data["IsLogin"] = this.IsLogin

		// if user forbided then do logout
		if !this.User.Active || this.User.Deleted {
			this.LogoutUser(this.Ctx)
			this.FlashRedirect("/login", 302, "UserForbid")
			return
		}
	}

	// Setting properties.
	this.Data["AppName"] = beego.AppName
	this.Data["AppVer"] = setting.AppVer
	this.Data["AppUrl"] = setting.AppUrl
	this.Data["AppLogo"] = setting.AppLogo
	this.Data["IsProMode"] = setting.IsProMode
	this.Data["AvatarUrl"] = setting.AvatarUrl

	// Redirect to make URL clean.
	if this.setLang() {
		i := strings.Index(this.Ctx.Request.RequestURI, "?")
		this.Redirect(this.Ctx.Request.RequestURI[:i], 302)
		return
	}

	// read flash message
	beego.ReadFromRequest(&this.Controller)

	// pass xsrf helper to template context
	xsrfToken := this.XsrfToken()
	this.Data["xsrf_token"] = xsrfToken
	this.Data["xsrf_html"] = template.HTML(this.XsrfFormHtml())

	// if method is GET then auto create a form once token
	if this.Ctx.Request.Method == "GET" {
		this.FormOnceCreate()
	}

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

// on router finished
func (this *BaseController) Finish() {

}

func (this *BaseController) LoginUser(u *user.User, remember bool) string {
	loginRedirect := strings.TrimSpace(this.Ctx.GetCookie("login_to"))
	if utils.IsMatchHost(loginRedirect) == false {
		loginRedirect = "/"
	} else {
		this.Ctx.SetCookie("login_to", "", -1, "/")
	}

	// login user
	this.LoginUserRememberCookie(u, this.Ctx, remember)

	//todo: add locale support, rewrite i18n.go
	this.setLangCookie(u.Lang)

	return loginRedirect
}

// check if user not active then redirect
func (this *BaseController) CheckActiveRedirect(args ...interface{}) bool {
	var redirect_to string
	code := 302
	needActive := true
	for _, arg := range args {
		switch v := arg.(type) {
		case bool:
			needActive = v
		case string:
			// custom redirect url
			redirect_to = v
		case int:
			code = v
		}
	}
	if needActive {
		// check login
		if this.CheckLoginRedirect() {
			return true
		}

		// redirect to active page
		if !this.User.Active {
			this.FlashRedirect("/settings/profile", code, "NeedActive")
			return true
		}
	} else {
		// no need active
		if this.User.Active {
			if redirect_to == "" {
				redirect_to = "/"
			}
			this.Redirect(redirect_to, code)
			return true
		}
	}
	return false

}

// check if not login then redirect
func (this *BaseController) CheckLoginRedirect(args ...interface{}) bool {
	var redirect_to string
	code := 302
	needLogin := true
	for _, arg := range args {
		switch v := arg.(type) {
		case bool:
			needLogin = v
		case string:
			// custom redirect url
			redirect_to = v
		case int:
			// custom redirect url
			code = v
		}
	}

	// if need login then redirect
	if needLogin && !this.IsLogin {
		if len(redirect_to) == 0 {
			req := this.Ctx.Request
			scheme := "http"
			if req.TLS != nil {
				scheme += "s"
			}
			redirect_to = fmt.Sprintf("%s://%s%s", scheme, req.Host, req.RequestURI)
		}
		redirect_to = "/login?to=" + url.QueryEscape(redirect_to)
		this.Redirect(redirect_to, code)
		return true
	}

	// if not need login then redirect
	if !needLogin && this.IsLogin {
		if len(redirect_to) == 0 {
			redirect_to = "/"
		}
		this.Redirect(redirect_to, code)
		return true
	}
	return false
}

// read beego flash message
func (this *BaseController) FlashRead(key string) (string, bool) {
	if data, ok := this.Data["flash"].(map[string]string); ok {
		value, ok := data[key]
		return value, ok
	}
	return "", false
}

// write beego flash message
func (this *BaseController) FlashWrite(key string, value string) {
	flash := beego.NewFlash()
	flash.Data[key] = value
	flash.Store(&this.Controller)
}

// check flash redirect, ensure browser redirect to uri and display flash message.
func (this *BaseController) CheckFlashRedirect(value string) (match bool, redirect bool) {
	v := this.GetSession("on_redirect")
	if params, ok := v.([]interface{}); ok {
		if len(params) != 5 {
			this.EndFlashRedirect()
			return match, redirect
		}
		uri := utils.ToStr(params[0])
		code := 302
		if c, ok := params[1].(int); ok {
			if c/100 == 3 {
				code = c
			}
		}
		flag := utils.ToStr(params[2])
		flagVal := utils.ToStr(params[3])
		times := 0
		if v, ok := params[4].(int); ok {
			times = v
		}

		times += 1
		if times > 3 {
			// if max retry times reached then end
			this.EndFlashRedirect()
			return match, redirect
		}

		// match uri or flash flag
		if uri == value || flag == value {
			match = true
		} else {
			// if no match then continue redirect
			this.FlashRedirect(uri, code, flag, flagVal, times)
			redirect = true
		}
	}
	return match, redirect
}

// set flash redirect
func (this *BaseController) FlashRedirect(uri string, code int, flag string, args ...interface{}) {
	flagVal := "true"
	times := 0
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			flagVal = v
		case int:
			times = v
		}
	}

	if len(uri) == 0 || uri[0] != '/' {
		panic("flash reirect only support same host redirect")
	}

	params := []interface{}{uri, code, flag, flagVal, times}
	this.SetSession("on_redirect", params)

	this.FlashWrite(flag, flagVal)
	this.Redirect(uri, code)
}

// clear flash redirect
func (this *BaseController) EndFlashRedirect() {
	this.DelSession("on_redirect")
}

// check form once, void re-submit
func (this *BaseController) FormOnceNotMatch() bool {
	notMatch := false
	recreat := false

	// get token from request param / header
	var value string
	if vus, ok := this.Input()["_once"]; ok && len(vus) > 0 {
		value = vus[0]
	} else {
		value = this.Ctx.Input.Header("X-Form-Once")
	}

	// exist in session
	if v, ok := this.GetSession("form_once").(string); ok && v != "" {
		// not match
		if value != v {
			notMatch = true
		} else {
			// if matched then re-creat once
			recreat = true
		}
	}

	this.FormOnceCreate(recreat)
	return notMatch
}

// create form once html
func (this *BaseController) FormOnceCreate(args ...bool) {
	var value string
	var creat bool
	creat = len(args) > 0 && args[0]
	if !creat {
		if v, ok := this.GetSession("form_once").(string); ok && v != "" {
			value = v
		} else {
			creat = true
		}
	}
	if creat {
		value = utils.GetRandomString(10)
		this.SetSession("form_once", value)
	}
	this.Data["once_token"] = value
	this.Data["once_html"] = template.HTML(`<input type="hidden" name="_once" value="` + value + `">`)
}

func (this *BaseController) validForm(form interface{}, names ...string) (bool, map[string]*validation.ValidationError) {
	// parse request params to form ptr struct
	utils.ParseForm(form, this.Input())

	// Put data back in case users input invalid data for any section.
	name := reflect.ValueOf(form).Elem().Type().Name()
	if len(names) > 0 {
		name = names[0]
	}
	this.Data[name] = form

	errName := name + "Error"

	// check form once
	if this.FormOnceNotMatch() {
		return false, nil
	}

	// Verify basic input.
	valid := validation.Validation{}
	if ok, _ := valid.Valid(form); !ok {
		errs := valid.ErrorMap()
		this.Data[errName] = &valid
		return false, errs
	}
	return true, nil
}

// valid form and put errors to tempalte context
func (this *BaseController) ValidForm(form interface{}, names ...string) bool {
	valid, _ := this.validForm(form, names...)
	return valid
}

// valid form and put errors to tempalte context
func (this *BaseController) ValidFormSets(form interface{}, names ...string) bool {
	valid, errs := this.validForm(form, names...)
	this.setFormSets(form, errs, names...)
	return valid
}

func (this *BaseController) SetFormSets(form interface{}, names ...string) *utils.FormSets {
	return this.setFormSets(form, nil, names...)
}

func (this *BaseController) setFormSets(form interface{}, errs map[string]*validation.ValidationError, names ...string) *utils.FormSets {
	formSets := utils.NewFormSets(form, errs, this.Locale)
	name := reflect.ValueOf(form).Elem().Type().Name()
	if len(names) > 0 {
		name = names[0]
	}
	name += "Sets"
	this.Data[name] = formSets

	return formSets
}

// add valid error to FormError
func (this *BaseController) SetFormError(form interface{}, fieldName, errMsg string, names ...string) {
	name := reflect.ValueOf(form).Elem().Type().Name()
	if len(names) > 0 {
		name = names[0]
	}
	errName := name + "Error"
	setsName := name + "Sets"

	if valid, ok := this.Data[errName].(*validation.Validation); ok {
		valid.SetError(fieldName, this.Tr(errMsg))
	}

	if fSets, ok := this.Data[setsName].(*utils.FormSets); ok {
		fSets.SetError(fieldName, errMsg)
	}
}

// check xsrf and show a friendly page
func (this *BaseController) CheckXsrfCookie() bool {
	return this.Controller.CheckXsrfCookie()
}

func (this *BaseController) SystemException() {

}

func (this *BaseController) IsAjax() bool {
	return this.Ctx.Input.Header("X-Requested-With") == "XMLHttpRequest"
}

func (this *BaseController) SetPaginator(per int, nums int64) *pagination.Paginator {
	p := pagination.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

func (this *BaseController) JsStorage(action, key string, values ...string) {
	value := action + ":::" + key
	if len(values) > 0 {
		value += ":::" + values[0]
	}
	this.Ctx.SetCookie("JsStorage", value, 1<<31-1, "/", nil, nil, false)
}

func (this *BaseController) setLangCookie(lang string) {
	this.Ctx.SetCookie("lang", lang, 60*60*24*365, "/", nil, nil, false)
}

//todo rewrite this part
// setLang sets site language version.
func (this *BaseController) setLang() bool {
	isNeedRedir := false
	//todo: add locale support, rewrite i18n.go

	hasCookie := false

	// get all lang names from i18n
	langs := setting.Langs

	// 1. Check URL arguments.
	lang := this.GetString("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. check if isLogin then use user setting
	if len(lang) == 0 && this.IsLogin {
		//todo:rewrite this part
		//lang = i18n.GetLangByIndex(this.User.Lang)
		lang = this.User.Lang
	}

	// 4. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := this.Ctx.Input.Header("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}
	// 4. DefaucurLang language is English.
	if len(lang) == 0 {
		lang = "en-US"
		isNeedRedir = false
	}

	// Save language information in cookies.
	if !hasCookie {
		this.setLangCookie(lang)
	}

	// Set language properties.
	this.Data["Lang"] = lang
	this.Data["Langs"] = langs
	this.Lang = lang

	return isNeedRedir
}

func (this *BaseController) CheckPermission(permissions ...string)  {
    if beego.RunMode =="dev" || len(permissions) == 0 {
        return
    }

	if this.User.Id <= 0{
		this.FlashRedirect("/login", 302, "not logged in")
		return
	}
    for _,p:= range permissions {
        if !this.UserService.HasPermission(&this.User, p) {
            this.Ctx.Abort(401, "not authorized!")
        }
    }
}
func (this *BaseController) AuthCheck(isRedirect bool) {
    if this.IsLogin && isRedirect {
        this.Redirect(beego.HttpAddr, 401)
    }
}

// get login redirect url from cookie
func (this *BaseController) GetLoginRedirect(ctx *context.Context) string {
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
func (this *BaseController) LoginUserRememberCookie(u *user.User, ctx *context.Context, remember bool) {
    // werid way of beego session regenerate id...
    ctx.Input.CruSession.SessionRelease(ctx.ResponseWriter)
    ctx.Input.CruSession = beego.GlobalSessions.SessionRegenerateId(ctx.ResponseWriter, ctx.Request)
    ctx.Input.CruSession.Set("auth_user_id", u.Id)

    if remember {
        this.WriteRememberCookie(u, ctx)
    }
}
func (this *BaseController) WriteRememberCookie(u *user.User, ctx *context.Context) {
    secret := utils.EncodeMd5(u.PasswordSalt + u.Password)
    days := 86400 * setting.LoginRememberDays
    ctx.SetCookie(setting.CookieUsername, u.Username, days)
    ctx.SetSecureCookie(secret, setting.CookieRememberName, u.Username, days)
}
func (this *BaseController) DeleteRememberCookie(ctx *context.Context) {
    ctx.SetCookie(setting.CookieUsername, "", -1)
    ctx.SetCookie(setting.CookieRememberName, "", -1)
}

func (this *BaseController) LoginUserFromRememberCookie(u *user.User, ctx *context.Context) (success bool) {
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

    this.LoginUserRememberCookie(u, ctx, true)

    return true
}

// logout user
func (this *BaseController) LogoutUser(ctx *context.Context) {
    this.DeleteRememberCookie(ctx)
    ctx.Input.CruSession.Delete("auth_user_id")
    ctx.Input.CruSession.Flush()
    beego.GlobalSessions.SessionDestroy(ctx.ResponseWriter, ctx.Request)
}

func (this *BaseController) GetUserIdFromSession(sess session.SessionStore) int64 {
    if id, ok := sess.Get("auth_user_id").(int64); ok && id > 0 {
        return id
    }
    return 0
}

// get user if key exist in session
func (this *BaseController) GetUserFromSession(u *user.User, sess session.SessionStore) bool {
    id := this.GetUserIdFromSession(sess)
    if id > 0 {
        u.Id = id
        if this.UserService.Read(u) == nil {
            return true
        }
    }

    return false
}
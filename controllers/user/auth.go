package user

import (
	"strings"

	"github.com/thanzen/eq/controllers/base"
	"github.com/thanzen/eq/models/user"
	"github.com/thanzen/eq/setting"
	"github.com/thanzen/eq/utils"
)

// AuthController serves login page.
type AuthController struct {
	base.BaseController
}

// Get implemented login page.
func (this *AuthController) Get() {
	this.Data["IsLoginPage"] = true
	this.TplNames = "auth/login.html"

	loginRedirect := strings.TrimSpace(this.GetString("to"))
	if utils.IsMatchHost(loginRedirect) == false {
		loginRedirect = "/"
	}

	// no need login
	if this.CheckLoginRedirect(false, loginRedirect) {
		return
	}

	if len(loginRedirect) > 0 {
		this.Ctx.SetCookie("login_to", loginRedirect, 0, "/")
	}

	form := LoginForm{}
	this.SetFormSets(&form)
}

// Login implemented user login.
func (this *AuthController) Login() {

	this.Data["IsLoginPage"] = true
	this.TplNames = "index.html"

	// no need login
	if this.CheckLoginRedirect(false) {
		return
	}

	var user user.User
	var key string
	ajaxErrMsg := "auth.login_error_ajax"

	form := LoginForm{}

	// valid form and put errors to template context
	if this.ValidFormSets(&form) == false {
		if this.IsAjax() {
			goto ajaxError
		}
		return
	}

	key = "auth.login." + form.Username + this.Ctx.Input.IP()
	if times, ok := utils.TimesReachedTest(key, setting.LoginMaxRetries); ok {
		if this.IsAjax() {
			ajaxErrMsg = "auth.login_error_times_reached"
			goto ajaxError
		}
		this.Data["ErrorReached"] = true

	} else if this.UserService.VerifyUser(&user, form.Username, form.Password) {
		loginRedirect := this.LoginUser(&user, form.Remember)
		if this.IsAjax() {
			this.Data["json"] = map[string]interface{}{
				"success":  true,
				"message":  this.Tr("auth.login_success_ajax"),
				"redirect": loginRedirect,
			}
			this.ServeJson()
			return
		}
		this.Redirect(loginRedirect, 302)
		return
	} else {
		utils.TimesReachedSet(key, times, setting.LoginFailedBlocks)
		if this.IsAjax() {
			goto ajaxError
		}
	}
	this.Data["Error"] = true
	return

ajaxError:
	this.Data["json"] = map[string]interface{}{
		"success": false,
		"message": this.Tr(ajaxErrMsg),
		"once":    this.Data["once_token"],
	}
	this.ServeJson()
}

// Logout implemented user logout page.
func (this *AuthController) Logout() {
	this.LogoutUser(this.Ctx)

	// write flash message
	this.FlashWrite("HasLogout", "true")

	this.Redirect("/login", 302)
}

package user

import (
	"github.com/astaxie/beego"
	"github.com/thanzen/eq/controllers/base"
	"github.com/thanzen/eq/models/user"
	"github.com/thanzen/eq/services/email"
	userServ "github.com/thanzen/eq/services/user"
)

// RegisterController serves register page.
type RegisterController struct {
	base.BaseController
}

// Get implemented Get method for RegisterController.
func (this *RegisterController) Get() {
	// no need login
	if this.CheckLoginRedirect(false) {
		return
	}

	this.Data["IsRegister"] = true
	this.TplNames = "auth/register.html"

	form := RegisterForm{Locale: this.Locale}
	this.SetFormSets(&form)
}

// Register implemented Post method for RegisterController.
func (this *RegisterController) Register() {
	this.Data["IsRegister"] = true
	this.TplNames = "auth/register.html"

	// no need login
	if this.CheckLoginRedirect(false) {
		return
	}

	form := RegisterForm{Locale: this.Locale}
	// valid form and put errors to template context
	if this.ValidFormSets(&form) == false {
		return
	}

	// Create new user.
	u := new(user.User)
	//set to default usertype
	ut := new(user.UserType)
	ut.Id = 1

	role := new(user.Role)
	role.Id = 1

	if err := this.UserService.RegisterUser(u, form.Username, form.Email, form.Password, ut, role); err == nil {
		email.SendRegisterMail(this.Locale, u)

		loginRedirect := this.LoginUser(u, false)
		if loginRedirect == "/" {
			this.FlashRedirect("/settings/profile", 302, "RegSuccess")
		} else {
			this.Redirect(loginRedirect, 302)
		}

		return

	} else {
		beego.Error("Register: Failed ", err)
	}
}

// Active implemented check Email actice code.
func (this *RegisterController) Active() {
	this.TplNames = "auth/active.html"

	// no need active
	if this.CheckActiveRedirect(false) {
		return
	}

	code := this.GetString(":code")

	var user user.User
	if this.UserService.VerifyUserActiveCode(&user, code) {
		user.Active = true
		user.PasswordSalt = userServ.GetUserSalt()
		if err := this.UserService.Update(&user, "Active", "PasswordSalt", "Updated"); err != nil {
			beego.Error("Active: user Update ", err)
		}
		if this.IsLogin {
			this.User = user
		}

		this.Redirect("/active/success", 302)

	} else {
		this.Data["Success"] = false
	}
}

// ActiveSuccess implemented success page when email active code verified.
func (this *RegisterController) ActiveSuccess() {
	this.TplNames = "auth/active.html"

	this.Data["Success"] = true
}

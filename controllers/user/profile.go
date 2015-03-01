package user

import (
	"github.com/astaxie/beego"
	"github.com/thanzen/eq/controllers/base"
	"github.com/thanzen/eq/services/email"
	"github.com/thanzen/eq/services/user"
)

// SettingsRouter serves user settings.
type ProfileController struct {
	base.BaseController
}

// Profile implemented user profile settings page.
func (this *ProfileController) Profile() {
	this.TplNames = "settings/profile.html"

	// need login
	if this.CheckLoginRedirect() {
		return
	}

	form := ProfileForm{Locale: this.Locale}
	form.SetFromUser(&this.User)
	this.SetFormSets(&form)

	formPwd := PasswordForm{}
	this.SetFormSets(&formPwd)
}

// ProfileSave implemented save user profile.
func (this *ProfileController) ProfileSave() {
	this.TplNames = "settings/profile.html"

	if this.CheckLoginRedirect() {
		return
	}

	action := this.GetString("action")

	if this.IsAjax() {
		switch action {
		case "send-verify-email":
			if this.User.Active {
				this.Data["json"] = false
			} else {
				email.SendActiveMail(this.Locale, &this.User)
				this.Data["json"] = true
			}

			this.ServeJson()
			return
		}
		return
	}

	profileForm := ProfileForm{Locale: this.Locale}
	profileForm.SetFromUser(&this.User)

	pwdForm := PasswordForm{User: &this.User}

	this.Data["Form"] = profileForm

	switch action {
	case "save-profile":
		if this.ValidFormSets(&profileForm) {
			if err := profileForm.SaveUserProfile(&this.User); err != nil {
				beego.Error("ProfileSave: save-profile", err)
			}
			this.FlashRedirect("/settings/profile", 302, "ProfileSave")
			return
		}

	case "change-password":
		if this.ValidFormSets(&pwdForm) {
			userServ := new(user.UserService)
			// verify success and save new password
			if err := userServ.SaveNewPassword(&this.User, pwdForm.Password); err == nil {
				this.FlashRedirect("/settings/profile", 302, "PasswordSave")
				return
			} else {
				beego.Error("ProfileSave: change-password", err)
			}
		}

	default:
		this.Redirect("/settings/profile", 302)
		return
	}

	if action != "save-profile" {
		this.SetFormSets(&profileForm)
	}
	if action != "change-password" {
		this.SetFormSets(&pwdForm)
	}
}

package user

import (
	"github.com/astaxie/beego"
)

func RegisterRoutes() {
	actr := new(AuthController)
	beego.Router("/login", actr, "get:Get;post:Login")
	beego.Router("/logout", actr, "get:Logout")

	register := new(RegisterController)
	beego.Router("/register", register, "get:Get;post:Register")
	beego.Router("/active/success", register, "get:ActiveSuccess")
	beego.Router("/active/:code([0-9a-zA-Z]+)", register, "get:Active")

    forgot := new(ForgotController)
    beego.Router("/forgot", forgot)
    beego.Router("/reset/:code([0-9a-zA-Z]+)", forgot, "get:Reset;post:ResetPost")

	settings := new(ProfileController)
	beego.Router("/settings/profile", settings, "get:Profile;post:ProfileSave")

	user := new(UserController)
	beego.Router("/user/:username", user, "get:Index")

	admin := new(AdminController)
	beego.Router("/admin", admin, "get:Index")
}

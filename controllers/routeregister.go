package controllers

import (
	"github.com/astaxie/beego"
	userApi "github.com/thanzen/eq/controllers/api/user"
	"github.com/thanzen/eq/controllers/user"
	"github.com/thanzen/eq/setting"
)

func RegisterControllers() {
	//register maincontroller
	mc := new(MainController)
	beego.Router("/", mc, "get:Index")
	beego.Router("/index", mc, "get:Index")

	// "robot.txt"
	beego.Router("/robot.txt", &RobotRouter{})

	// Add Filters
	//todo:enable image filter later
	//  beego.InsertFilter("/img/*", beego.BeforeRouter, attachment.ImageFilter)

	beego.InsertFilter("/captcha/*", beego.BeforeRouter, setting.Captcha.Handler)
	//register user related controllers
	user.RegisterRoutes()
	userApi.RegisterRoutes()
}

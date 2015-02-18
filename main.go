package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/thanzen/eq/controllers"
	"github.com/thanzen/eq/services"
	"github.com/thanzen/eq/setting"
)

func initialize() {

}
func main() {
	//todo:move to init if we can
	setting.Initialize()
	beego.SetLogFuncCall(true)
	beego.Info("AppPath:", beego.AppPath)
	beego.SetViewsPath("views")

	if setting.IsProMode {
		beego.Info("Product mode enabled")
	} else {
		beego.Info("Develment mode enabled")
	}
	beego.Info(beego.AppName, setting.APP_VER, setting.AppUrl)
	if !setting.IsProMode {
		beego.SetStaticPath("/static_source", "static_source")
		beego.DirectoryIndex = true
	}
	orm.RegisterDriver("postgres", orm.DR_Postgres)

	orm.RegisterDataBase("default", "postgres", "user=postgres password=root dbname=eq sslmode=disable")
	services.Register()
	orm.RunCommand()
	orm.Debug = true

	controllers.RegisterControllers()


	//beego.Router("/user", uctr, "get:GetOne")
	//beego.BeeLogger.DelLogger("console")
	//beego.SetLevel(beego.LevelInformational)
	if beego.RunMode == "dev" {
		//	beego.Router("/test/:tmpl(mail/.*)", new(base.TestRouter))
	}
	beego.Run()
}

package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/thanzen/eq/controllers"
	"github.com/thanzen/eq/services"
	"github.com/thanzen/eq/setting"
	"github.com/thanzen/migrate/migrate"
)

func main() {
	setting.Initialize()
	beego.SetLogFuncCall(true)
	beego.SetViewsPath("views")

	if setting.IsProMode {
		beego.Info("Product mode enabled")
        beego.Info(setting.PostgresMigrateConnection)
        //auto migrate db
        //todo: we may want to disable this later
        dbMigrate()
	} else {
		beego.Info("Develment mode enabled")
	}
	beego.Info(beego.AppName, setting.APP_VER, setting.AppUrl)
	if !setting.IsProMode {

		beego.SetStaticPath("/static", "static")
		beego.DirectoryIndex = true
	}
	orm.RegisterDriver("postgres", orm.DR_Postgres)

	orm.RegisterDataBase("default", "postgres", setting.PostgresConnection)
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
func dbMigrate() {
	allErrors, ok := migrate.UpSync(setting.PostgresMigrateConnection, "./conf/db/migrations")
	if !ok {
		beego.Error(allErrors)
	}

}

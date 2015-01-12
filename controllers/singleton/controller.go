package singleton

import (
	"github.com/gin-gonic/gin"
	"github.com/thanzen/eq/controllers"
	"github.com/thanzen/eq/services/singleton"
)

var UserController *controllers.UserController
var UserTypeController *controllers.UserTypeController

func RegisterControllers(engine *gin.Engine, group ...*gin.RouterGroup) {
	UserController := &controllers.UserController{UserService: singleton.Users}
	UserController.Register(engine, group...)
	UserTypeController := &controllers.UserTypeController{Service: singleton.UserTypes}
	UserTypeController.Register(engine, group...)

}

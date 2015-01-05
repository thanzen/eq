package singleton

import (
	"github.com/gin-gonic/gin"
	"github.com/thanzen/eq/controllers"
	"github.com/thanzen/eq/services/singleton"
)

var UserController *controllers.UserController

func RegisterControllers(engine *gin.Engine, group ...*gin.RouterGroup) {
	UserController := &controllers.UserController{UserService: singleton.Users}
	UserController.Register(engine, group...)
}

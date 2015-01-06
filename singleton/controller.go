package singleton

import (
	"github.com/thanzen/eq/controllers"

	"github.com/gin-gonic/gin"
)

var UserController *controllers.UserController

func RegisterControllers(engine *gin.Engine, group ...*gin.RouterGroup) {
	UserController := &controllers.UserController{UserService: Users}
	UserController.Register(engine, group...)
}

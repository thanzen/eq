package web

import (
	"github.com/gin-gonic/gin"
)

//Defines a interface for controllers.
type Controller interface {
	Register(engine *gin.Engine, group ...*gin.RouterGroup)
}

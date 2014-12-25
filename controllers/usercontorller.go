package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanzen/eq/models/user"
	"github.com/thanzen/eq/services/userservice"
)

type UserController struct {
	UserService *userservice.UserService
}

func (ct *UserController) Register(engine *gin.Engine, group ...*gin.RouterGroup) {
	engine.GET("user/:id", ct.get)
	engine.GET("users", ct.getall)
}
func (ct UserController) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.Error(err, "Invalid id")
		c.Abort(500)
	} else {
		var u *user.User
		ct.UserService.Get(u, id)
		c.JSON(200, u)
	}

}
func (ct UserController) getall(c *gin.Context) {

	var users []*user.User
	ct.UserService.GetList(&users, nil, 1, 5)

	c.JSON(200, users)

}

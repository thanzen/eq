package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thanzen/eq/models/user"
	"github.com/thanzen/eq/services"
	"github.com/thanzen/eq/services/userservice"
	//"log"
	"strconv"
)

type UserController struct {
	UserService *userservice.UserService
}

func (ct *UserController) Register(engine *gin.Engine, group ...*gin.RouterGroup) {
	//log.Println("register called")
	engine.GET("user/:id", ct.get)
	engine.GET("/users", ct.getall)
	engine.POST("/login", ct.login)
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
func (ct UserController) login(c *gin.Context) {
	// Example for binding JSON ({"user": "manu", "password": "123"})
	var json, u user.LoginAccount
	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	conds := services.SearchOptions{"user_name": "lnelson0"}
	ct.UserService.Login(&u, conds)
	//todo: add real authentication logic
	if u.Username == json.Username {
		c.JSON(200, gin.H{"status": "you are logged in"})
	} else {
		c.JSON(401, gin.H{"status": "unauthorized"})
	}
}

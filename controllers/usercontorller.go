package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thanzen/eq/models/user"
	"github.com/thanzen/eq/services"
	"github.com/thanzen/eq/services/userservice"
	"log"
	"strconv"
)

type UserController struct {
	UserService *userservice.UserService
}

func (uc *UserController) Register(engine *gin.Engine, group ...*gin.RouterGroup) {
	log.Println("register called")
	engine.GET("user/:id", uc.get)
	engine.GET("/users", uc.getall)
	engine.POST("/login", uc.login)
	engine.POST("/register", uc.register)
	engine.GET("/count", uc.count)
}
func (uc UserController) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.Error(err, "Invalid id")
		c.Abort(500)
	} else {
		var u *user.User
		uc.UserService.Get(u, id)
		c.JSON(200, u)
	}
}

func (uc UserController) getall(c *gin.Context) {

	var users []*user.User
	conds := services.SearchOptions{"deleted": false, "active": true}
	uc.UserService.GetList(&users, conds, 1, 5)

	c.JSON(200, users)
}
func (uc UserController) count(c *gin.Context) {

	var u user.User
	conds := services.SearchOptions{"deleted": false, "active": true}
	n, err := uc.UserService.Count(&u, conds)
	if err != nil {
		//todo process error
	}

	c.JSON(200, n)
}
func (uc UserController) login(c *gin.Context) {
	// Example for binding JSON ({"user": "manu", "password": "123"})
	var json, u user.LoginAccount
	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	conds := services.SearchOptions{"user_name": "lnelson0", "deleted": false, "active": true}
	uc.UserService.Login(&u, conds)
	//todo: add real authentication logic
	if u.Username == json.Username {
		c.JSON(200, gin.H{"status" + u.Username: "you are logged in" + json.Username})
	} else {
		c.JSON(401, gin.H{"status  " + u.Username: "unauthorized  " + json.Username})
	}
}

func (uc UserController) register(c *gin.Context) {
	//todo: more validations required
	var json user.LoginAccount
	c.Bind(&json)

	if json.Username == "" {
		c.JSON(500, gin.H{"error :": "username cannot be empty!"})
	}
	err := uc.UserService.CreateAccount(&json)
	if err != nil {
		c.JSON(500, gin.H{"error :": "cannot create user!"})
	}
}

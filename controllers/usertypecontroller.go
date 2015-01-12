package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thanzen/eq/models/user"
	//	"github.com/thanzen/eq/services"
	"github.com/thanzen/eq/services/userservice"
	"strconv"
)

type UserTypeController struct {
	Service *userservice.UserTypeService
}

func (uc *UserTypeController) Register(engine *gin.Engine, group ...*gin.RouterGroup) {
	engine.GET("usertype/:id", uc.get)
	engine.GET("/usertypes", uc.getall)
	engine.POST("/usertype", uc.insert)

}
func (uc UserTypeController) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.Error(err, "Invalid id")
		c.Abort(500)
	} else {
		var u user.UserType
		uc.Service.Get(&u, id)
		c.JSON(200, u)
	}
}
func (uc UserTypeController) getall(c *gin.Context) {

	var uts []*user.UserType
	uc.Service.GetList(&uts, nil)

	c.JSON(200, uts)
}

func (uc UserTypeController) insert(c *gin.Context) {
	//todo: more validations required
	var json user.UserType
	c.Bind(&json)

	if json.Name == "" {
		c.JSON(500, gin.H{"error :": "username cannot be empty!"})
	}
	err := uc.Service.Save(&json)
	if err != nil {
		c.JSON(500, gin.H{"error :": "cannot create user!"})
	}
}

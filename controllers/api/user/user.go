package user

import (
	"github.com/thanzen/eq/conf/permissions"
	"github.com/thanzen/eq/controllers/base"
	"github.com/thanzen/eq/models/user"
    "github.com/astaxie/beego"
)

type UserApiController struct {
	base.BaseController
}

func (this *UserApiController) GetUsers() {
	this.CheckPermission(permissions.UserFuzzSearch)
	query := this.Ctx.Input.Param(":query")
	var offset, limit int64
	var err error
	offset, err = this.GetInt64("offset")
	if err != nil || offset < 0 {
		this.Ctx.Abort(500, "invalid offset")
	}
	limit, err = this.GetInt64(":limit")
	if err != nil || limit < 1 {
        beego.Info(limit)
        beego.Info(err)
        this.Ctx.Abort(500, "invalid limit")
    }
	var users []*user.User
	this.UserService.FuzzySearch(&users, query, offset, limit)
	this.Data["json"] = &users
	this.ServeJson()
}

package user


//todo: refactor error code for http request
//todo: swagger support
import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/thanzen/eq/conf/permissions"
	"github.com/thanzen/eq/controllers/base"
	"github.com/thanzen/eq/models/user"
)

type AdminApiController struct {
	base.BaseController
}

func (this *AdminApiController) GetUsers() {
	this.CheckPermission(permissions.UserFuzzSearch)
	query := this.Ctx.Input.Param(":query")
	var roleId, offset, limit int64
	var err error
	roleId, err = this.GetInt64(":roleId")
	if err != nil || offset < 0 {
		this.Ctx.Abort(500, "invalid role id")
	}
	offset, err = this.GetInt64("offset")
	if err != nil || offset < 0 {
		this.Ctx.Abort(500, "invalid offset")
	}
	limit, err = this.GetInt64(":limit")
	if err != nil || limit < 1 {
		this.Ctx.Abort(500, "invalid limit")
	}
	var users []*user.User
	this.UserService.FuzzySearch(&users, query, roleId, offset, limit)
	this.Data["json"] = users
	this.ServeJson(true)
}
func (this *AdminApiController) Update() {
	this.CheckPermission(permissions.UserAdminUpdate)
	var u user.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &u)
	beego.Info(u)
	var err error
	if u.Id > 0 {
		err = this.UserService.Update(&u, "Firstname", "Lastname", "Active", "Email", "Company")
		if err != nil {
			this.Ctx.Abort(500, err.Error())
		}
	} else {
		this.Ctx.Abort(500, "invalid user id")
	}
	this.ServeJson()
}
func (this *AdminApiController) ResetPassword(){
    var ChangePasswordModel = struct {
        Id int64
        Password string
    }{}
    this.CheckPermission(permissions.UserAdminUpdate)
    var u user.User
    json.Unmarshal(this.Ctx.Input.RequestBody, &ChangePasswordModel)
    u.Id = ChangePasswordModel.Id
    u.Password = ChangePasswordModel.Password
    var err error
    if u.Id > 0 {
        err = this.UserService.SaveNewPassword(&u,u.Password)
        if err != nil {
            this.Ctx.Abort(500, err.Error())
        }
    } else {
        this.Ctx.Abort(500, "invalid user id")
    }
    this.ServeJson()
}

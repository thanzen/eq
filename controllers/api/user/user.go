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

type UserFuzzSearchParam struct{
	Query string `json:"query"`
	RoleId int64 `json:"roleId"`
	Offset int64 `json:"offset"`
	Limit int64 `json:"limit"`
	IncludeTotal bool `json:"includeTotal"`
}

type UserFuzzSearchResponse struct{
	Users []*user.User `json:"users"`
	Total int64 `json:"total"`
}

func (this *AdminApiController) GetUsers() {
	this.CheckPermission(permissions.UserFuzzSearch)
	var param UserFuzzSearchParam
	json.Unmarshal(this.Ctx.Input.RequestBody, &param)
	var err error
	beego.Info(param)
	if err != nil || param.RoleId < 0 {
		this.Ctx.Abort(500, "invalid role id")
	}
	if err != nil || param.Offset < 0 {
		this.Ctx.Abort(500, "invalid offset")
	}
	if err != nil || param.Limit < 0 {
		this.Ctx.Abort(500, "invalid limit")
	}
	var users []*user.User
	n,_ := this.UserService.FuzzySearch(&users, param.Query, param.RoleId, param.Offset, param.Limit,param.IncludeTotal)
	resp := UserFuzzSearchResponse{Users:users,Total:n}
	this.Data["json"] = resp
	this.ServeJson(true)
}


func (this *AdminApiController) Update() {
	this.CheckPermission(permissions.UserAdminUpdate)
	var u user.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &u)
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

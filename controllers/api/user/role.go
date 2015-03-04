package user

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/thanzen/eq/conf/permissions"
	"github.com/thanzen/eq/controllers/base"
	"github.com/thanzen/eq/models/user"
	userServes "github.com/thanzen/eq/services/user"
	"strings"
)

type roleUserModel struct {
	roleId int64
	users  []*user.User
}

type rolePermissionModel struct {
	roleId      int64
	permissions []*user.Permission
}

type RoleApiController struct {
	base.BaseController
	roleService userServes.RoleService
}

func NewRoleApiController() *RoleApiController {
	ct := &RoleApiController{roleService: userServes.RoleService{}}
	return ct
}

func (this *RoleApiController) GetRoles() {
	this.CheckPermission(permissions.RoleViewAll)
	cachedRoles := this.roleService.GetAllRoles()
	var roles []*user.Role = make([]*user.Role, 0)
	for _, r := range cachedRoles {
		roles = append(roles, r)
	}
	this.Data["json"] = roles
	this.ServeJson()
}

//AddUsers add users to given role
func (this *RoleApiController) AddUsers() {
	this.CheckPermission(permissions.RoleCreate)
	var model roleUserModel
	json.Unmarshal(this.Ctx.Input.RequestBody, &model)
	if model.roleId > 0 && len(model.users) > 0 {
		r := &user.Role{Id: model.roleId}
		err := this.roleService.InsertUsers(r, model.users...)
		if err != nil {
			this.Ctx.Abort(500, err.Error())
		}
	} else if model.roleId <= 0 {
		this.Ctx.Abort(500, "invalid role id")
	}
	this.ServeJson()
}

//DeleteUsers delete users for given role
func (this *RoleApiController) DeleteUsers() {
	this.CheckPermission(permissions.RoleDeleteUsers)
	var model roleUserModel
	json.Unmarshal(this.Ctx.Input.RequestBody, &model)
	if model.roleId > 0 && len(model.users) > 0 {
		r := &user.Role{Id: model.roleId}
		err := this.roleService.DeleteUsers(r, model.users...)
		if err != nil {
			this.Ctx.Abort(500, err.Error())
		}
	} else if model.roleId <= 0 {
		this.Ctx.Abort(500, "invalid role id")
	}
	this.ServeJson()
}

//AddPermissions add permissions to given role
func (this *RoleApiController) AddPermissions() {
	this.CheckPermission(permissions.RoleAddPermissions)
	var model rolePermissionModel
	json.Unmarshal(this.Ctx.Input.RequestBody, &model)
	if model.roleId > 0 && len(model.permissions) > 0 {
		r := &user.Role{Id: model.roleId}
		err := this.roleService.InsertPermissions(r, model.permissions...)
		if err != nil {
			this.Ctx.Abort(500, err.Error())
		}
	} else if model.roleId <= 0 {
		this.Ctx.Abort(500, "invalid role id")
	}
	this.ServeJson()
}

//DeletePermissions delete permissions for given role
func (this *RoleApiController) DeletePermissions() {
	this.CheckPermission(permissions.RoleDeletePermissions)
	var model rolePermissionModel
	json.Unmarshal(this.Ctx.Input.RequestBody, &model)
	if model.roleId > 0 && len(model.permissions) > 0 {
		r := &user.Role{Id: model.roleId}
		err := this.roleService.DeletePermissions(r, model.permissions...)
		if err != nil {
			this.Ctx.Abort(500, err.Error())
		}
	} else if model.roleId <= 0 {
		this.Ctx.Abort(500, "invalid role id")
	}
	this.ServeJson()
}

//Add add a role
func (this *RoleApiController) Add() {
	this.CheckPermission(permissions.RoleCreate)
	var model user.Role
	json.Unmarshal(this.Ctx.Input.RequestBody, &model)
	model.Name = strings.Trim(model.Name, " ")
	this.validate(&model)
	err := this.roleService.Insert(&model)
	if err != nil {
		this.Ctx.Abort(500, err.Error())
	}
	this.Data["json"] = model
	this.ServeJson()
}

//Update update a role
func (this *RoleApiController) Update() {
	this.CheckPermission(permissions.RoleUpdate)
	var model user.Role
	json.Unmarshal(this.Ctx.Input.RequestBody, &model)
	model.Name = strings.Trim(model.Name, " ")
	if model.Id <= 0 {
		this.Ctx.Abort(500, "invalid role id")
	}
	this.validate(&model)
	err := this.roleService.Update(&model, "Name", "IsSystemRole", "Description")
	if err != nil {
		this.Ctx.Abort(500, err.Error())
	}
	this.ServeJson()
}

func (this *RoleApiController) validate(role *user.Role) {
	valid := validation.Validation{}
	b, _ := valid.Valid(role)
	if !b {
		m := ""
		for _, err := range valid.Errors {
			m += fmt.Sprintf("%s:%s\n", err.Key, err.Message)
		}
		this.Ctx.Abort(500, m)
	}
}


package user

import (
    "github.com/thanzen/eq/controllers/base"
    userServes "github.com/thanzen/eq/services/user"
    "github.com/thanzen/eq/models/user"
    "encoding/json"
)

type roleUserModel struct {
    roleId int64
    users []*user.User
}
type rolePermissionModel struct {
    roleId int64
    permissions []*user.Permission
}

type RoleApiController struct {
    base.BaseRouter
    roleService userServes.RoleService
}

func NewRoleApiController() *RoleApiController {
    ct := &RoleApiController{roleService: userServes.RoleService{}}
    return ct
}

func (this *RoleApiController) GetRoles() {
    roles := this.roleService.GetAllRoles()
    this.Data["json"] = roles
    this.ServeJson()
}

//todo: permission key should be in a
func (this *RoleApiController) AddUsers() {
    if !this.UserService.HasPermission(this.User, "eq_role_adduserstorole") {
        this.Ctx.Abort("401", "not authorized!")
        return
    }
    var model roleUserModel
    json.Unmarshal(this.Ctx.Input.RequestBody, &model)
    if model.roleId>0 && len(model.users)>0 {
        this.roleService.InsertUsers(model.users)
    }
}

func (this *RoleApiController) AddPermissions() {
    if !this.UserService.HasPermission(this.User, "eq_role_addpermissionstorole") {
        this.Ctx.Abort("401", "not authorized!")
        return
    }
    var model rolePermissionModel
    json.Unmarshal(this.Ctx.Input.RequestBody, &model)
    if model.roleId>0 && len(model.permissions)>0 {
        this.roleService.InsertUsers(model.permissions)
    }
}
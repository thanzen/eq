package user

import (
    "github.com/thanzen/eq/controllers/base"
//  "github.com/thanzen/eq/models/user"
    userServes "github.com/thanzen/eq/services/user"
    "github.com/thanzen/eq/models/user"
    "encoding/json"
)

type roleUserModel struct {
    roleId int64
    users []*user.User
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
func (this *RoleApiController) AddUsers() {
    var model roleUserModel
    json.Unmarshal(this.Ctx.Input.RequestBody, model)

    this.roleService.InsertUsers(model.users)

}


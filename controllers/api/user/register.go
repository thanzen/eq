package user
import (
    "github.com/astaxie/beego"
)
func RegisterRoutes() {
    role := NewRoleApiController()
    beego.Router("/api/role", role, "get:GetRoles;post:Add;put:Update;delete:Delete")
    uct :=&AdminApiController{}
    beego.Router("/api/admin/user/getusers",uct,"post:GetUsers")
    beego.Router("/api/admin/user/update",uct,"post:Update")
    beego.Router("/api/admin/user/reset",uct,"post:ResetPassword")
}
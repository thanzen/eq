package user
import (
    "github.com/astaxie/beego"
)
func RegisterRoutes() {
    role := NewRoleApiController()
    beego.Router("/api/role", role, "get:GetRoles;post:Add;put:Update;delete:Delete")
    uct :=&AdminApiController{}
    beego.Router("/api/admin/user/query/?:query/role/?:roleId([0-9]+)/offset/?:offset([0-9]+)/limit/?:limit([1-9][0-9]*)", uct, "get:GetUsers")
    beego.Router("/api/admin/user/update",uct,"post:Update")
    beego.Router("/api/admin/user/reset",uct,"post:ResetPassword")
}
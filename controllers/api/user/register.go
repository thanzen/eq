package user
import (
    "github.com/astaxie/beego"
)
func RegisterRoutes() {

    role := NewRoleApiController()
    beego.Router("/api/role", role, "get:GetRoles;post:Add;put:Update;delete:Delete")

//    beego.Router("/api/role/:id", role, "get")
}
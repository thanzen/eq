package user
import (
    "github.com/astaxie/beego"
)
func RegisterRoutes() {

    role := NewRoleApiController()
    beego.Router("/api/role", role, "get:GetRoles;post:Add;put:Update;delete:Delete")

    uct :=&UserApiController{}
    beego.Router("/api/user/list/:query/:offset([0-9]+)/:limit([1-9][0-9]*)", uct, "get:GetUsers")
//    beego.Router("/api/role/:id", role, "get")
}
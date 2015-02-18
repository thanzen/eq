package services

import (
    "github.com/astaxie/beego/orm"
    "github.com/thanzen/eq/models/user"
)
func Register(){
    orm.RegisterModel(new(user.UserType), new(user.User),new(user.Role),new(user.Permission))
}
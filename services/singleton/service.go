package singleton

import (
	"github.com/thanzen/eq/services"
	"github.com/thanzen/eq/services/userservice"
	"github.com/thanzen/modl"
)

var Users *userservice.UserService
var UserTypes *userservice.UserTypeService

func RegisterServices(m *modl.DbMap) {
	Users = userservice.Create(m)
	UserTypes = &userservice.UserTypeService{services.DefaultRepository{Modl: m}}
}

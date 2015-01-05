package singleton

import (
	"github.com/thanzen/eq/services/userservice"
	"github.com/thanzen/modl"
)

var Users *userservice.UserService

func RegisterServices(m *modl.DbMap) {
	Users = userservice.Create(m)
}

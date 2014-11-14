package singleton

import (
	"github.com/thanzen/eq/services"
	"github.com/thanzen/eq/services/userservice"
)

var Users *userservice.UserService

func RegisterServices(dbcontext *services.DbContext) {
	Users = &userservice.UserService{DbContext: dbcontext}
}

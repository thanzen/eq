package userservice

import (
	"github.com/thanzen/eq/services"

	"github.com/thanzen/eq/models/user"
)

type UserService struct {
	*services.DbContext
}

func CreateUserService(dbcontext *services.DbContext) UserService {
	c := UserService{}
	return c
}

func (serv *UserService) GetById(id int) *user.User {
	return &user.User{}
	//result, err := serv.Gorp.Get(user.User{}, id)
	//if err != nil {
	//	return nil
	//}
	//return result.(*user.User)
}

func (serv *UserService) Insert(model user.User) *user.User {
	err := serv.Gorp.Insert(model)
	if err != nil {
		return nil
	}
	return &model
}

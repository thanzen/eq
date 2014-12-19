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
	if id <= 0 {
		return nil
	}
	var u *user.User
	err := serv.Modl.Get(u, id)
	if err != nil {
		return nil
	}
	return u
}

func (serv *UserService) Insert(model user.User) *user.User {
	err := serv.Modl.Insert(model)
	if err != nil {
		return nil
	}
	return &model
}

func (serv *UserService) GetList(options services.SearchOptions) []user.User {

	query := "select * from user_meta"

	// pass a slice to Select()
	var list []user.User
	err := serv.Modl.Select(&list, query)
	if err != nil {
		panic(err)
	}
	return list
}

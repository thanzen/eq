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
	result, err := serv.Gorp.Get(user.User{}, id)
	if err != nil {
		return nil
	}
	return result.(*user.User)
}

func (serv *UserService) Insert(model user.User) *user.User {
	err := serv.Gorp.Insert(model)
	if err != nil {
		return nil
	}
	return &model
}

func (serv *UserService) GetList(options services.SearchOptions) []user.User {

	query := "select * from user_meta"

	// pass a slice to Select()
	var list []user.User
	_, err := serv.Gorp.Select(&list, query)
	if err != nil {
		panic(err)
	}
	return list
}

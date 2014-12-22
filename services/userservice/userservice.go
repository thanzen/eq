package userservice

import (
	//"errors"
	"github.com/thanzen/eq/models/user"
	"github.com/thanzen/eq/services"
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

//Save provides Insert and Update for user.User.
//When u(user) is nil, it performs insert, otherwise, it performs update.
//Todo: added error log
func (serv *UserService) Save(u *user.User) *user.User {
	if u == nil {
		return nil
	}
	var err error
	if u.Id <= 0 {
		err = serv.Modl.Insert(u)
	} else {
		_, err = serv.Modl.Update(u)
	}
	if err != nil {
		return nil
	}
	return u
}

func (serv *UserService) GetList() []*user.User {

	//query := "select * from user_meta"

	// pass a slice to Select()
	list := []*user.User{}
	dfservice := &services.DefaultRepository{Modl: serv.DbContext.Modl}
	err := dfservice.GetList(&list, nil)
	//err := serv.Modl.Select(&list, query)
	if err != nil {
		panic(err)
	}
	return list
}

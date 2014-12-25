package userservice

import (
	//"errors"
	"github.com/thanzen/eq/models/user"
	"github.com/thanzen/eq/services"
	"github.com/thanzen/modl"
)

type UserService struct {
	repo services.Repositoryer
}

func Create(modl *modl.DbMap) *UserService {
	c := &UserService{}
	c.repo = &services.DefaultRepository{Modl: modl}
	return c
}

func (serv *UserService) Get(u *user.User, id int) error {
	return serv.repo.Get(u, id)
}

func (serv *UserService) Save(u *user.User) error {
	return serv.repo.Save(u)
}

func (serv *UserService) GetList(list *[]*user.User, options services.SearchOptions, pos ...int) error {
	return serv.repo.GetList(list, options, pos...)
}

package userservice

import (
	//"errors"
	"github.com/thanzen/eq/models/user"
	"github.com/thanzen/eq/services"
	"github.com/thanzen/modl"
	"log"
)

type UserService struct {
	repo      services.Repositoryer
	loginRepo services.Repositoryer
}

func Create(modl *modl.DbMap) *UserService {
	c := &UserService{}
	c.repo = &services.DefaultRepository{Modl: modl}
	c.loginRepo = &services.DefaultRepository{Modl: modl}
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

func (serv *UserService) Login(u *user.LoginAccount, options services.SearchOptions) error {
	var list []*user.LoginAccount
	err := serv.loginRepo.GetList(&list, options)
	log.Println("list1:", list)
	if len(list) == 1 {
		u = list[0]
	}
	log.Println("list2:", u)
	return err
}

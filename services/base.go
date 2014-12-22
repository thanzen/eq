package services

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/thanzen/modl"
	"reflect"
)

type SearchOptions map[string]interface{}

//Database related context
type DbContext struct {
	Modl *modl.DbMap
}

//Repositoryer represent interface for repository
//Repositoryer should be good enough for most basic CRUD
//If the service needs more logic or functionality, it can
//composit this interface with its custom functions.
type Repositoryer interface {
	Get(id int) interface{}
	Save(model interface{}) error
	GetList(options SearchOptions) []interface{}
	Delete(id int) error
}

//DefaultRepository provides basic implementation of Repositoryer
//It assumes that the model must have Id as its primary key
type DefaultRepository struct {
	Modl      *modl.DbMap
	selectAll string
}

//Get returns the entity by given keys.
//The sequence of the keys follows the sequence of the properties in the dest class
func (repo *DefaultRepository) Get(dest interface{}, keys ...interface{}) error {
	if len(keys) <= 0 {
		return errors.New("Keys can not be empty")
	}
	err := repo.Modl.Get(dest, keys...)
	return err
}

//Save provides Insert and Update for user.User.
//When u(user) is nil, it performs insert, otherwise, it performs update.
//Todo: added error log
func (repo *DefaultRepository) Save(dest interface{}) error {
	if dest == nil {
		return errors.New("Keys can not be empty")
	}
	var err error
	v := reflect.ValueOf(dest)
	//reflect all the keys and remove the id limitation
	if v.FieldByName("Id").Interface().(int) <= 0 {
		err = repo.Modl.Insert(dest)
	} else {
		_, err = repo.Modl.Update(dest)
	}
	return err
}

func (repo *DefaultRepository) GetList(dest interface{}, options SearchOptions) error {
	sql := repo.GenerateSelectSql(dest, options)
	if sql == "" {
		return errors.New("Generate sql error")
	}
	err := repo.Modl.Select(dest, sql)
	return err
}

//GenerateSelectSql generate select template for given search options
//note: in order to avoid sql injection, GenerateSelectSql function skip to
//fill the search option values in, instead, use ? as parameters so that
//necessary validation will be performed  by database/sql package.
func (repo *DefaultRepository) GenerateSelectSql(dest interface{}, options SearchOptions) string {
	table := repo.Modl.TableFor(dest)
	if table == nil || len(table.Columns) < 1 {
		return ""
	}
	sql := repo.getSelectAll(table)
	s := bytes.Buffer{}
	s.WriteString(" where ")
	x := 0
	for key, _ := range options {
		if x > 0 {
			s.WriteString(" and ")
		}
		s.WriteString(repo.Modl.Dialect.QuoteField(key))
		s.WriteString("=")
		s.WriteString(repo.Modl.Dialect.BindVar(x))
		x++
	}
	if x > 0 {
		sql += s.String()
	}
	return sql
}
func (repo *DefaultRepository) getSelectAll(table *modl.TableMap) string {
	if repo.selectAll == "" {
		s := bytes.Buffer{}
		s.WriteString("select ")
		x := 0
		for _, col := range table.Columns {
			if !col.Transient {
				if x > 0 {
					s.WriteString(",")
				}
				s.WriteString(repo.Modl.Dialect.QuoteField(col.ColumnName))
				x++
			}
		}
		s.WriteString(" from ")
		s.WriteString(repo.Modl.Dialect.QuoteField(table.TableName))
		repo.selectAll = s.String()
	}
	return repo.selectAll
}

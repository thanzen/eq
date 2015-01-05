package services

import (
	"bytes"
	"errors"
	//	"github.com/thanzen/eq/util"
	"github.com/thanzen/modl"
	//	"log"
	"reflect"
	"strconv"
)

type SearchOptions map[string]interface{}

//Repositoryer represent interface for repository
//Repositoryer should be good enough for most basic CRUD
//If the service needs more logic or functionality, it can
//composit this interface with its custom functions.
type Repositoryer interface {
	Get(dest interface{}, keys ...interface{}) error
	Save(dest interface{}) error
	Insert(dest interface{}) error
	Update(dest interface{}) error
	GetList(dest interface{}, options SearchOptions, pos ...int) error
	Count(dest interface{}, options SearchOptions) (int, error)
	Delete(dest interface{}) error
	First(dest interface{}, options SearchOptions) error
}

//make sure DefaultRepository implements interface Repositoryer
var _ Repositoryer = &DefaultRepository{nil, ""}

//DefaultRepository provides basic implementation of Repositoryer
type DefaultRepository struct {
	Modl      *modl.DbMap
	selectAll string
}

//Get returns the entity by given keys.
//The sequence of the keys follows the sequence of the properties in the dest class
func (repo *DefaultRepository) Get(dest interface{}, keys ...interface{}) error {
	if len(keys) <= 0 {
		return errors.New("keys can not be empty")
	}
	return repo.Modl.Get(dest, keys...)
}

//Get returns the entity by given keys.
//The sequence of the keys follows the sequence of the properties in the dest class
func (repo *DefaultRepository) First(dest interface{}, options SearchOptions) error {
	sql, params := repo.GenerateSelectSql(dest, options)
	if sql == "" {
		return errors.New("sql generation error")
	}
	sql += " limit 1"
	if len(options) > 0 {
		return repo.Modl.Dbx.Get(dest, sql, params...)
	}
	return nil
}

//Save provides Insert and Update for given type instance.
//When the id for give type instance(dest) is nil, it performs insert, otherwise, it performs update.
//Therefore, it assumes that the given type contain Id column as its primary key in the database
//Then the given dest contains other  keys than Id, please do not call this function for inserting
//,since it may cuase unexpected result
func (repo *DefaultRepository) Save(dest interface{}) error {
	v := reflect.ValueOf(dest)
	if v.FieldByName("Id").Interface().(int) <= 0 {
		return repo.Insert(dest)
	} else {
		return repo.Update(dest)
	}
}
func (repo *DefaultRepository) Insert(dest interface{}) error {
	if dest == nil {
		return errors.New("object can not be empty")
	}
	return repo.Modl.Insert(dest)
}
func (repo *DefaultRepository) Update(dest interface{}) error {
	if dest == nil {
		return errors.New("object can not be empty")
	}
	_, err := repo.Modl.Update(dest)
	return err

}

func (repo *DefaultRepository) GetList(dest interface{}, options SearchOptions, pos ...int) error {
	sql, params := repo.GenerateSelectSql(dest, options, pos...)
	if sql == "" {
		return errors.New("sql generation error")
	}
	if len(options) > 0 {
		return repo.Modl.Select(dest, sql, params...)
	} else {
		return repo.Modl.Select(dest, sql)
	}
}

func (repo *DefaultRepository) Count(dest interface{}, options SearchOptions) (int, error) {
	//sql, params := repo.GenerateSelectSql(dest, options)

	table := repo.Modl.TableFor(dest)

	if table == nil || len(table.Columns) < 1 {
		return 0, nil
	}
	sql := "select count(1) from " + table.TableName
	if sql == "" {
		return 0, errors.New("sql generation error")
	}
	s, params := repo.generateWhere(options)
	sql += s.String() + ";"
	var n int
	err := repo.Modl.Dbx.Get(&n, sql, params...)
	return n, err
}

func (repo *DefaultRepository) Delete(dest interface{}) error {
	if dest == nil {
		return errors.New("object can not be empty")
	}
	_, err := repo.Modl.Delete(dest)
	return err
}

//GenerateSelectSql generate select template for given search options
//note: in order to avoid sql injection, GenerateSelectSql function skip to
//fill the search option values in, instead, use ? as parameters so that
//necessary validation will be performed  by database/sql package.
func (repo *DefaultRepository) GenerateSelectSql(dest interface{}, options SearchOptions, pos ...int) (string, []interface{}) {
	table := repo.Modl.TableFor(dest)
	if table == nil || len(table.Columns) < 1 {
		return "", nil
	}
	sql := repo.getSelectAll(table)
	if sql == "" {
		return sql, nil
	}
	//generate where and params
	s, params := repo.generateWhere(options)
	//generate order by
	for i, col := range table.Keys {
		if i == 0 {
			s.WriteString(" order by ")
		}
		if i > 0 {
			s.WriteString(",")
		}
		s.WriteString(repo.Modl.Dialect.QuoteField(col.ColumnName))
	}

	if s.Len() > 0 {
		//generate limit offset if applicable
		if len(pos) == 2 {

			if pos[0] > pos[1] {
				pos[0], pos[1] = pos[1], pos[0]
			}
			if pos[0] <= 0 {
				pos[1] = 1
			}
			s.WriteString(" limit " + strconv.Itoa(pos[1]-pos[0]+1) + " offset " + strconv.Itoa(pos[0]-1))

		}
		sql += s.String()
	}
	return sql, params
}
func (repo *DefaultRepository) generateWhere(options SearchOptions) (bytes.Buffer, []interface{}) {
	s := bytes.Buffer{}

	x := 0
	var params []interface{}
	for key, val := range options {
		if x == 0 {
			params = make([]interface{}, len(options))
			s.WriteString(" where ")
		}
		if x > 0 {
			s.WriteString(" and ")
		}
		s.WriteString(repo.Modl.Dialect.QuoteField(key))
		s.WriteString("=")
		s.WriteString(repo.Modl.Dialect.BindVar(x))
		params[x] = val
		x++
	}
	return s, params
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

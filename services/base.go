package services

import (
	"fmt"

	"github.com/jmoiron/modl"
)

type SearchOptions map[string]interface{}

//Database related context
type DbContext struct {
	Modl *modl.DbMap
}

func (dbcontext *DbContext) GenerateWhere(options SearchOptions) string {
	where := ""
	for key, _ := range options {
		if where == "" {
			where = fmt.Sprintf(" %s = : %s", key, key)
		} else {
			where += fmt.Sprintf(" AND %s = : %s", key, key)
		}
	}
	return where

}

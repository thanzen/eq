package services

import "github.com/coopernurse/gorp"

//Database related context
type DbContext struct {
	Gorp *gorp.DbMap
}

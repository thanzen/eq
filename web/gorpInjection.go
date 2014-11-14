package web

import (
	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
)

//InjectGorp injects a db pointer and a DbMap(gorp) pointer to current gin.Context
func InjectGorp(dbMap *gorp.DbMap) gin.HandlerFunc {
	return func(c *gin.Context) {
		// inject db and gorp object
		c.Set(GORP_CONTEXT, dbMap)

		c.Next()
	}
}

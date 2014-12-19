package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/modl"
	_ "github.com/lib/pq"
	"github.com/thanzen/eq/models/user"
	"github.com/thanzen/eq/services"
	"github.com/thanzen/eq/singleton"
)

func initDb() *modl.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish

	db, err := sql.Open("postgres", "user=postgres password=root dbname=linkedtec sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &modl.DbMap{Db: db, Dialect: modl.PostgresDialect{}}

	// add a table, setting the table name to 'user_meta' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(user.User{}, "user_meta").SetKeys(true, "id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func main() {
	router := gin.Default()
	dbmap := initDb()
	defer dbmap.Db.Close()

	//router.Use(web.InjectGorp(dbmap))
	root := router.Group("/v1")
	gin.SetMode(gin.TestMode)
	router.LoadHTMLTemplates("templates/*")
	router.GET("/index", func(c *gin.Context) {
		obj := gin.H{"title": "Main website"}
		c.HTML(200, "index.tmpl", obj)
	})

	// Global middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	dbcontext := services.DbContext{dbmap}

	singleton.RegisterServices(&dbcontext)
	singleton.RegisterControllers(router, root)
	// Listen and server on 0.0.0.0:8080
	router.Run(":8080")

}

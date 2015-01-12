package main

import (
	//"database/sql"
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/thanzen/eq/models/user"
	//	"github.com/thanzen/eq/services"
	controller "github.com/thanzen/eq/controllers/singleton"
	service "github.com/thanzen/eq/services/singleton"
	"github.com/thanzen/modl"
	"log"
	"os"
)

func initDb() *modl.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish

	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=testdb sslmode=disable")
	//db, err = sqlx.Connect("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	checkErr(err, "sql.Open failed")

	dbmap := modl.NewDbMap(db.DB, modl.PostgresDialect{})

	//fmt.Println("connected!")
	// construct a gorp DbMap
	//dbmap := &modl.DbMap{Db: db.DB, Dialect: modl.PostgresDialect{}}

	// add a table, setting the table name to 'user_meta' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(user.User{}, "user_info").SetKeys(true, "id")
	dbmap.AddTableWithName(user.LoginAccount{}, "user_meta").SetKeys(true, "id")
	dbmap.AddTableWithName(user.UserType{}, "user_type").SetKeys(true, "id")
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
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	defer dbmap.Db.Close()

	//router.Use(web.InjectGorp(dbmap))
	root := router.Group("/api")
	//gin.SetMode(gin.TestMode)
	router.LoadHTMLTemplates("templates/*")
	router.GET("/index", func(c *gin.Context) {
		obj := gin.H{"title": "Main website"}
		c.HTML(200, "index.tmpl", obj)
	})

	// Global middlewares
	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	router.Use(CORSMiddleware())
	//dbcontext := services.DbContext{dbmap}

	service.RegisterServices(dbmap)
	controller.RegisterControllers(router, root)

	// Listen and server on 0.0.0.0:8080
	router.Run(":8080")
	// Turn off tracing
	dbmap.TraceOff()

}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, Content-Range, Content-Disposition, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.Abort(200)
			return
		}
		// c.Next()
	}
}

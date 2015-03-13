package test

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/thanzen/migrate/migrate"
	"os"
	"path/filepath"
	"strings"
)

func TestInitializeTestingDatabase(path string) {
	allErrors, ok := migrate.ResetSync("postgres://postgres:root@localhost:5432/eqtest?sslmode=disable", path)
	if !ok  {
		for _, err := range allErrors {
			fmt.Println(err)
		}
	}
}
func getMigrationFolder() string {
	path, _ := filepath.Abs(os.Args[0])
	pos := strings.Index(path, "\\eq\\")
	count := strings.Count(path[pos+3:], "\\") - 2
	fmt.Println(path[pos+3:])
	if count <= 0 {
		path = "./conf/db/migrations"
	} else {
		path = "./"
		for i := 0; i < count; i++ {
			path += "../"
		}
		path += "conf/db/migrations"
	}
	return path
}
func init() {
	TestInitializeTestingDatabase(getMigrationFolder())
}

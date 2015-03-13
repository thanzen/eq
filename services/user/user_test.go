package user

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
	model "github.com/thanzen/eq/models/user"
	_ "github.com/thanzen/eq/test"
	"testing"
)

func TestUserRegister(t *testing.T) {

	orm.RegisterDriver("postgres", orm.DR_Postgres)

	orm.RegisterDataBase("default", "postgres", "user=postgres password=root dbname=eqtest sslmode=disable")
	us := UserService{}
	var users []*model.User
	us.FuzzySearch(&users, "t", 2, 0, 200)

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		x := 1
		So(len(users), ShouldBeGreaterThan, 0)
		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}

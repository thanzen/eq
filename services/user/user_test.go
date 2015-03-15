package user

import (
	"github.com/astaxie/beego/orm"
	//	_ "github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/thanzen/eq/models/user"
	model "github.com/thanzen/eq/models/user"
	_ "github.com/thanzen/eq/test"
	"github.com/thanzen/eq/utils"
	"testing"
)

var userServ UserService = UserService{}

func TestUserSearchSpec(t *testing.T) {
	us := UserService{}
	var users []*model.User
	us.FuzzySearch(&users, "t", 2, 0, 200)

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		So(len(users), ShouldBeGreaterThan, 0)
	})
}

func TestPasswordSpec(t *testing.T) {
	password := "abcDefg01!"
    var err error
	Convey("Authentication Testing", t, func() {
		Convey("generateSalt()", func() {
			salt := GetUserSalt()
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, 10)
		})

		Convey("hashPassword()", func() {
			hash := utils.EncodePassword(password, GetUserSalt())
			So(hash, ShouldNotBeBlank)
			So(len(hash), ShouldEqual, 100)
		})

		Convey("Create a user", func() {
			u := new(user.User)
			u.Username = "testuser"
			u.Password = password
			db := orm.NewOrm()
			//ensure testuser not exist in the database
			_,err=db.Raw("delete from role where role_info_id = (select id from user_info where user_name = ? limit 1)", u.Username).Exec()
            if err != nil{
                Println(err)
            }
			_,err=db.Raw("delete from user_info where user_name=?", u.Username).Exec()
            if err != nil{
                Println(err)
            }
			err:=userServ.RegisterUser(u,"testuser","test@test.com",password,&model.UserType{Id:1},&model.Role{Id:2})
            if err != nil{
                Println(err)
            }
            So(err,ShouldEqual,nil)
			So(u.Id, ShouldBeGreaterThan, 0)

			Convey("VerifyUser()", func() {
               So(userServ.VerifyUser(u,u.Username,password),ShouldEqual,true)
			})
            
		})

	})
}

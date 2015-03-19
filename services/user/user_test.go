package user

import (
	"github.com/astaxie/beego/orm"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/thanzen/eq/models/user"
	model "github.com/thanzen/eq/models/user"
	_ "github.com/thanzen/eq/test"
	"github.com/thanzen/eq/utils"
	"testing"
)

var userServ UserService = UserService{}



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

            Convey("CanRegister()",func(){
                var canName, canEmail bool = true,true
                canName, canEmail,_= userServ.CanRegistered("testuser","test@test.com")
                So(canName,ShouldEqual,false)
                So(canEmail,ShouldEqual,false)
                canName, canEmail,_ =userServ.CanRegistered("testuser","test@whatever.com")
                So(canName,ShouldEqual,false)
                So(canEmail,ShouldEqual,true)
                canName, canEmail,_=userServ.CanRegistered("whatever","test@test.com")
                So(canName,ShouldEqual,true)
                So(canEmail,ShouldEqual,false)
                canName, canEmail,_=userServ.CanRegistered("whatever","test@whatever.com")
                So(canName,ShouldEqual,true)
                So(canEmail,ShouldEqual,true)
            })

            Convey("HasUser()",func(){
                So(userServ.HasUser(u,"testuser"),ShouldEqual,true)
            })

            Convey("SaveNewPassword()",func(){
                newPassword := "this is new pass!"
                userServ.SaveNewPassword(u,newPassword)
                So(userServ.VerifyUser(u,u.Username,password),ShouldEqual,false)
                So(userServ.VerifyUser(u,u.Username,newPassword),ShouldEqual,true)
            })

            Convey("FuzzySearch()", func() {
                var users []*model.User
                userServ.FuzzySearch(&users, "t", 2, 0, 200)
                So(len(users),ShouldBeGreaterThanOrEqualTo, 1)
            })
		})

	})
}

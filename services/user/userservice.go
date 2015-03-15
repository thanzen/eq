package user

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/thanzen/eq/cachemanager"
	"github.com/thanzen/eq/models/user"
	. "github.com/thanzen/eq/services"
	"github.com/thanzen/eq/setting"
	"github.com/thanzen/eq/utils"
	"strings"
)

const USER_ID_CACHE_PATTERN = "eq_user_id_%d_deleted_%t"
const USER_CACHE_PATTERN = "eq_user_"

type UserService struct {
}

func (this UserService) Queryable() orm.QuerySeter {
	return orm.NewOrm().QueryTable("user_info").OrderBy("-Id")
}

func (this UserService) Insert(u *user.User) error {
	return this.InsertWithScope(orm.NewOrm(), u)
}

func (this UserService) InsertWithScope(tr orm.Ormer, u *user.User) error {
	u.PasswordSalt = GetUserSalt()
	if _, err := tr.Insert(u); err != nil {
		return err
	}
	return nil
}

func (this UserService) Read(u *user.User, fields ...string) error {
	if len(fields) == 1 && strings.ToUpper(fields[0]) == "ID" {
		return this.getById(u)
	} else if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (this UserService) Update(u *user.User, fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	//invalid cache
	cachemanager.Delete(fmt.Sprintf(USER_ID_CACHE_PATTERN, u.Id, u.Deleted))
	return nil
}

func (this UserService) Delete(u *user.User) error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}

	//invalid cache
	cachemanager.Delete(fmt.Sprintf(USER_ID_CACHE_PATTERN, u.Id, u.Deleted))
	return nil
}

func (this UserService) Query(users *[]*user.User, options SearchOptions) (int64, error) {
	qs := this.Queryable()
	cond := GenerateCondition(options)
	qs.SetCond(cond)
	return qs.All(users)
}

func (this UserService) getById(u *user.User) error {
	var err error
	cache, hit := cachemanager.Get(fmt.Sprintf(USER_ID_CACHE_PATTERN, u.Id, u.Deleted), func(params ...interface{}) interface{} {
		if err = this.Read(u, "Id", "Deleted"); err != nil {
			return nil
		}
		return *u
	})
	if hit {
		*u = cache.(user.User)
	}
	return err
}

func (this *UserService) CanRegistered(userName string, email string) (canName bool, canEmail bool, err error) {
	cond := orm.NewCondition()
	cond = cond.Or("Username", userName).Or("email", email)
	var maps []orm.Params
    var n int64
	n, err = this.Queryable().SetCond(cond).Values(&maps, "Username", "email")
	if err != nil {
		return false, false, err
	}

    canName = true
    canEmail = true

	if n > 0 {
		for _, m := range maps {
			if canName && orm.ToStr(m["Username"]) == userName {
                canName = false
			}
			if canEmail && orm.ToStr(m["Email"]) == email {
                canEmail = false
			}
		}
	}

	return canName, canEmail, nil
}

// check if exist user by username or email, ignore "deleted" users
func (this *UserService) HasUser(user *user.User, username string) bool {
	var err error
	qs := orm.NewOrm()
	if strings.IndexRune(username, '@') == -1 {
		user.Username = username
		err = qs.Read(user, "Username", "Deleted")
	} else {
		user.Email = username
		err = qs.Read(user, "Email", "Deleted")
	}
	if err == nil {
		return true
	}
	return false
}

// register a user
func (this *UserService) RegisterUser(u *user.User, username, email, password string, userType *user.UserType,role *user.Role) error {
	u.UserType = userType
	// use random salt encode password
	salt := GetUserSalt()
	pwd := utils.EncodePassword(password, salt)

	u.Username = strings.ToLower(username)
	u.Email = strings.ToLower(email)

	// save salt and encode password, use $ as split char
	u.Password = fmt.Sprintf("%s$%s", salt, pwd)
	u.PasswordSalt = salt

	var err error
	tr := orm.NewOrm()
	tr.Begin()
	if err = this.InsertWithScope(tr, u); err == nil {
		roleService := RoleService{}
		err = roleService.InsertUsersWithScope(tr, role, u)
	}
	if err == nil {
		tr.Commit()
	} else {
		tr.Rollback()
	}
	return err
}

// set a new password to user
func (this *UserService) SaveNewPassword(u *user.User, password string) error {
	salt := GetUserSalt()
	u.Password = fmt.Sprintf("%s$%s", salt, utils.EncodePassword(password, salt))
    u.PasswordSalt = salt
	_, err := orm.NewOrm().Update(u, "Password", "PasswordSalt", "Updated")
	return err
}

// verify username/email and password
func (this *UserService) VerifyUser(u *user.User, username, password string) (success bool) {
	// search user by username or email
	if this.HasUser(u, username) == false {
		return
	}

	if this.VerifyPassword(password, u.Password) {
		// success
		success = true

		// re-save discuz password
		if len(u.Password) == 39 {
			if err := this.SaveNewPassword(u, password); err != nil {
				beego.Error("SaveNewPassword err: ", err.Error())
			}
		}
	}
	return
}

// compare raw password and encoded password
func (this *UserService) VerifyPassword(rawPwd, encodedPwd string) bool {

	// for discuz accounts
	if len(encodedPwd) == 39 {
		salt := encodedPwd[:6]
		encoded := encodedPwd[7:]
		return encoded == utils.EncodeMd5(utils.EncodeMd5(rawPwd)+salt)
	}

	// split
	var salt, encoded string
	if len(encodedPwd) > 11 {
		salt = encodedPwd[:10]
		encoded = encodedPwd[11:]
	}

	return utils.EncodePassword(rawPwd, salt) == encoded
}

// get user by erify code
func (this *UserService) getVerifyUser(u *user.User, code string) bool {
	if len(code) <= utils.TimeLimitCodeLength {
		return false
	}

	// use tail hex username query user
	hexStr := code[utils.TimeLimitCodeLength:]
	if b, err := hex.DecodeString(hexStr); err == nil {
		u.Username = string(b)
		if this.Read(u, "Username") == nil {
			return true
		}
	}

	return false
}

// verify active code when active account
func (this *UserService) VerifyUserActiveCode(u *user.User, code string) bool {
	minutes := setting.ActiveCodeLives

	if this.getVerifyUser(u, code) {
		// time limit code
		prefix := code[:utils.TimeLimitCodeLength]
		data := utils.ToStr(u.Id) + u.Email + u.Username + u.Password + u.PasswordSalt

		return utils.VerifyTimeLimitCode(data, minutes, prefix)
	}

	return false
}

// create a time limit code for user active
func (this *UserService) CreateUserActiveCode(u *user.User, startInf interface{}) string {
	minutes := setting.ActiveCodeLives
	data := utils.ToStr(u.Id) + u.Email + u.Username + u.Password + u.PasswordSalt
	code := utils.CreateTimeLimitCode(data, minutes, startInf)

	// add tail hex username
	code += hex.EncodeToString([]byte(u.Username))
	return code
}

// verify code when reset password
func (this *UserService) VerifyUserResetPwdCode(u *user.User, code string) bool {
	minutes := setting.ResetPwdCodeLives

	if this.getVerifyUser(u, code) {
		// time limit code
		prefix := code[:utils.TimeLimitCodeLength]
		data := utils.ToStr(u.Id) + u.Email + u.Username + u.Password + u.PasswordSalt + u.LastActivityDate.String()

		return utils.VerifyTimeLimitCode(data, minutes, prefix)
	}

	return false
}

// create a time limit code for user reset password
func (this *UserService) CreateUserResetPwdCode(u *user.User, startInf interface{}) string {
	minutes := setting.ResetPwdCodeLives
	data := utils.ToStr(u.Id) + u.Email + u.Username + u.Password + u.PasswordSalt + u.LastActivityDate.String()
	code := utils.CreateTimeLimitCode(data, minutes, startInf)

	// add tail hex username
	code += hex.EncodeToString([]byte(u.Username))
	return code
}

// load roles
func (this *UserService) LoadRoles(u *user.User) error {
	if len(u.Roles) > 0 {
		return nil
	}
	_, err := orm.NewOrm().LoadRelated(u, "Roles")
	if err == nil {
		for _, r := range u.Roles {
			if r.IsSystemRole {
				u.IsSystemAccount = true
			}
		}
		cachemanager.Put(fmt.Sprintf(USER_ID_CACHE_PATTERN, u.Id, u.Deleted), *u)
	}
	return err
}

// HasPermission checks whether the user has the given permission
func (this *UserService) HasPermission(u *user.User, permission string) bool {
	if u == nil {
		return false
	}
	err := this.LoadRoles(u)
	if err != nil {
		return false
	}
	roleService := RoleService{}
	allRoles := roleService.GetAllRoles()
	for _, r := range u.Roles {
		if roleService.HasPermission(allRoles[r.Id], permission) {
			return true
		}
	}
	return false
}

func (this *UserService) FuzzySearch(users *[]*user.User, text string, roleId int64, offset int64, limit int64) (n int64,err error) {
    //todo: evaluate return value n for total number with a extra flag, currently return 0 only, we can
    // make another query to get total, even using go routine to perform concurrent query
	if limit <= 0 || offset < 0 {
		return 0, errors.New("invalid range")
	}
    sql:="SELECT user_info.id, user_name,email,first_name,last_name,cell_phone,office_phone,company,active from user_info"
	if roleId > 0 {
        sql+=" INNER JOIN user_role ON user_info.Id = user_role.user_info_id AND user_role.role_id = ?"
	}
    sql += " WHERE UPPER(user_name) LIKE UPPER(?) OR UPPER(email) LIKE UPPER(?)"
    sql +=" OR UPPER(first_name) LIKE UPPER(?) OR UPPER(last_name) LIKE UPPER(?)"
    sql+=" OR UPPER(cell_phone) LIKE UPPER(?) OR UPPER(office_phone) LIKE UPPER(?)"
    sql+= " OR UPPER(company) LIKE UPPER(?) ORDER BY id DESC OFFSET ? LIMIT ?"
    text = "%"+text+"%"
    if roleId > 0 {
        n,err =orm.NewOrm().Raw(sql,roleId,text,text,text,text,text,text,text,offset,limit).QueryRows(users)
    }else {
        n,err =orm.NewOrm().Raw(sql,text,text,text,text,text,text,text,offset,limit).QueryRows(users)
    }
    if err!=nil{
        beego.Info(err)
    }
	return 0, err
}




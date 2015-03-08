package user

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
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

func (this *UserService) CanRegistered(userName string, email string) (bool, bool, error) {
	cond := orm.NewCondition()
	cond = cond.Or("Username", userName).Or("email", email)
	var maps []orm.Params
	n, err := this.Queryable().SetCond(cond).Values(&maps, "Username", "email")
	if err != nil {
		return false, false, err
	}

	e1 := true
	e2 := true

	if n > 0 {
		for _, m := range maps {
			if e1 && orm.ToStr(m["Username"]) == userName {
				e1 = false
			}
			if e2 && orm.ToStr(m["Email"]) == email {
				e2 = false
			}
		}
	}

	return e1, e2, nil
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

// register a regular user user
//todo: refactor this method later to enable register different type users(UserType and Role are different)
func (this *UserService) RegisterRegularUser(u *user.User, username, email, password string) error {
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
		r := &user.Role{Id: 1}
		err = roleService.InsertUsersWithScope(tr, r, u)
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

// get login redirect url from cookie
func (this *UserService) GetLoginRedirect(ctx *context.Context) string {
	loginRedirect := strings.TrimSpace(ctx.GetCookie("login_to"))
	if utils.IsMatchHost(loginRedirect) == false {
		loginRedirect = "/"
	} else {
		ctx.SetCookie("login_to", "", -1, "/")
	}
	return loginRedirect
}

// login user
//todo: added token (oauth2) session if it is rest api call
func (this *UserService) LoginUser(u *user.User, ctx *context.Context, remember bool) {
	// werid way of beego session regenerate id...
	ctx.Input.CruSession.SessionRelease(ctx.ResponseWriter)
	ctx.Input.CruSession = beego.GlobalSessions.SessionRegenerateId(ctx.ResponseWriter, ctx.Request)
	ctx.Input.CruSession.Set("auth_user_id", u.Id)

	if remember {
		this.WriteRememberCookie(u, ctx)
	}
}
func (this *UserService) WriteRememberCookie(u *user.User, ctx *context.Context) {
	secret := utils.EncodeMd5(u.PasswordSalt + u.Password)
	days := 86400 * setting.LoginRememberDays
	ctx.SetCookie(setting.CookieUsername, u.Username, days)
	ctx.SetSecureCookie(secret, setting.CookieRememberName, u.Username, days)
}
func (this *UserService) DeleteRememberCookie(ctx *context.Context) {
	ctx.SetCookie(setting.CookieUsername, "", -1)
	ctx.SetCookie(setting.CookieRememberName, "", -1)
}

func (this *UserService) LoginUserFromRememberCookie(u *user.User, ctx *context.Context) (success bool) {
	userName := ctx.GetCookie(setting.CookieUsername)
	if len(userName) == 0 {
		return false
	}

	defer func() {
		if !success {
			this.DeleteRememberCookie(ctx)
		}
	}()

	u.Username = userName
	if err := this.Read(u, "Username"); err != nil {
		return false
	}

	secret := utils.EncodeMd5(u.PasswordSalt + u.Password)
	value, _ := ctx.GetSecureCookie(secret, setting.CookieRememberName)
	if value != userName {
		return false
	}

	this.LoginUser(u, ctx, true)

	return true
}

// logout user
func (this *UserService) LogoutUser(ctx *context.Context) {
	this.DeleteRememberCookie(ctx)
	ctx.Input.CruSession.Delete("auth_user_id")
	ctx.Input.CruSession.Flush()
	beego.GlobalSessions.SessionDestroy(ctx.ResponseWriter, ctx.Request)
}

func (this *UserService) GetUserIdFromSession(sess session.SessionStore) int64 {
	if id, ok := sess.Get("auth_user_id").(int64); ok && id > 0 {
		return id
	}
	return 0
}

// get user if key exist in session
func (this *UserService) GetUserFromSession(u *user.User, sess session.SessionStore) bool {
	id := this.GetUserIdFromSession(sess)
	if id > 0 {
		u.Id = id
		if this.Read(u) == nil {
			return true
		}
	}

	return false
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

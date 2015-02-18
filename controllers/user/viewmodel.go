package user

import (
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
	"github.com/thanzen/eq/models/user"
	"github.com/thanzen/eq/services"
	userServ "github.com/thanzen/eq/services/user"
	"github.com/thanzen/eq/setting"
	"github.com/thanzen/eq/utils"
)

// Register model
type RegisterForm struct {
	Username   string      `valid:"Required;AlphaDash;MinSize(5);MaxSize(30)"`
	Email      string      `valid:"Required;Email;MaxSize(80)"`
	Password   string      `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
	PasswordRe string      `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
	Captcha    string      `form:"type(captcha)" valid:"Required"`
	CaptchaId  string      `form:"type(empty)"`
	Locale     i18n.Locale `form:"-"`
}

func (form *RegisterForm) Valid(v *validation.Validation) {

	// Check if passwords of two times are same.
	if form.Password != form.PasswordRe {
		v.SetError("PasswordRe", "auth.repassword_not_match")
		return
	}
	serv := &userServ.UserService{}

	e1, e2, _ := serv.CanRegistered(form.Username, form.Email)

	if !e1 {
		v.SetError("Username", "auth.username_already_taken")
	}

	if !e2 {
		v.SetError("Email", "auth.email_already_taken")
	}

	if !setting.Captcha.Verify(form.CaptchaId, form.Captcha) {
		v.SetError("Captcha", "auth.captcha_wrong")
	}
}

func (form *RegisterForm) Labels() map[string]string {
	return map[string]string{
		"Username":   "auth.login_username",
		"Email":      "auth.login_email",
		"Password":   "auth.login_password",
		"PasswordRe": "auth.retype_password",
		"Captcha":    "auth.captcha",
	}
}

func (form *RegisterForm) Helps() map[string]string {
	return map[string]string{
		"Username": form.Locale.Tr("valid.min_length_is", 5) + ", " + form.Locale.Tr("valid.only_contains", "a-z 0-9 - _"),
		"Captcha":  "auth.captcha_click_refresh",
	}
}

func (form *RegisterForm) Placeholders() map[string]string {
	return map[string]string{
		"Username":   "auth.plz_enter_username",
		"Email":      "auth.plz_enter_email",
		"Password":   "auth.plz_enter_password",
		"PasswordRe": "auth.plz_reenter_password",
		"Captcha":    "auth.plz_enter_captcha",
	}
}

// Login form
type LoginForm struct {
	Username string `valid:"Required"`
	Password string `form:"type(password)" valid:"Required"`
	Remember bool
}

func (model *LoginForm) Labels() map[string]string {
	return map[string]string{
		"Username": "auth.username_or_email",
		"Password": "auth.login_password",
		"Remember": "auth.login_remember_me",
	}
}

// Forgot model
type ForgotModel struct {
	Email string     `valid:"Required;Email;MaxSize(80)"`
	User  *user.User `form:"-"`
}

func (model *ForgotModel) Labels() map[string]string {
	return map[string]string{
		"Email": "auth.login_email",
	}
}

func (model *ForgotModel) Helps() map[string]string {
	return map[string]string{
		"Email": "auth.forgotform_email_help",
	}
}

func (model *ForgotModel) Valid(v *validation.Validation) {
	serv := &userServ.UserService{}
	if serv.HasUser(model.User, model.Email) == false {
		v.SetError("Email", "auth.forgotform_wrong_email")
	}
}

// Reset password model
type ResetPwdModel struct {
	Password   string `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
	PasswordRe string `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
}

func (model *ResetPwdModel) Valid(v *validation.Validation) {
	// Check if passwords of two times are same.
	if model.Password != model.PasswordRe {
		v.SetError("PasswordRe", "auth.repassword_not_match")
		return
	}
}

func (model *ResetPwdModel) Labels() map[string]string {
	return map[string]string{
		"PasswordRe": "auth.retype_password",
	}
}

func (model *ResetPwdModel) Placeholders() map[string]string {
	return map[string]string{
		"Password":   "auth.plz_enter_password",
		"PasswordRe": "auth.plz_reenter_password",
	}
}

// Settings Profile form
type ProfileForm struct {
	NickName    string `valid:"Required;MaxSize(30)"`
	Url         string `valid:"MaxSize(100)"`
	Company     string `valid:"MaxSize(30)"`
	Location    string `valid:"MaxSize(30)"`
	Info        string `form:"type(textarea)" valid:"MaxSize(255)"`
	Email       string `valid:"Required;Email;MaxSize(100)"`
	PublicEmail bool   `valid:""`
	GrEmail     string `valid:"Required;MaxSize(80)"`
	Github      string `valid:"MaxSize(30)"`
	Twitter     string `valid:"MaxSize(30)"`
	Google      string `valid:"MaxSize(30)"`
	Weibo       string `valid:"MaxSize(30)"`
	Linkedin    string `valid:"MaxSize(30)"`
	Facebook    string `valid:"MaxSize(30)"`
	Lang        int    `form:"type(select);attr(rel,select2)" valid:""`
	//LangAdds    models.SliceStringField `form:"type(select);attr(rel,select2);attr(multiple,multiple)" valid:""`
	Locale i18n.Locale `form:"-"`
}

func (form *ProfileForm) LangSelectData() [][]string {
	langs := setting.Langs
	data := make([][]string, 0, len(langs))
	for i, lang := range langs {
		data = append(data, []string{lang, utils.ToStr(i)})
	}
	return data
}

func (form *ProfileForm) LangAddsSelectData() [][]string {
	langs := setting.Langs
	data := make([][]string, 0, len(langs))
	for i, lang := range langs {
		data = append(data, []string{lang, utils.ToStr(i)})
	}
	return data
}

func (form *ProfileForm) Valid(v *validation.Validation) {
	if len(i18n.GetLangByIndex(form.Lang)) == 0 {
		v.SetError("Lang", "Can not be empty")
	}

	//if len(model.LangAdds) > 0 {
	//	adds := make(models.SliceStringField, 0, len(model.LangAdds))
	//	for _, l := range model.LangAdds {
	//		if d, err := utils.StrTo(l).Int(); err == nil {
	//			if model.Lang == d {
	//				continue
	//			}
	//			if len(i18n.GetLangByIndex(model.Lang)) == 0 {
	//				v.SetError("Lang", "Can not be empty")
	//				return
	//			}
	//			adds = append(adds, l)
	//		}
	//	}
	//	model.LangAdds = adds
	//}
}

func (form *ProfileForm) SetFromUser(u *user.User) {
	utils.SetFormValues(u, form)
}

func (form *ProfileForm) SaveUserProfile(u *user.User) error {
	// set md5 value if the value is an email
	if strings.IndexRune(form.GrEmail, '@') != -1 {
		form.GrEmail = utils.EncodeMd5(form.GrEmail)
	}

	changes := utils.FormChanges(u, form)
	if len(changes) > 0 {
		// if email changed then need re-active
		if u.Email != form.Email {
			u.Active = false
			changes = append(changes, "Active")
		}
		serv := &userServ.UserService{}
		utils.SetFormValues(form, u)
		return serv.Update(u, changes...)
	}
	return nil
}

func (form *ProfileForm) Labels() map[string]string {
	return map[string]string{
		"Lang":        "auth.profile_lang",
		"LangAdds":    "auth.profile_lang_additional",
		"NickName":    "model.user_nickname",
		"PublicEmail": "auth.profile_publicemail",
		"GrEmail":     "auth.profile_gremail",
		"Info":        "auth.profile_info",
		"Company":     "model.user_company",
		"Location":    "model.user_location",
		"Google":      ".Google+",
	}
}

func (form *ProfileForm) Helps() map[string]string {
	return map[string]string{
		"GrEmail": "auth.profile_gremail_help",
		"Info":    "auth.plz_enter_your_info",
	}
}

func (form *ProfileForm) Placeholders() map[string]string {
	return map[string]string{
		"GrEmail": "auth.plz_enter_gremail",
		"Url":     "auth.plz_enter_website",
	}
}

// Change password form
type PasswordForm struct {
	PasswordOld string     `form:"type(password)" valid:"Required"`
	Password    string     `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
	PasswordRe  string     `form:"type(password)" valid:"Required;MinSize(4);MaxSize(30)"`
	User        *user.User `form:"-"`
}

func (form *PasswordForm) Valid(v *validation.Validation) {
	// Check if passwords of two times are same.
	if form.Password != form.PasswordRe {
		v.SetError("PasswordRe", "auth.repassword_not_match")
		return
	}

	serv := &userServ.UserService{}
	if serv.VerifyPassword(form.PasswordOld, form.User.Password) == false {
		v.SetError("PasswordOld", "auth.old_password_wrong")
	}
}

func (form *PasswordForm) Labels() map[string]string {
	return map[string]string{
		"PasswordOld": "auth.old_password",
		"Password":    "auth.new_password",
		"PasswordRe":  "auth.retype_password",
	}
}

func (form *PasswordForm) Placeholders() map[string]string {
	return map[string]string{
		"PasswordOld": "auth.plz_enter_old_password",
		"Password":    "auth.plz_enter_new_password",
		"PasswordRe":  "auth.plz_reenter_password",
	}
}

type UserAdminModel struct {
	Create      bool   `form:"-"`
	Id          int    `form:"-"`
	Username    string `valid:"Required;AlphaDash;MinSize(5);MaxSize(30)"`
	Email       string `valid:"Required;Email;MaxSize(100)"`
	PublicEmail bool   ``
	NickName    string `valid:"Required;MaxSize(30)"`
	Url         string `valid:"MaxSize(100)"`
	Company     string `valid:"MaxSize(30)"`
	Location    string `valid:"MaxSize(30)"`
	Info        string `form:"type(textarea)" valid:"MaxSize(255)"`
	GrEmail     string `valid:"Required;MaxSize(80)"`
	Github      string `valid:"MaxSize(30)"`
	Twitter     string `valid:"MaxSize(30)"`
	Google      string `valid:"MaxSize(30)"`
	Weibo       string `valid:"MaxSize(30)"`
	Linkedin    string `valid:"MaxSize(30)"`
	Facebook    string `valid:"MaxSize(30)"`
	Followers   int    ``
	Following   int    ``
	IsAdmin     bool   ``
	IsActive    bool   ``
	IsForbid    bool   ``
	Lang        int    `form:"type(select);attr(rel,select2)" valid:""`
	//LangAdds    models.SliceStringField `form:"type(select);attr(rel,select2);attr(multiple,multiple)" valid:""`
}

func (model *UserAdminModel) LangSelectData() [][]string {
	langs := setting.Langs
	data := make([][]string, 0, len(langs))
	for i, lang := range langs {
		data = append(data, []string{lang, utils.ToStr(i)})
	}
	return data
}

func (model *UserAdminModel) LangAddsSelectData() [][]string {
	langs := setting.Langs
	data := make([][]string, 0, len(langs))
	for i, lang := range langs {
		data = append(data, []string{lang, utils.ToStr(i)})
	}
	return data
}

func (model *UserAdminModel) Valid(v *validation.Validation) {
	qs := userServ.UserService{}.Queryable()
	if services.CheckIsExist(qs, "user_name", model.Username, model.Id) {
		v.SetError("Username", "auth.username_already_taken")
	}

	if services.CheckIsExist(qs, "email", model.Email, model.Id) {
		v.SetError("Email", "auth.email_already_taken")
	}

	if len(i18n.GetLangByIndex(model.Lang)) == 0 {
		v.SetError("Lang", "Can not be empty")
	}

	//if len(model.LangAdds) > 0 {
	//	adds := make(models.SliceStringField, 0, len(model.LangAdds))
	//	for _, l := range model.LangAdds {
	//		if d, err := utils.StrTo(l).Int(); err == nil {
	//			if model.Lang == d {
	//				continue
	//			}
	//			if len(i18n.GetLangByIndex(model.Lang)) == 0 {
	//				v.SetError("Lang", "Can not be empty")
	//				return
	//			}
	//			adds = append(adds, l)
	//		}
	//	}
	//	model.LangAdds = adds
	//}
}

func (model *UserAdminModel) Helps() map[string]string {
	return nil
}

func (model *UserAdminModel) Labels() map[string]string {
	return nil
}

func (model *UserAdminModel) SetFromUser(u *user.User) {
	utils.SetFormValues(u, model)
}

func (model *UserAdminModel) SetToUser(u *user.User) {
	// set md5 value if the value is an email
	if strings.IndexRune(model.GrEmail, '@') != -1 {
		model.GrEmail = utils.EncodeMd5(model.GrEmail)
	}

	utils.SetFormValues(model, u)
}

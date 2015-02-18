package email

import (
	"github.com/astaxie/beego"
	"github.com/thanzen/eq/models/user"
	"github.com/thanzen/eq/setting"
)

// Create New mail message use MailFrom and MailUser
func NewMailMessage(To []string, subject, body string) Message {
	msg := NewHtmlMessage(To, setting.MailFrom, subject, body)
	msg.User = setting.MailUser
	return msg
}

func GetMailTmplData(lang string, u *user.User) map[interface{}]interface{} {
	data := make(map[interface{}]interface{}, 10)
	data["AppName"] = beego.AppName
	data["AppVer"] = setting.AppVer
	data["AppUrl"] = setting.AppUrl
	data["AppLogo"] = setting.AppLogo
	data["IsProMode"] = setting.IsProMode
	data["Lang"] = lang
	data["ActiveCodeLives"] = setting.ActiveCodeLives
	data["ResetPwdCodeLives"] = setting.ResetPwdCodeLives
	if u != nil {
		data["User"] = u
	}
	return data
}

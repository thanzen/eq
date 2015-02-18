package base

import (
"github.com/thanzen/eq/services/email"
)

type TestRouter struct {
    BaseRouter
}

func (this *TestRouter) Get() {
    this.TplNames = this.GetString(":tmpl")
    this.Data = email.GetMailTmplData(this.Locale.Lang, &this.User)
    this.Data["Code"] = "CODE"
}

package controllers

import (
	"github.com/thanzen/eq/controllers/base"
)

type MainController struct {
	base.BaseRouter
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}
func (c *MainController) Index() {
	c.TplNames = "index.html"
}

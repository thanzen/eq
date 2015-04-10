package user

import (
    "github.com/thanzen/eq/controllers/base"
)

type AdminController struct {
    base.BaseController
}

func (this *AdminController) Index()  {
    this.TplNames = "admin/index.html"
}
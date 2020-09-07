package controllers

import "github.com/astaxie/beego"

type AdminController struct {
	beego.Controller
}

func (admin *AdminController) init()  {
}

func (admin *AdminController) Get() {
	admin.Data["title"] = "网站后台管理模版"

	admin.TplName = "admin/index.html"
}
func (admin *AdminController) Welcome()  {
	admin.Layout = "admin/admin.tpl"
	admin.TplName = "admin/welcome.html"
}
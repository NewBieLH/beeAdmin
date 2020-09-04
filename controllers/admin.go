package controllers

import "github.com/astaxie/beego"

type AdminController struct {
	beego.Controller
}

func (admin *AdminController) Get() {
	admin.Data["title"] = "网站后台管理模版"
	admin.Layout = "admin/admin.tpl"
	admin.TplName = "admin/welcome.html"
}

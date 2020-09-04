package controllers

import "github.com/astaxie/beego"

type AdminController struct {
	beego.Controller
}

func (admin *AdminController)init () {
	admin.Data["title"] = "网站后台管理模版"
	admin.Layout = ""
}
func (admin *AdminController) Get() {
	admin.TplName = "admin/index.html"
}

func (admin *AdminController) Welcome() {
	admin.TplName = "admin/welcome.html"
}
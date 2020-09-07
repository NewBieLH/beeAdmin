package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (user *UserController) List()  {
	user.Layout = "admin/admin.tpl"
	user.TplName = "admin/user/list.html"
}
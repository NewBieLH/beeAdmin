package routers

import (
	"beeAdmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{})
}

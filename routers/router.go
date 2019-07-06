package routers

import (
	"webhooks/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hooks", &controllers.HooksController{})
}

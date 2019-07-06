package routers

import (
	"webhooks/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hooks", &controllers.HooksController{})
	
	beego.Get("/hello",func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})

}

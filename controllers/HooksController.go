package controllers

import (
	"github.com/astaxie/beego"
)

type HooksController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "ucat.me"
	c.Data["Email"] = "longy0115@gmail.com"
	c.TplName = "index.tpl"
}

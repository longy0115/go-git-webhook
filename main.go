package main

import (
	_ "webhooks/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}


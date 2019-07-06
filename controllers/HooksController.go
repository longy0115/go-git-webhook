package controllers

import (
	"github.com/astaxie/beego"
	"os/exec"
)

type HooksController struct {
	beego.Controller
}

func (c *HooksController) Get() {
	command := "../gohooks.sh"
	exec.Command("/bin/sh", "-c", command)
     // 略 Other Code
     // 永远也不可能执行到这里来
     // bash 这条命令 (kill $pidid|go run main.go&)
     // 会杀掉现在的进程，程序直接结束了，不可能执行后面的 go run main.go &
}

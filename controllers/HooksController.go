package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"log"
	"os/exec"
)

const ShellToUse = "bash"

type HooksController struct {
	beego.Controller
}

func (this *HooksController) Get() {

	err, out, errout := Shellout("cd /home/Go/golang/src/webhooks && git pull origin master")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	this.Ctx.WriteString("--- stdout ---")
	this.Ctx.WriteString(out)
	this.Ctx.WriteString("--- stderr ---")
	this.Ctx.WriteString(errout)
	this.Ctx.WriteString("succ----")
}

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

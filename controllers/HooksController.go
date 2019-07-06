package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"os/exec"
)

const ShellToUse = "bash"

type HooksController struct {
	beego.Controller
}

func (c *HooksController) Get() {
	err, out, errout := Shellout("ls -ltr")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println("--- stdout ---")
	fmt.Println(out)
	fmt.Println("--- stderr ---")
	fmt.Println(errout)
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

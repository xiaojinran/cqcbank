package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os/exec"
)

type CommandController struct {
	beego.Controller
}


// @Title Commands
// @Description Commands
// @Success 200
// @router / [get]
func (c *CommandController)Get(){
	
	cmd := exec.Command(addPrefix("demo.bat"))
	out,err := cmd.CombinedOutput()
	fmt.Println(err)
	if err != nil {
	//	log.Fatal(err)
	}
	fmt.Println(string(out))
	c.Data["json"] = string(out)
	c.ServeJSON()
}

func addPrefix(c string) string{
	return "scripts\\" + c
}

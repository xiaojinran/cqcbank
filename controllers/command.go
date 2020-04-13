package controllers

import (
	"cqcbank/models"
	"fmt"
	"github.com/astaxie/beego"
)

type CommandController struct {
	beego.Controller
}


// @Title Get
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:scriptsName [get]
func (c *CommandController)Get(){
	scriptsName := c.Ctx.Input.Param(":scriptsName")
	fmt.Println(scriptsName)
	cr := models.ExeSysCommand(scriptsName)
	c.Data["json"] = cr
	c.ServeJSON()
}






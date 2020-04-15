package controllers

import (
	"cqcbank/models"
	"github.com/astaxie/beego"
)

type CommandController struct {
	beego.Controller
}


// @Title Post
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:scriptsName [post]
func (c *CommandController)Post(){
	scriptsName := c.Ctx.Input.Param(":scriptsName")
    //fmt.Println(string(c.Ctx.Input.RequestBody))
	models.WriteRequestBody(scriptsName,c.Ctx.Input.RequestBody)
	cr := models.ExeSysCommand(scriptsName+".sh")
	c.Data["json"] = cr
	c.ServeJSON()
}






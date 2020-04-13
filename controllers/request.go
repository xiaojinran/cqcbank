package controllers

import (
	"cqcbank/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RequestController struct {
	beego.Controller
}


// @Title Get
// @Description Execute scripts by Name
// @Param	scriptsName		path 	string	true		"the script you want to run"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:fileName [get]
func (c *RequestController)Get(){
	fileName := c.Ctx.Input.Param(":fileName")
	fmt.Println(fileName)
	cr := models.RequestServer(fileName)
	fmt.Println(cr)
	c.Data["json"] = cr
	c.ServeJSON()
}






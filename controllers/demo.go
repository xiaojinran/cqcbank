package controllers

import (
	"cqcbank/models"
	"github.com/astaxie/beego"
)
type DemoController struct {
	beego.Controller
}

// @Title Demo
// @Description Demo
// @Success 200 {object} models.Demo
// @router / [get]
func (o *DemoController) Get() {
	obs := []*models.Demo{
		models.NewDemo("设备1",21),
		models.NewDemo("设备2",18),
		models.NewDemo("设备3",19),
	}
	
	
	o.Data["json"] = obs
	o.ServeJSON()
}



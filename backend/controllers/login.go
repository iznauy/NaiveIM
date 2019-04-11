package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}


func (control *LoginController) Post() {
	name := control.GetString("name")
	if len(name) == 0 {
		control.Ctx.WriteString("error")
		return
	}
	control.Ctx.WriteString("ok")
}
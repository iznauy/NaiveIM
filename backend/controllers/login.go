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
		control.Redirect("/ok", 302)
		return
	}
	control.Redirect("/room", 302)
}
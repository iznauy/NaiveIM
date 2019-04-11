package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

type LoginController struct {
	beego.Controller
}


func (controller *LoginController) Post() {
	name := controller.GetString("name")
	if len(name) == 0 {
		http.Error(controller.Ctx.ResponseWriter, "error!", 400)
		return
	}

}
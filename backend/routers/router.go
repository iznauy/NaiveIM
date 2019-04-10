package routers

import (
	"NaiveIM/backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/room", &controllers.WebSocketController{})
}

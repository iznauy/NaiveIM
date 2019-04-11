package routers

import (
	"NaiveIM/backend/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: 	[]string{"http://localhost:8081"},
		AllowMethods: 	[]string{"GET", "POST"},
		AllowHeaders: 	[]string{"Origin"},
		ExposeHeaders:	[]string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.LoginController{})
    beego.Router("/room", &controllers.WebSocketController{})
}

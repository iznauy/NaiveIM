package main

import (
	_ "NaiveIM/backend/controllers"
	_ "NaiveIM/backend/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}


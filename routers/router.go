package routers

import (
	"github.com/astaxie/beego"
	"phoneLocationOpenApi/controllers"
)

func init() {
	//beego.Router("/phoneLocation", &controllers.MainController{})
	beego.Router("/v1.0/api/phoneLocation", &controllers.PlController{})
}

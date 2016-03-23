package routers

import (
	"github.com/dkkxhn/hellogo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

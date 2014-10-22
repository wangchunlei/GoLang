package routers

import (
	"github.com/wangchunlei/GoLang/website/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

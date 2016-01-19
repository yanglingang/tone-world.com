package routers

import (
	"github.com/astaxie/beego"
	"tone-world.com/controllers"
	"tone-world.com/models"
)

func init() {
	models.RegisterDB()
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/menu", &controllers.MenuController{})
	beego.AutoRouter(&controllers.GoodsController{})
	beego.AutoRouter(&controllers.WeixinController{})

	beego.Router("/token/new", &controllers.TokenController{},"post:NewToken")
	beego.Router("/token/check", &controllers.TokenController{},"post:CheckToken")
}

package controllers

import (
	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {

	c.TplNames = "index.html"
}

func (c *GoodsController) Retail() {
	c.Data["IsGoodsRetail"] = true
	c.TplNames = "goods/retail.html"
}

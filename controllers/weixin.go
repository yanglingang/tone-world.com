package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"tone-world.com/outer/weixin"
)

type WeixinController struct {
	beego.Controller
}

func (c *WeixinController) Get() {
}

func (c *WeixinController) Token() {

	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")

	sb := weixin.SignatureBody{signature, timestamp, nonce}
	passed := weixin.CheckSignature(&sb)
	if passed {
		c.Ctx.WriteString(echostr)
	} else {
		c.Ctx.WriteString("Fail")
	}
}

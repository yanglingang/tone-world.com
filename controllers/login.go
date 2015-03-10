package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/mysql"
	"tone-world.com/models"
)

var GlobalSessions *session.Manager

func init() {
	GlobalSessions, _ = session.NewManager("mysql", `{"cookieName":"sessiontoneworldcom","gclifetime":3600,"ProviderConfig":"`+beego.AppConfig.String("sessiondb::url")+`"}`)
	go GlobalSessions.GC()
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "login.html"
}

func (this *LoginController) Post() {

	isValidate, user := checkAccount(this.Ctx)
	if isValidate {
		sess, error := GlobalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		if error != nil {
			this.Redirect("/login", 301)
			return
		}
		defer sess.SessionRelease(this.Ctx.ResponseWriter)

		sess.Set("uid", user.Id)
	} else {
		this.Redirect("/login", 301)
		return
	}

	this.Redirect("/index", 301)
	return
}

func checkAccount(ctx *context.Context) (bool, *models.User) {
	uname := ctx.Request.Form["uname"][0]
	pwd := ctx.Request.Form["pwd"][0]
	user := models.ValidateUser(uname, pwd)

	return user.Id != 0, user
}

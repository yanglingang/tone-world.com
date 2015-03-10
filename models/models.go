package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int64
	Uname    string `orm:"index;size(30)"`
	Pwd      string `orm:"index;size(30)"`
	UserName string `orm:size(30)`
	Created  time.Time
}

func RegisterDB() {
	orm.RegisterModel(new(User))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql::url"))
}

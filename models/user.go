package models

import (
	"github.com/astaxie/beego/orm"
)

func ValidateUser(uname string, pwd string) *User {
	o := orm.NewOrm()
	user := &User{}
	qs := o.QueryTable("User")
	err := qs.Filter("uname", uname).Filter("pwd", pwd).One(user)
	if err != nil {
		return user
	}
	return user
}

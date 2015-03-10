package controllers

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"tone-world.com/common/mongo"
)

type MenuController struct {
	beego.Controller
}

func (this *MenuController) Get() {
	var mType string
	this.Ctx.Input.Bind(&mType, "type")

	conn := mongo.Conn()
	defer conn.Close()

	collection := conn.DB(mongo.TS).C("menu")

	var result []map[string]interface{}
	iter := collection.Find(bson.M{"name": mType}).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		return
	}

	this.Ctx.Output.Json(result, true, false)
}

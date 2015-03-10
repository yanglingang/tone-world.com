package mongo

import (
	// "github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func Conn() *mgo.Session {
	return session.Copy()
}

/*
func Close() {
	session.Close()
}
*/
var TS string

func init() {

	url := "mongodb://127.0.0.1:27017"

	// url := beego.AppConfig.String("mongodb::url")

	sess, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	session = sess
	session.SetMode(mgo.Monotonic, true)
	TS = "se3w"
	// TS = beego.AppConfig.String("mongodb::table_space_name")
}

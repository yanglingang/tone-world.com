package log

import (
	"time"
	"tone-world.com/common/mongo"
)

type OuterLog struct {
	Type     string      `json:"type"`
	Url      string      `json:"url"`
	Method   string      `json:"method"`
	Response interface{} `json:"response"`
	Time     time.Time   `json:"time"`
}

func (this *OuterLog) Log() {
	conn := mongo.Conn()
	defer conn.Close()

	c := conn.DB(mongo.TS).C("outer_log")

	err := c.Insert(this)

	if err != nil {
		panic(err)
	}
}

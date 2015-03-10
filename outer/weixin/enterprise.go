package main

import (
	// "encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"tone-world.com/common/log"
	_ "tone-world.com/common/mongo"
)

// type ErrorResponse struct {
// 	ErrCode string `json:"errcode"`
// 	ErrMsg  string `json:"errmsg"`
// }
// type OkResponse struct {
// 	Token string `json:"access_token"`
// }

func GetToken() (token string) {

	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=id&corpsecret=secrect"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// var sbody = hex.EncodeToString(body)

	var jsonObject map[string]interface{}
	err = json.Unmarshal(body, &jsonObject)
	if err != nil {
		panic(err)
	}

	logObject := log.OuterLog{Type: "gettoken", Url: url, Method: "get", Response: jsonObject, Time: time.Now()}
	logObject.Log()

	errcode, _ := jsonObject["errcode"]
	if errcode != nil {
		return
	} else {
		access_token, _ := jsonObject["access_token"]
		token = access_token.(string)
		return token
	}
}

func main() {
	fmt.Println(GetToken())
}

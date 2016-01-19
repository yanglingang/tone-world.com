package controllers

import (
	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
	// "time"
	"io/ioutil"
	"net/http"
	"fmt"
)

type TokenController struct {
	beego.Controller
}
var (
	privateKey []byte
	publicKey []byte
)
func init() {
	privateKey, _ = ioutil.ReadFile("/Users/yang/Documents/Code/Server/Go/src/github.com/dgrijalva/jwt-go/test/sample_key")
	publicKey, _ = ioutil.ReadFile("/Users/yang/Documents/Code/Server/Go/src/github.com/dgrijalva/jwt-go/test/sample_key.pub")

}
func (c *TokenController) NewToken() {
	 // Create the token
    token := jwt.New(jwt.SigningMethodRS512)
    // Set some claims
    // token.Claims["foo"] = "bar"
    // token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
    // Sign and get the complete encoded token as a string

	tokenString, err := token.SignedString(privateKey)

	if err != nil {
		panic(err.Error())
	}
    c.Ctx.WriteString(tokenString)
}

func (c *TokenController) CheckToken() {
	token, err := jwt.ParseFromRequest(c.Ctx.Input.Request, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return publicKey, nil
		}
    })

    if err == nil && token.Valid {
        c.Ctx.WriteString("ok")
    } else {
    	fmt.Println(err)
        c.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)

    }
}
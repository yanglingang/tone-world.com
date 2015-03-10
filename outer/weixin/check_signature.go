package weixin

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
)

var token string = "ofMA_tz8JXR_Grf6Rn3A5x6kClCk"

type SignatureBody struct {
	Signature string
	Timestamp string
	Nonce     string
}

func CheckSignature(body *SignatureBody) bool {
	strs := []string{token, body.Timestamp, body.Nonce}
	sort.Strings(strs)
	var str string

	for _, value := range strs {
		str += value
	}
	fmt.Println(str)
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	str = hex.EncodeToString(bs)
	fmt.Println(str)
	return str == body.Signature
}
func main() {
	sb := SignatureBody{"a6c0b1fc0c57d12fa258d763a551e64408a4f3e7", "d", "c"}

	fmt.Println(CheckSignature(&sb))
}

package base

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"crypto/sha1"
	"io"
)

const TOKEN  = "fshare"

type Base struct {
	token string
	Resp http.ResponseWriter
	Req *http.Request
}


func (t *Base) Valid(echostr string) {
	fmt.Println(echostr)
	if (t.checkSignature()) {
		fmt.Println("aaaa")
		fmt.Fprintf(t.Resp, echostr)
	}
}

func (t *Base) checkSignature() bool {
	t.token = TOKEN

	timestamp, nonce, signature := t.Req.Form["timestamp"][0], t.Req.Form["nonce"][0], t.Req.Form["signature"][0]
	var tmpArr = []string{t.token, timestamp, nonce}

	fmt.Println(t.Req.Form)
	sort.Strings(tmpArr)

	tmpStr := strings.Join(tmpArr, " ")
	fmt.Println(tmpArr, tmpStr)

	ts := sha1.New()
	io.WriteString(ts, tmpStr)
	tmpStr = fmt.Sprintf("%x",ts.Sum(nil))

	fmt.Println(tmpStr)

	if (signature == tmpStr) {
		return true
	}

	return false
}
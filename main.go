package main

import (
	"net/http"
	"github.com/skiy/express-for-go/base"
	//"fmt"
)

func ReceiveMsg(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()

	bs := &base.Base{}
	bs.Req = r
	bs.Resp = w

	if r.Form["echostr"] != nil {
		bs.Valid(r.Form["echostr"][0])
	} else {
		//fmt.Println(r.Form["echostr"])
		bs.ResponseMsg()
	}
}

func main() {

	http.HandleFunc("/", ReceiveMsg)
	http.ListenAndServe(":9003", nil)
}

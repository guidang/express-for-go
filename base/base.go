package base

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"crypto/sha1"
	"io"
	"encoding/xml"
	"time"
	"io/ioutil"
	"log"
)

const TOKEN  = "fshare"

type Base struct {
	token string
	Resp http.ResponseWriter
	Req *http.Request
}

type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}


func (t *Base) Valid(echostr string) {
	//fmt.Println(echostr)
	if t.checkSignature() {
		//fmt.Println("aaaa")
		fmt.Fprintf(t.Resp, echostr)
	}
}

func (t *Base) checkSignature() bool {
	t.token = TOKEN

	timestamp, nonce, signature := t.Req.Form["timestamp"][0], t.Req.Form["nonce"][0], t.Req.Form["signature"][0]
	var tmpArr = []string{t.token, timestamp, nonce}

	//fmt.Println(t.Req.Form)
	sort.Strings(tmpArr)

	/*
	//高效做法
	buffer := bytes.Buffer
	for _, v := range tmpArr {
		buffer.WriteString(v)
	}
	tmpStr := buffer.String()
	fmt.Println(tmpStr)
	 */
	tmpStr := strings.Join(tmpArr, " ")
	//fmt.Println(tmpArr, tmpStr)

	ts := sha1.New()
	io.WriteString(ts, tmpStr)
	tmpStr = fmt.Sprintf("%x",ts.Sum(nil))

	//fmt.Println(tmpStr)

	if signature == tmpStr {
		return true
	}

	return false
}

func (t *Base) ResponseMsg() {
	fmt.Fprintf(t.Resp, string(t.parseTextRequest()))
}

/* 接收的数据 xml 格式 */
func (t *Base) parseTextRequest() *TextRequestBody {
	body, err := ioutil.ReadAll(t.Req.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println(string(body))
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}
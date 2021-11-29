package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/imroc/req"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime"
)

func stack() []byte {
	buf := make([]byte, 9128)
	n := runtime.Stack(buf, false)
	return buf[:n]
}

func GetStack() string {
	return fmt.Sprintf("%s", stack())
}

func JsonToString(obj interface{}) string {
	info, err := jsoniter.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(info)
}

func DumpRequest(request *http.Request, response *http.Response, section string) {
	println(fmt.Sprintf("%v start", section))
	body, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	println(string(body))
	println("===========================")
	body, err = httputil.DumpResponse(response, true)
	if err != nil {
		println(fmt.Sprintf("httputil.DumpResponse err: %v", err))
	} else {
		println(string(body))
	}
	println(fmt.Sprintf("%v end", section))
}

func DumpResp(resp *req.Resp, section string) {
	println(fmt.Sprintf("%v start", section))
	println(resp.Dump())
	println(fmt.Sprintf("%v end", section))
}

func GetServiceName() string {
	_, fileName, _, ok := runtime.Caller(1)
	if ok {
		return fileName
	} else {
		return ""
	}
}

func GetFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		return runtime.FuncForPC(pc).Name()
	} else {
		return ""
	}
}

func GetCallerFuncName() string {
	pc, _, _, ok := runtime.Caller(2)
	if ok {
		return runtime.FuncForPC(pc).Name()
	} else {
		return ""
	}
}

func ReceiveStr(str string) {
	var strBuffer bytes.Buffer
	_ = json.Indent(&strBuffer, []byte(str), "", "    ")
	fmt.Println(strBuffer.String())
}

func ReceiveStruct(t interface{}) {
	b, err := jsoniter.Marshal(t)
	if err != nil {
		log.Fatalln(err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")

	if err != nil {
		log.Fatalln(err)
	}

	out.WriteTo(os.Stdout)
}

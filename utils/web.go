package utils

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
)

func DoRemoteRequest(req *http.Request) {
	//调用接口
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	DumpRequest(req, resp, "section list")
	var body []byte
	if resp.Header.Get("Content-Encoding") == "gzip" {
		println("--------------------gzip")
		res := new(bytes.Buffer)
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(res, gr)
		if err != nil {
			panic(err)
		}
		body = res.Bytes()
	} else {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
	}

	println(string(body))
}

package test

import (
	"GoWild/route"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_MsgProduce(t *testing.T) {
	w := httptest.NewRecorder()
	values := url.Values{}
	values.Add("msg", "first msg hello")

	queryString := values.Encode()
	queryString, _ = url.QueryUnescape(queryString)
	request, err := http.NewRequest("PUT", "/nsqConsumer/genMsg?"+values.Encode(), nil)
	if err != nil {
		panic(err)
	}

	route.Route().ServeHTTP(w, request)
	println(fmt.Sprintf("%v", w))
}

func Test_Hello(t *testing.T) {
	w := httptest.NewRecorder()
	values := url.Values{}
	values.Add("msg", "first msg hello")

	queryString := values.Encode()
	queryString, _ = url.QueryUnescape(queryString)
	request, err := http.NewRequest("GET", "/test/hello?"+values.Encode(), nil)
	if err != nil {
		panic(err)
	}

	route.Route().ServeHTTP(w, request)
	println(fmt.Sprintf("%v", w))
}

package http

import (
	"GoWild/utils"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func Test_ActorSongList(t *testing.T) {
	values := url.Values{}

	queryString := values.Encode()
	queryString, _ = url.QueryUnescape(queryString)
	request, err := http.NewRequest("POST", Local+"/nsqConsumer/stopConsumer?"+values.Encode(), nil)
	if err != nil {
		log.Fatal(err)
	}
	//调用远端方法
	utils.DoRemoteRequest(request)

}

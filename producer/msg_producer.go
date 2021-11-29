package producer

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

var msgProducer *nsq.Producer
var msgProducerOnce sync.Once

func GetMsgProducer() (pro *nsq.Producer, err error) {
	onceBody := func() {
		config := nsq.NewConfig()
		msgProducer, err = nsq.NewProducer("10.171.4.4:4150", config)
		if err != nil {
			log.Default().Println("连接失败...")
			panic(err)
		}
	}
	msgProducerOnce.Do(onceBody)
	return msgProducer, nil
}

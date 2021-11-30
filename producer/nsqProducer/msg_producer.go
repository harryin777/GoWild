package nsqProducer

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

var msgProducer *nsq.Producer
var msgProducerOnce sync.Once

func getMsgProducer() (pro *nsq.Producer, err error) {
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

func NsqProduce(msg interface{}) error {
	mp, err := getMsgProducer()
	if err != nil {
		panic(err)
	}
	err = mp.Ping()
	if err != nil {
		log.Default().Println("ping 不通...")
		return err
	}

	b, err := json.Marshal(msg)
	if err != nil {
		log.Default().Println("marshal fail ...")
		return err
	}
	err = mp.Publish("test", b)
	if err != nil {
		log.Default().Println("publish fail ...")
		return err
	}
	//生产者停止运行
	mp.Stop()
	return nil
}

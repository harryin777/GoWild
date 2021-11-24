package nsqService

import (
	"GoWild/producer"
	"encoding/json"
	"log"
)

func ProducerMsg(msg string) {
	mp, err := producer.GetMsgProducer()
	if err != nil {
		panic(err)
	}
	err = mp.Ping()
	if err != nil {
		log.Default().Println("ping 不通...")
		panic(err)
	}

	b, err := json.Marshal(msg)
	if err != nil {
		log.Default().Println("marshal fail ...")
	}
	err = mp.Publish("test", b)
	if err != nil {
		log.Default().Println("publish fail ...")
	}
	//生产者停止运行
	mp.Stop()
}

func ConsumeMsg(channel string) {

}

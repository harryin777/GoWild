package nsqConsumer

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
	"time"
)

var nsqConsumerOnce sync.Once
var NsqConsumer *nsq.Consumer

func init() {

}

func MsgConsumer() {
	onceBody := func() {
		config := nsq.NewConfig()
		//设置断开重连时间
		config.LookupdPollTimeout = time.Second
		var err error
		NsqConsumer, err = nsq.NewConsumer("test", "test_Chan", config)
		if err != nil {
			log.Default().Println("创建消费者失败...")
			panic(err)
		}
		NsqConsumer.SetLoggerLevel(nsq.LogLevelError)
		NsqConsumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
			var data string
			err = jsoniter.Unmarshal(message.Body, &data)
			if err != nil {
				panic(err)
			}
			fmt.Printf("消费!!!!%v \n", data)
			return nil
		}))

		lookupAddr := []string{
			"10.171.4.4:4161",
		}
		err = NsqConsumer.ConnectToNSQLookupds(lookupAddr)
		if err != nil {
			log.Default().Println("ConnectToNSQLookupds failed...")
			panic(err)
		}

		// 4. 接收消费者停止通知
		<-NsqConsumer.StopChan

		// 5. 获取统计结果
		stats := NsqConsumer.Stats()
		fmt.Sprintf("message received %d, finished %d, requeued:%s, connections:%s",
			stats.MessagesReceived, stats.MessagesFinished, stats.MessagesRequeued, stats.Connections)
	}

	nsqConsumerOnce.Do(onceBody)

}

package consumer

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

func init() {

}

func MsgConsumer() {
	config := nsq.NewConfig()
	c, err := nsq.NewConsumer("test", "test_Chan", config)
	if err != nil {
		log.Default().Println("创建消费者失败...")
		panic(err)
	}

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Printf("消费!!!!%v \n", message.Body)
		return nil
	}))

	lookupAddr := []string{
		"10.171.4.4:4161",
	}
	err = c.ConnectToNSQLookupds(lookupAddr)
	if err != nil {
		log.Default().Println("ConnectToNSQLookupds failed...")
		panic(err)
	}

	// 4. 接收消费者停止通知
	<-c.StopChan

	// 5. 获取统计结果
	stats := c.Stats()
	fmt.Sprintf("message received %d, finished %d, requeued:%s, connections:%s",
		stats.MessagesReceived, stats.MessagesFinished, stats.MessagesRequeued, stats.Connections)

}

package nsqService

import (
	"GoWild/consumer/nsqConsumer"
	"GoWild/producer/nsqProducer"
)

/**
 * @Description 生产消息
 * @Param
 * @return
 **/
func ProducerMsg(msg string) {
	err := nsqProducer.NsqProduce(msg)
	if err != nil {
		panic(err)
	}
}

/**
 * @Description 停止消费
 * @Param
 * @return
 **/
func StopConsumer() {
	nsqConsumer.NsqConsumer.StopChan <- 1
}

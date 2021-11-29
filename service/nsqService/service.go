package nsqService

import (
	"GoWild/producer/nsq"
)

/**
 * @Description 生产消息
 * @Param
 * @return
 **/
func ProducerMsg(msg string) {
	err := nsq.NsqProduce(msg)
	if err != nil {
		panic(err)
	}
}

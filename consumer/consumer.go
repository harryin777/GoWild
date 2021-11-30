package consumer

import "GoWild/consumer/nsqConsumer"

func StartAllConsumer() {
	go nsqConsumer.MsgConsumer()
}

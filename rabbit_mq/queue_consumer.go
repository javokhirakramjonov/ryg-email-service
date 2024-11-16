package rabbit_mq

import (
	ampq "github.com/rabbitmq/amqp091-go"
	"log"
)

type QueueConsumer interface {
	Consume()
}

type BaseQueueConsumer struct {
	Ch *ampq.Channel
}

func (bqc *BaseQueueConsumer) Consume() {
	log.Print("This is the base consumer")
}

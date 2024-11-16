package rabbit_mq

import (
	"fmt"
	ampq "github.com/rabbitmq/amqp091-go"
	"log"
)

const (
	genericEmailQueueName  = "generic_email"
	genericEmailRoutingKey = genericEmailQueueName
)

type GenericEmailQueueConsumer struct {
	BaseQueueConsumer
}

func NewGenericEmailQueueConsumer(ch *ampq.Channel, exchangeName string) QueueConsumer {
	bindQueue(ch, exchangeName)

	return &GenericEmailQueueConsumer{
		BaseQueueConsumer{
			Ch: ch,
		},
	}
}

func bindQueue(ch *ampq.Channel, exchangeName string) {
	q, err := ch.QueueDeclare(
		genericEmailQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,
		"generic_email",
		exchangeName,
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")
}

func (c *GenericEmailQueueConsumer) Consume() {
	msgs, err := c.Ch.Consume(
		genericEmailQueueName,
		genericEmailRoutingKey,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, fmt.Sprintf("Failed to register a consumer for queue: %s", genericEmailQueueName))

	go func() {
		for msg := range msgs {
			log.Printf("Generic Email Consumer Received a message: %s", msg.Body)
		}
	}()
}

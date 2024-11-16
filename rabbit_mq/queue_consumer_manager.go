package rabbit_mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

const exchangeName = "email_service_topics"

type QueueConsumerManager struct {
	conn      *amqp.Connection
	ch        *amqp.Channel
	consumers []QueueConsumer
}

func NewQueueConsumerManager() QueueConsumerManager {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	log.Printf("Connected to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	log.Printf("Opened a channel")

	genericEmailQueue := NewGenericEmailQueueConsumer(ch, exchangeName)

	return QueueConsumerManager{
		conn:      conn,
		ch:        ch,
		consumers: []QueueConsumer{genericEmailQueue},
	}
}

func (qcm *QueueConsumerManager) Start() {
	for _, consumer := range qcm.consumers {
		consumer.Consume()
	}

	log.Printf("Consumers started")

	// Block forever
	<-make(chan bool)
}

func (qcm *QueueConsumerManager) Close() {
	err := qcm.ch.Close()
	failOnError(err, "Failed to close channel")

	err = qcm.conn.Close()
	failOnError(err, "Failed to close connection")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

package rabbit_mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func ConnectAMQ() (*amqp.Connection, *amqp.Channel, amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	log.Printf("Connected to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	log.Printf("Opened a channel")

	q, err := ch.QueueDeclare(
		"simple_message", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return conn, ch, q
}

func ConsumeForever(ch *amqp.Channel, q *amqp.Queue) {
	msgs, err := ch.Consume(
		q.Name,          // queue
		"email_service", // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	failOnError(err, "Failed to register a consumer")
	log.Printf("Registered a consumer(email_service)")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever // Block forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

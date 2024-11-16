package rabbit_mq

import (
	"fmt"
	ampq "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"log"
	"net/smtp"
	"ryg-email-service/conf"
	"ryg-email-service/gen_proto/email_service"
)

const (
	genericEmailQueueName  = "generic_email"
	genericEmailRoutingKey = genericEmailQueueName
)

type GenericEmailQueueConsumer struct {
	emailConfig struct {
		host     string
		port     string
		username string
	}
	auth smtp.Auth
	BaseQueueConsumer
}

func NewGenericEmailQueueConsumer(cnf *conf.Config, ch *ampq.Channel, exchangeName string) QueueConsumer {
	bindQueue(ch, exchangeName)

	return &GenericEmailQueueConsumer{
		emailConfig: struct {
			host     string
			port     string
			username string
		}{
			host:     cnf.EmailHost,
			port:     cnf.EmailPort,
			username: cnf.EmailUsername,
		},
		auth: smtp.PlainAuth("", cnf.EmailUsername, cnf.EmailPassword, cnf.EmailHost),
		BaseQueueConsumer: BaseQueueConsumer{
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
			var ge email_service.GenericEmail
			err := proto.Unmarshal(msg.Body, &ge)
			if err != nil {
				log.Printf("Failed to unmarshal email message: %v", err)
			}

			log.Printf("\n To: %s\n Subject: %s\n Body: %s\n", ge.To, ge.Subject, ge.Body)

			err = c.sendEmail(ge.To, ge.Subject, ge.Body)
			if err != nil {
				log.Printf("Failed to send email: %v", err)
			} else {
				log.Printf("Email sent successfully")
			}
		}
	}()
}

func (c *GenericEmailQueueConsumer) sendEmail(to, subject, body string) error {
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(c.emailConfig.host+":"+c.emailConfig.port, c.auth, c.emailConfig.username, []string{to}, msg)
	if err != nil {
		return err
	}
	return nil
}

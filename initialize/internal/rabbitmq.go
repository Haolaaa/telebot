package internal

import (
	"log"
	"telebot/global"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type RabbitMQ struct {
	Conn      *amqp.Connection
	Channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Mqurl     string
}

func (r RabbitMQ) FailOnError(err error, str string) {
	if err != nil {
		global.LOG.Error("str", zap.Error(err))
	}
}

func (r *RabbitMQ) PublishPub(message string) {
	err := r.Channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.FailOnError(err, "Failed to declare an exchange")

	err = r.Channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	r.FailOnError(err, "Failed to publish a message")
}

func (r *RabbitMQ) ReceiveSub(keyStr string) {
	err := r.Channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.FailOnError(err, "Failed to declare an exchange")

	q, err := r.Channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.FailOnError(err, "Failed to declare a queue")

	err = r.Channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
	)

	message, err := r.Channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range message {
			log.Printf(keyStr+"Received a message: %s", d.Body)
		}
	}()
	log.Println("退出请按 CTRL+C")
	<-forever
}

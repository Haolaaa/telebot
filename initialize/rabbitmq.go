package initialize

import (
	"telebot/initialize/internal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQ(mqUrl, queueName, exchange, key string) *internal.RabbitMQ {
	rabbitmq := &internal.RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     mqUrl,
	}
	var err error
	rabbitmq.Conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnError(err, "创建链接错误")
	rabbitmq.Channel, err = rabbitmq.Conn.Channel()
	rabbitmq.FailOnError(err, "获取channel失败")

	return rabbitmq
}

func NewRabbitMQPubSub(mqURL, exchangeName string) *internal.RabbitMQ {
	rabbitmq := NewRabbitMQ(mqURL, "", exchangeName, "")
	var err error
	rabbitmq.Conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnError(err, "创建链接错误")
	rabbitmq.Channel, err = rabbitmq.Conn.Channel()
	rabbitmq.FailOnError(err, "获取channel失败")

	return rabbitmq
}

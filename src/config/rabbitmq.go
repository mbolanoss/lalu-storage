package config

import (
	"fmt"

	"github.com/streadway/amqp"
)

func RabbitMQConn() *amqp.Channel{
	conn, err := amqp.Dial("amqp://lalu_mq:lalu_mq@localhost:5672/")

	if err != nil {
		panic("Error conecting to rabbitmq")
	}

	fmt.Println("Succesfully connected to RabbitMQ")

	channel, err := conn.Channel()

	if err != nil {
		panic("Error creating channel")
	}

	return channel
}
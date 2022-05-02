package config

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func RabbitMQConn() *amqp.Channel{

	user := os.Getenv("MQ_USER")
	password := os.Getenv("MQ_PASSWORD")
	host := os.Getenv("MQ_HOST")

	rabbitMQUrl := fmt.Sprintf("amqp://%s:%s@%s:5672/", user, password, host)

	conn, err := amqp.Dial(rabbitMQUrl)

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
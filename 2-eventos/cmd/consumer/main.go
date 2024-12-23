package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"pos-go-expert/2-eventos/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, "orders")
	for msg := range msgs {
		println(string(msg.Body))
		msg.Ack(false)
	}
}

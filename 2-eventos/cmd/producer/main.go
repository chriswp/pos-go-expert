package main

import "pos-go-expert/2-eventos/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = rabbitmq.Publish(ch, []byte("Hello, World!"), "amq.direct")
	if err != nil {
		return
	}
}

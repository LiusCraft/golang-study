package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"liuscraft.op/study/common/utils"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:123456@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ Server")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	helloQueue, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		helloQueue.Name, // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	go func() {
		for {
			ch.PublishWithContext(context.TODO(), "", helloQueue.Name, false, false, amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("Hello, world!"),
			})
			time.Sleep(1 * time.Second)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

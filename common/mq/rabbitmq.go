package main

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {
	conn, err := amqp.Dial("amqp://mengwei:mengwei@192.168.0.106:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello1", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	body := "Hello World!"
	go func(){
		for ;; {
			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType: "text/plain",
					Body:        []byte(body),
				})
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	for ;; {
		time.Sleep(time.Second)
	}
}

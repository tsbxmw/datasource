package main

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {
	conn,err := amqp.Dial("amqp://mengwei:mengwei@192.168.0.106:5672/")
	FailError1(err,"receive:Failed to connect to RabbitMQ")
	defer conn.Close()
	ch,err := conn.Channel()
	FailError1(err,"receive:Failed to open a channel")
	defer ch.Close()
	q,err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	err = ch.Qos(
		1, //// 在没有返回ack之前，最多只接收1个消息
		0,
		false,
	)
	FailError1(err,"Failed to set Qos")
	msgs,err := ch.Consume(
		q.Name,
		"",
		false,//将autoAck设置为false，则需要在消费者每次消费完成
		// 消息的时候调用d.Ack(false)来告诉RabbitMQ该消息已经消费
		false,
		false,
		false,
		nil,
	)
	FailError1(err,"Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs{
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			fmt.Println()
			time.Sleep(t * time.Second)
			log.Printf("Done")
			//multiple为true的时候：此次交付和之前没有确认的交付都会在通过同一个通道交付，这在批量处理的时候很有用
			//为false的时候只交付本次。只有该方法执行了，RabbitMQ收到该确认才会将消息删除
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
func FailError1(err error,msg string)  {
	if err != nil {
		log.Fatal("%s : %s",msg,err)
	}
}

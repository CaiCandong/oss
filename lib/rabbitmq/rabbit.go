package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	channel  *amqp.Channel
	Name     string
	exchange string
}

//New函数用于创建一个新的rabbitmq.RabbitMQ结构体，
//该结构体的Bind方法可以将自己的消息队列和一个exchange绑定，
//所有发往该exchange的消息都能在自己的消息队列中被接收到。
func New(s string) *RabbitMQ {
	conn, err := amqp.Dial(s)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	q, e := ch.QueueDeclare(
		"",
		false,
		true,
		false,
		false,
		nil,
	)
	if e != nil {
		panic(e)
	}
	mq := new(RabbitMQ)
	mq.channel = ch
	mq.Name = q.Name
	return mq
}

func (q *RabbitMQ) Bind(exchange string) {
	e := q.channel.QueueBind(
		q.Name,
		"",
		exchange,
		false,
		nil)
	if e != nil {
		panic(e)
	}
	q.exchange = exchange
}

//Send方法可以往某个消息队列发送消息。
func (q *RabbitMQ) Send(queue string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}
	e = q.channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    str,
		})
	if e != nil {
		panic(e)
	}
}

//Publish方法可以往某个exchange发送消息。
func (q *RabbitMQ) Publish(exchange string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}
	e = q.channel.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    str,
		})
	if e != nil {
		panic(e)
	}
}

//Consume方法用于生成一个接收消息的go channel，使客户程序可以通过Go语言的原生机制接收队列中的消息。
func (q *RabbitMQ) Consume() <-chan amqp.Delivery {
	c, e := q.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if e != nil {
		panic(e)
	}
	return c
}

//Close方法用于关闭消息队列
func (q *RabbitMQ) Close() {
	q.channel.Close()
}

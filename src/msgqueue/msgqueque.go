package msgqueue

import (
	"log"
	"github.com/streadway/amqp"
	"fmt"
	"flag"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func Publish() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"pdna", // name
		true, // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil, // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "LA DEE DA, THIS IS OUR SUPER COOL MESSAGE"
	err = ch.Publish(
		"", // exchange
		q.Name, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

}
func Producer(amqpURI string, msg string) {

	connection, _ := amqp.Dial(amqpURI)

	channel, _ := connection.Channel()

	for i := 0; i < 10; i++ {
		channel.Publish(
			"", //exchange
			"pdna", //routingKey
			false,
			false,
			amqp.Publishing{
				Headers: amqp.Table{},
				ContentType: "text/plain",
				ContentEncoding: "UTF-8",
				Body: []byte(msg + fmt.Sprint(i)),
				DeliveryMode: amqp.Transient,
				Priority: 0,
			},
		)
	}
}
type Consumerx struct {
	conn	*amqp.Connection
	channel	*amqp.Channel
	tag	string
	done	chan error
}
var (
	uri		= flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchange	= flag.String("exchange", "logs", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType	= flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	queue		= flag.String("queue", "logs", "Ephemeral AMQP queue name")
	bindingKey	= flag.String("key", "", "AMQP binding key")
	consumerTag	= flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
	wsPort		= flag.Int("ws-port", 23456, "WebSockets port to listen to")
)
func Consumer(amqpURI string, msg string)   {
	c := &Consumerx{
		conn:		nil,
		channel:	nil,
		tag:		"",
		done:		make(chan error),
	}
	c.conn, _ = amqp.Dial(amqpURI)

	c.channel, _ = c.conn.Channel()

	deliveries, err := c.channel.Consume(
		"pdna",	// name
		c.tag,	// consumerTag,
		false,	// noAck
		false,	// exclusive
		false,	// noLocal
		false,	// noWait
		nil,	// arguments
	)
	if err != nil {
		  fmt.Errorf("Queue Consume: %s", err)
	}

	for d := range deliveries {
		log.Printf("got %s", d.Body)
	}

}

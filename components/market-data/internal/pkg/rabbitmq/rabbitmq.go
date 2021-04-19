package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

var conn *amqp.Connection
var ch *amqp.Channel
var q amqp.Queue

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connect() {
	n, err := amqp.Dial("amqp://user:password@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	conn = n
}

func openChannel() {
	n, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	ch = n
}

func declareQueue() {
	n, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	q = n
}

func init() {
	connect()
	//defer conn.Close()

	openChannel()
	//defer ch.Close()

	declareQueue()
}

// Enqueue will just enqueue a message in the queue
func Enqueue(body []byte) {
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

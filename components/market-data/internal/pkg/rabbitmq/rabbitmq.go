package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Queue is the main struct that will hold the queue details + the connection and channel info so that
// we may invoke the close handlers
type Queue struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	name       string
}

// Teardown will terminate the connections to rabbitmq
func (q Queue) Teardown() {
	q.connection.Close()
	q.channel.Close()
}

// Enqueue will just enqueue a message in the queue
func (q Queue) Enqueue(body []byte) {
	err := q.channel.Publish(
		"",     // exchange
		q.name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

// Create handles the queue creation. It stores the data in the Queue struct
func (q *Queue) Create(name string) (*Queue, error) {
	connection, err := connect()
	channel, err := openChannel(connection)
	q.connection = connection
	q.channel = channel
	q.name = name
	declareQueue(channel, name)
	return q, err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connect() (*amqp.Connection, error) {
	n, err := amqp.Dial("amqp://user:password@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	return n, err
}

func openChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	n, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return n, err
}

func declareQueue(ch *amqp.Channel, name string) error {
	_, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	msg := fmt.Sprintf("Failed to declare a queue named %s", name)
	failOnError(err, msg)
	return err
}

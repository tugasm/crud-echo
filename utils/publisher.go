package utils

import (
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// Message is a struct containing the name and payload of a message queue item's Body
type Message struct {
	Type    string
	Payload interface{}
}

// Connect returns an AMQP channel connected to the AMQP server
func Connect() *amqp.Connection {
	godotenv.Load()

	host := os.Getenv("MQ_HOST")
	port := os.Getenv("MQ_PORT")
	url := "amqp://guest:guest@" + host + ":" + port + "/"

	conn, err := amqp.Dial(url)
	FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

// Produce adds a message into the rabbitMQ queue
func Produce(m Message) {
	conn := Connect()
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"customer",
		false,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")

	b := &m.Payload
	body, err := json.Marshal(b)
	FailOnError(err, "Failed to read message body")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Type:        m.Type,
			Body:        body,
		},
	)
	if err != nil {
		logHandler := logrus.WithField("err", err)
		logHandler.Error("Failed to publish a message")
	}
	logrus.Info(`Successfully published message with topics :`, m.Type)
}

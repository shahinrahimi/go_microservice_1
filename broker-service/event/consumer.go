package event

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	authServiceURL = "http://auth-service:3002/auth"
	logServiceURL  = "http://logger-service:3003/log"
	mailServiceURL = "http://mailer-service:3004/send"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	log.Printf("creating consumer with conn: %v\n", conn)
	consumer := Consumer{
		conn: conn,
	}
	log.Printf("setting up consumer\n")
	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}
	return consumer, nil
}

func (c *Consumer) setup() error {
	log.Printf("setting up consumer setup")
	channel, err := c.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (c *Consumer) Listen(topics []string) error {
	log.Printf("listening to topics: %v\n", topics)
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	log.Printf("binding queue [%s] to exchange [logs_topic]\n", q.Name)
	for _, s := range topics {
		err = ch.QueueBind(
			q.Name,
			s,
			"logs_topic",
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}

	log.Printf("consume messages from queue [%s]\n", q.Name)
	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	log.Printf("forever read messages from queue [%s]\n", q.Name)
	forever := make(chan bool)
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)

			go handlePayload(payload)
		}
	}()

	log.Printf("waiting for message [Exchange, Queue] [logs_topic, %s]\n", q.Name)

	<-forever

	return nil
}

func handlePayload(payload Payload) {
	fmt.Println("handled payload:", payload)
	switch payload.Name {
	case "log":
		// write to database
		err := logEvent(payload)
		if err != nil {
			fmt.Println(err)
		}
	case "event":
		// publish event
		err := logEvent(payload)
		if err != nil {
			fmt.Println(err)
		}
	case "auth":

	default:
		err := logEvent(payload)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func logEvent(entry Payload) error {
	jsonData, err := json.MarshalIndent(entry, "", "\t")
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return errors.New("error calling logger service")
	}
	return nil

}

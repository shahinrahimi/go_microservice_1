package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	conn *amqp.Connection
}

func (e *Emitter) setup() error {
	log.Println("setting up event emitter setup")
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	return declareExchange(channel)
}

func (e *Emitter) Push(event string, severity string) error {
	log.Printf("Grabbing channel")
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	log.Println("Publishing message")

	err = channel.Publish(
		"logs_topic",
		severity,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func NewEventEmitter(conn *amqp.Connection) (Emitter, error) {
	log.Printf("creating event emitter with conn: %v\n", conn)
	emitter := Emitter{
		conn: conn,
	}
	log.Printf("setting up event emitter\n")
	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}
	return emitter, nil
}

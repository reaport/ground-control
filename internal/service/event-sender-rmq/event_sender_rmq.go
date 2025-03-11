package eventsenderrmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	URL   string `mapstructure:"url"`
	Queue string `mapstructure:"queue"`
}

type EventSenderRMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    *amqp.Queue
}

func New(cfg *Config) (*EventSenderRMQ, error) {
	conn, err := amqp.Dial(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	queue, err := ch.QueueDeclare(
		cfg.Queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %w", err)
	}

	return &EventSenderRMQ{
		conn: conn,
		ch:   ch,
		q:    &queue,
	}, nil
}

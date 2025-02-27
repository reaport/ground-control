package eventsenderrmq

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/reaport/ground-control/internal/entity"
)

func (es *EventSenderRMQ) SendEvent(ctx context.Context, event *entity.Event) error {
	message, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = es.ch.PublishWithContext(
		ctx,
		"",
		es.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	return nil
}

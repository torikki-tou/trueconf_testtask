package repo

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/torikki-tou/trueconf_testtask/config"
	"github.com/torikki-tou/trueconf_testtask/dto"
)

type RabbitRepository interface {
	ProduceMessage(notification dto.Notification) error
}

type rabbitRepository struct {
	connection *amqp.Connection
}

func NewRabbitRepository(connection *amqp.Connection) RabbitRepository {
	return &rabbitRepository{connection: connection}
}

func (r *rabbitRepository) ProduceMessage(notification dto.Notification) error {
	ch, err := r.connection.Channel()
	if err != nil {
		return err
	}
	defer func(ch *amqp.Channel) { _ = ch.Close() }(ch)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	err = ch.PublishWithContext(ctx,
		"",
		config.QueueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		return err
	}
	return nil
}
package repo

import (
	"context"
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
	"github.com/torikki-tou/trueconf_testtask/config"
)

type QueueRepository interface {
}

type queueRepository struct {
	connection *ampq.Connection
}

func NewQueueRepository(connection *ampq.Connection) QueueRepository {
	return &queueRepository{connection: connection}
}


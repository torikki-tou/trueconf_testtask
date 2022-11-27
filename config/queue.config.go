package config

import amqp "github.com/rabbitmq/amqp091-go"

var QueueName = "notifications"

func InitQueue(con *amqp.Connection) {
	ch, err := con.Channel()
	if err != nil {
		panic(err)
	}
	defer func(ch *amqp.Channel) { _ = ch.Close() }(ch)

	_, err = ch.QueueDeclare(
		QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
}

func SetupRabbitMQConnection() *amqp.Connection {
	con, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")
	if err != nil {
		panic(err)
	}
	return con
}

func CloseRabbitMQConnection(con *amqp.Connection) {
	_ = con.Close()
}

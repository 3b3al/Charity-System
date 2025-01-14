package queue

import (
	"errors"

	"github.com/3b3al/Charity-System/setting"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect(q setting.Queue) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(q.MQAddress)
	if err != nil {
		return nil, nil, errors.New("could not connect to RabbitMQ: " + err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, errors.New("could not open channel: " + err.Error())
	}

	err = ch.ExchangeDeclare(
		q.ExchangeName,
		q.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, errors.New("could not declare an exchange: " + err.Error())
	}

	_, err = ch.QueueDeclare(
		q.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)

	for _, routeKey := range q.Bindings {
		err = ch.QueueBind(
			q.QueueName,
			routeKey,
			q.ExchangeName,
			false,
			nil)
		if err != nil {
			return nil, nil, errors.New("could not bind a queue: " + err.Error())
		}
	}

	return conn, ch, nil
}

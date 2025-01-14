package queue

import (
	"context"
	"errors"

	"github.com/3b3al/Charity-System/core"
	amqp "github.com/rabbitmq/amqp091-go"
)

type mq struct {
	conn         *amqp.Connection
	ch           *amqp.Channel
	exchangeName string
	queueName    string
}

func New(
	conn *amqp.Connection,
	ch *amqp.Channel,
	exchangeName string,
	queueName string,
) core.QueueAPI {
	return &mq{conn, ch, exchangeName, queueName}
}

func (q *mq) Consume() (*amqp.Connection, <-chan amqp.Delivery, error) {
	conn := q.conn
	ch := q.ch

	err := ch.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		return nil, nil, errors.New("could not set QoS: " + err.Error())
	}

	deliveries, err := ch.Consume(
		q.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, errors.New("could not consume: " + err.Error())
	}

	return conn, deliveries, nil
}

func (q *mq) Publish(msg []byte, routeKey string) error {
	ch := q.ch

	err := ch.PublishWithContext(
		context.Background(),
		q.exchangeName,
		routeKey,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         msg,
		})
	if err != nil {
		return errors.New("could not publish a message: " + err.Error())
	}

	return nil
}

package core

import amqp "github.com/rabbitmq/amqp091-go"

type QueueAPI interface {
	Consume() (*amqp.Connection, <-chan amqp.Delivery, error)
	Publish(msg []byte, routeKey string) error
}

type QueueDetails struct {
	ExchangeName string           `json:"exchangeName"`
	ExchangeType string           `json:"exchangeType"`
	BindKey      string           `json:"bindingKey"`
	QueueName    string           `json:"queueName"`
	Message      []byte           `json:"message"`
	Connection   *amqp.Connection `json:"connection"`
	Channel      *amqp.Channel    `json:"channel"`
	RouteKey     string           `json:"routeKey"`
	RouteKeys    []string         `json:"routeKeys"`
	MQAddress    string           `json:"mqAdress"`
}

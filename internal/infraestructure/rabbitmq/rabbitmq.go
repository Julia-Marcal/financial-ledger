package rabbitmq

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

const (
	maxRetries    = 3
	retryInterval = 500 * time.Millisecond
)

var conn *amqp091.Connection
var channel *amqp091.Channel

func InitRabbitMQ() error {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@localhost:5672/"
	}
	var err error
	conn, err = amqp091.Dial(url)
	if err != nil {
		return err
	}
	channel, err = conn.Channel()
	if err != nil {
		return err
	}

	err = channel.Confirm(false)
	if err != nil {
		return err
	}

	return nil
}

func PublishTransaction(exchange string, routingKey string, body []byte) error {
	if channel == nil {
		return amqp091.ErrClosed
	}

	var lastErr error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		confirms := channel.NotifyPublish(make(chan amqp091.Confirmation, 1))

		err := channel.PublishWithContext(
			context.Background(),
			exchange,
			routingKey,
			false,
			false,
			amqp091.Publishing{
				ContentType:  "application/json",
				Body:         body,
				DeliveryMode: amqp091.Persistent,
			},
		)
		if err != nil {
			lastErr = err
			time.Sleep(retryInterval)
			continue
		}

		confirmed, ok := <-confirms
		if !ok {
			lastErr = fmt.Errorf("confirm channel closed by broker")
			time.Sleep(retryInterval)
			continue
		}
		if !confirmed.Ack {
			lastErr = fmt.Errorf("broker refused a mensage (nack) in the %d try", attempt)
			time.Sleep(retryInterval)
			continue
		}

		return nil
	}

	return fmt.Errorf("fail on publishing after %d tries: %w", maxRetries, lastErr)
}

func CloseRabbitMQ() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

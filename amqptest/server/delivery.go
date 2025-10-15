package server

import (
	"time"

	"github.com/bluearchive/generic_cache_wabbit"
	amqp "github.com/bluearchive/generic_cache_wabbit/amqp"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

type (
	// Delivery is an interface to delivered messages
	Delivery struct {
		data          []byte
		headers       wabbit.Option
		tag           uint64
		consumerTag   string
		originalRoute string
		messageId     string
		channel       *Channel
		contentType   string
		timestamp     time.Time
		delivery      *amqp.Delivery
	}
)

func NewDelivery(ch *Channel, data []byte, tag uint64, messageId string, hdrs wabbit.Option, contentType string, route string, timestamp time.Time) *Delivery {
	d := amqp091.Delivery{
		// Acknowledger: ch,
		Body:        data,
		DeliveryTag: tag,
		MessageId:   messageId,
		ContentType: contentType,
		RoutingKey:  route,
		Timestamp:   timestamp,
	}
	d2 := amqp.Delivery{&d}
	return &Delivery{
		data:          data,
		headers:       hdrs,
		channel:       ch,
		tag:           tag,
		messageId:     messageId,
		contentType:   contentType,
		originalRoute: route,
		timestamp:     timestamp,
		delivery:      &d2,
	}
}

func (d *Delivery) Ack(multiple bool) error {
	return d.channel.Ack(d.tag, multiple)
}

func (d *Delivery) Nack(multiple, requeue bool) error {
	return d.channel.Nack(d.tag, multiple, requeue)
}

func (d *Delivery) Reject(requeue bool) error {
	return d.channel.Nack(d.tag, false, requeue)
}

func (d *Delivery) Body() []byte {
	return d.delivery.Body()
}

func (d *Delivery) Headers() wabbit.Option {
	return d.headers
}

func (d *Delivery) DeliveryTag() uint64 {
	return d.tag
}

func (d *Delivery) ConsumerTag() string {
	return d.consumerTag
}

func (d *Delivery) MessageId() string {
	return d.messageId
}

func (d *Delivery) Timestamp() time.Time {
	return d.timestamp
}

func (d *Delivery) ContentType() string {
	return d.contentType
}

func (d *Delivery) RoutingKey() string {
	return d.originalRoute
}

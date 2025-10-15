package server

import (
	"time"

	"github.com/bluearchive/generic_cache_wabbit"
	amqp "github.com/bluearchive/generic_cache_wabbit/amqp"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

type (
	// Delivery is an interface to delivered messages
	// FIXME: this shadows the actual Delivery in amqp/delivery.go, thus not testing its calls
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
	// amqp Delivery object
	d := amqp091.Delivery{
		Acknowledger: ch,
		Headers:     amqp091.Table(hdrs),
		ContentType: contentType,
		MessageId:   messageId,
		Timestamp:   timestamp,
		// valid on Consume only:
		ConsumerTag: "",
		// valid on Get only:
		DeliveryTag: tag,
		RoutingKey:  route,
		Body:        data,
	}
	// wabbit Delivery wrappers the amqp object
	d2 := amqp.Delivery{&d}

	// return a test Delivery object that delegates all its methods to the actual
	d3 := &Delivery{
		// some tests depend on the fields of this mock, so leave them for now
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
	return d3;
}

func (d *Delivery) Ack(multiple bool) error {
	return d.delivery.Acknowledger.Ack(d.delivery.DeliveryTag(), multiple)
}

func (d *Delivery) Nack(multiple, requeue bool) error {
	return d.delivery.Acknowledger.Nack(d.delivery.DeliveryTag(), multiple, requeue)
}

func (d *Delivery) Reject(requeue bool) error {
	return d.delivery.Acknowledger.Nack(d.delivery.DeliveryTag(), false, requeue)
}

func (d *Delivery) Body() []byte {
	return d.delivery.Body()
}

func (d *Delivery) Headers() wabbit.Option {
	return wabbit.Option(d.delivery.Headers())
}

func (d *Delivery) DeliveryTag() uint64 {
	return d.delivery.DeliveryTag()
}

func (d *Delivery) ConsumerTag() string {
	return d.delivery.ConsumerTag()
}

func (d *Delivery) MessageId() string {
	return d.delivery.MessageId()
}

func (d *Delivery) Timestamp() time.Time {
	return d.delivery.Timestamp()
}

func (d *Delivery) ContentType() string {
	return d.delivery.ContentType()
}

func (d *Delivery) RoutingKey() string {
	return d.delivery.RoutingKey()
}

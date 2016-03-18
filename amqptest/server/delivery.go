package server

type (
	// Delivery is an interface to delivered messages
	Delivery struct {
		data          []byte
		consumerTag   string
		originalRoute string
		channID       string // channel that published the message
	}
)

func NewDelivery(data []byte) *Delivery {
	return &Delivery{
		data: data,
	}
}

func (d *Delivery) Ack(multiple bool) error {
	return nil
}

func (d *Delivery) Nack(multiple, request bool) error {
	return nil
}

func (d *Delivery) Reject(requeue bool) error {
	return nil
}

func (d *Delivery) Body() []byte {
	return d.data
}

func (d *Delivery) DeliveryTag() uint64 {
	return 0
}

func (d *Delivery) ConsumerTag() string {
	return d.consumerTag
}

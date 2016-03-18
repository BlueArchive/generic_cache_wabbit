package server

import "fmt"

type Exchange interface {
	route(route string, message []byte) error
	addBinding(route string, q *Queue)
	delBinding(route string)
}

type TopicExchange struct {
	name     string
	bindings map[string]*Queue
}

func NewTopicExchange(name string) *TopicExchange {
	return &TopicExchange{
		name:     name,
		bindings: make(map[string]*Queue),
	}
}

func (t *TopicExchange) addBinding(route string, q *Queue) {
	t.bindings[route] = q
}

func (t *TopicExchange) delBinding(route string) {
	delete(t.bindings, route)
}

func (t *TopicExchange) route(route string, msg []byte) error {
	for bname, q := range t.bindings {
		if topicMatch(bname, route) {
			d := NewDelivery(msg)
			q.data <- d
			return nil
		}
	}

	return fmt.Errorf("Route '%s' doesn't match any routing-key", route)
}

type DirectExchange struct {
	name     string
	bindings map[string]*Queue
}

func NewDirectExchange(name string) *DirectExchange {
	return &DirectExchange{
		name:     name,
		bindings: make(map[string]*Queue),
	}
}

func (d *DirectExchange) addBinding(route string, q *Queue) {
	d.bindings[route] = q
}

func (d *DirectExchange) delBinding(route string) {
	delete(d.bindings, route)
}

func (d *DirectExchange) route(route string, msg []byte) error {
	if q, ok := d.bindings[route]; ok {
		q.data <- NewDelivery(msg)
		return nil
	}

	return fmt.Errorf("No bindings to route: %s", route)

}

package events

type Collector struct {
	events []*Event
}

func (c *Collector) Record(event Publishable) error {
	message := &Event{
		Key:     event.Key(),
		Topic:   event.Topic(),
		Payload: event,
	}
	c.events = append(c.events, message)
	return nil
}

func (c *Collector) All() []*Event {
	return c.events
}

func (c *Collector) Drain() []*Event {
	events := c.events
	c.events = make([]*Event, 0)
	return events
}

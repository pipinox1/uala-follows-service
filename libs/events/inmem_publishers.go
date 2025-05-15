package events

import (
	"context"
	"fmt"
)

type InmemEvents struct {
}

func NewInmemEvents() *InmemEvents {
	return &InmemEvents{}
}

func (k *InmemEvents) Publish(ctx context.Context, event Publishable) error {
	fmt.Println(fmt.Sprintf("publishing event with key: %s", event.Key()))
	return nil
}

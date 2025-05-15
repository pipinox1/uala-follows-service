package events

import (
	"context"
)

type Publishable interface {
	Key() string
	Topic() string
}

type Publisher interface {
	Publish(ctx context.Context, event Publishable) error
}

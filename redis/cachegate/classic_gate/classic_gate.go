package classic_gate

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type ClassicGate struct {
	client *redis.Client
}

func NewClassicGate(client *redis.Client) *ClassicGate {
	return &ClassicGate{client: client}
}

func (c *ClassicGate) Ping() bool {
	return c.client.Ping(context.Background()).Err() == nil
}

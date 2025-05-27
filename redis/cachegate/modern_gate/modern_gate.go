package modern_gate

import (
	"context"

	"github.com/redis/rueidis"
)

type ModernGate struct {
	client    rueidis.Client
	multihost bool
}

func NewModernGate(c rueidis.Client, host []string) *ModernGate {
	m := &ModernGate{client: c}
	if len(host) > 1 {
		m.multihost = true
	} else {
		m.multihost = false
	}
	return m
}

func (m *ModernGate) Ping() bool {
	ctx := context.Background()
	if !m.multihost {
		return m.client.Do(ctx, m.client.B().Arbitrary("ping").Build()).Error() == nil
	}
	r := m.client.DoMulti(ctx, m.client.B().Multi().Build(), m.client.B().Arbitrary("ping").Build())
	for _, i := range r {
		if i.Error() != nil {
			return false
		}
	}
	return true
}

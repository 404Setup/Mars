package redis

import (
	"log/slog"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/redis/rueidis"

	"Mars/redis/cachegate/classic_gate"
	"Mars/redis/cachegate/modern_gate"
	"Mars/redis/shared"
	"Mars/shared/configure"
)

// TODO

func NewRedis() {
	if !configure.Get().Redis.Enabled || configure.Get().Redis.Servers == nil || len(configure.Get().Redis.Servers) == 0 {
		return
	}
	if configure.Get().Redis.Driver == "modern" {
		bootModern()
	} else {
		bootClassic()
	}
}

func bootClassic() {
	rds := redis.NewClient(&redis.Options{
		Addr:     configure.Get().Redis.Servers[0],
		Username: configure.Get().Redis.Username,
		Password: configure.Get().Redis.Password, // no password set
		DB:       0,                              // use default DB
	})
	shared.Gateway = classic_gate.NewClassicGate(rds)
}

func bootModern() {
	var sentinel rueidis.SentinelOption
	if configure.Get().Redis.Username != "" || configure.Get().Redis.Password != "" {
		sentinel = rueidis.SentinelOption{
			Username: configure.Get().Redis.Username,
			Password: configure.Get().Redis.Password,
		}
	}
	rue, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress:           configure.Get().Redis.Servers,
		ClientTrackingOptions: []string{"PREFIX", "mars:", "PREFIX", "tranic:", "BCAST"},
		Sentinel:              sentinel,
	})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(-3)
	}
	shared.Gateway = modern_gate.NewModernGate(rue, configure.Get().Redis.Servers)
}

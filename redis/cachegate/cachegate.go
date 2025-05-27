package cachegate

type RedisCacheGate interface {
	Ping() bool
}

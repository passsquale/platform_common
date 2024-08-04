package redis

import (
	"time"

	"github.com/a1exCross/common/pkg/storage"

	"github.com/go-redis/redis"
)

type redisStorage struct {
	client *redis.Client
}

// NewRedisConnection - создает новое инстанс клиента Redis
func NewRedisConnection(opts *redis.Options) (storage.Redis, error) {
	client := redis.NewClient(opts)

	return &redisStorage{
		client: client,
	}, nil
}

// Ping - тестирует запрос к базе
func (r *redisStorage) Ping() error {
	return r.client.Ping().Err()
}

// Close - закрывает соединение с базой
func (r *redisStorage) Close() error {
	return r.client.Close()
}

// Get - позволяет получить значение по ключу
func (r *redisStorage) Get(key string) *redis.StringCmd {
	return r.client.Get(key)
}

// Set - сохраняет значение по ключу
func (r *redisStorage) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(key, value, expiration)
}

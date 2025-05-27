package cache

import (
	"github.com/banraiphisan/banraicore-api/config"
	"github.com/banraiphisan/banraicore-api/pkg/cache/inmem"
	"github.com/banraiphisan/banraicore-api/pkg/cache/redis"
	"time"
)

type Engine interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte, exp time.Duration) error
	Delete(key string) error
	Reset() error
	Close() error
	Ping() error
}

func NewCache(configuration *config.Configuration) (Engine, error) {
	switch configuration.Server.CacheDeploymentType {
	case 1:
		client, err := redis.NewStandaloneConn(configuration)

		return client, err
	case 2:
		client, err := redis.NewClusterConn(configuration)

		return client, err
	default:
		client := inmem.NewInMemoryCache()

		return client, nil
	}
}

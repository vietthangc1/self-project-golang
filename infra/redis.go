package infra

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/thangpham4/self-project/pkg/envx"
	"github.com/thangpham4/self-project/pkg/logger"
)

var (
	defaultRedisAddr string = "localhost:6379"
)

type RedisConfig struct {
	Address  string
	Password string
	DB       int32
}

func NewRedisClient(ctx context.Context, config RedisConfig) (*redis.Client, error) {
	l := logger.Factory("Setup Redis")
	rdp := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       int(config.DB),
	})
	_, err := rdp.Ping(ctx).Result()
	if err != nil {
		l.V(logger.LogErrorLevel).Error(err, "Cannot set up Redis", "address", config.Address, "password", config.Password)
		return nil, err
	}
	l.V(logger.LogInfoLevel).Info("Successfully setup Redis", "address", config.Address)
	return rdp, nil
}

func NewRedisConfig() RedisConfig {
	return RedisConfig{
		Address:  envx.String("REDIS_ADDRS", defaultRedisAddr),
		Password: envx.String("REDIS_PW", ""),
		DB:       envx.Int("REDIS_DB", 0),
	}
}

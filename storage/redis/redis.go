package redis

import (
    "auth/config"
    "auth/storage"
    "context"
    "fmt"
    "time"

    "github.com/redis/go-redis/v9"
)

type Store struct {
    db *redis.Client
}

func New(cfg config.Config) storage.IRedisStorage {
    fmt.Println("cfg.RedisHost: ", cfg.RedisHost)
    fmt.Println("cfg.RedisPort: ", cfg.RedisPort)
    fmt.Println("cfg.RedisPassword: ", cfg.RedisPassword) // Optional: for debugging purposes

    redisClient := redis.NewClient(&redis.Options{
        Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
        Password: cfg.RedisPassword, // Set the password
    })

    return Store{
        db: redisClient,
    }
}

func (s Store) SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error {
    statusCmd := s.db.SetEx(ctx, key, value, duration)
    if statusCmd.Err() != nil {
        return statusCmd.Err()
    }

    return nil
}

func (s Store) Get(ctx context.Context, key string) interface{} {
    return s.db.Get(ctx, key).Val()
}

func (s Store) Del(ctx context.Context, key string) error {
    statusCmd := s.db.Del(ctx, key)
    if statusCmd.Err() != nil {
        return statusCmd.Err()
    }
    return nil
}

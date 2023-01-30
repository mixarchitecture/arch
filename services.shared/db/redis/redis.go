package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Service interface {
	Set(k string, v string) error
	HSet(k string, v ...interface{}) error
	Get(k string) (string, error)
	HGetAll(k string) (map[string]string, error)
	SetEx(k string, v interface{}, d int64) error
	Del(k ...string) error
	Exist(k string) (bool, error)
}

type redisClient struct {
	client *redis.Client
}

type Config struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func New(cnf *Config) Service {
	rClient := redis.NewClient(&redis.Options{
		Addr:     cnf.Host + ":" + cnf.Port,
		Password: cnf.Password,
		DB:       cnf.DB,
	})
	return &redisClient{
		client: rClient,
	}
}

func (r *redisClient) Set(k string, v string) error {
	return r.client.Set(ctx, k, v, 0).Err()
}

func (r *redisClient) HSet(k string, v ...interface{}) error {
	return r.client.HSet(ctx, k, v).Err()
}

func (r *redisClient) Get(k string) (string, error) {
	return r.client.Get(ctx, k).Result()
}

func (r *redisClient) HGetAll(k string) (map[string]string, error) {
	return r.client.HGetAll(ctx, k).Result()
}

func (r *redisClient) SetEx(k string, v interface{}, d int64) error {
	_, err := r.client.Set(ctx, k, v, 0).Result()
	if err != nil {
		return err
	}
	_, err = r.client.Expire(ctx, k, time.Duration(d)*time.Minute).Result()
	return err
}

func (r *redisClient) Del(k ...string) error {
	return r.client.Del(ctx, k...).Err()
}

func (r *redisClient) Exist(k string) (bool, error) {
	res, err := r.client.Exists(ctx, k).Result()
	if err != nil {
		return false, err
	}
	return res == 1, nil
}

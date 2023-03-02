package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisService interface {
	SaveToken(key string, token string, TTL time.Duration)
	ValidateToken(key string) bool
	ClearToken(key string)
	CheckValueKey(key string, token string) bool
	GetValueByKey(key string) string
}

type redisService struct {
	Client *redis.Client
}

func NewRedisService(client *redis.Client) RedisService {
	return &redisService{
		Client: client,
	}
}

func (r redisService) CheckValueKey(key string, token string) bool {
	val := r.Client.Get(context.TODO(), key)
	if val.Val() == token {
		return true
	}
	return false
}

func (r redisService) GetValueByKey(key string) string {
	val := r.Client.Get(context.TODO(), key)
	val2, _ := val.Result()
	//fmt.Println(val.Val())
	//fmt.Println(val2)
	return val2
}

func (r redisService) SaveToken(key string, token string, TTL time.Duration) {
	r.Client.Set(context.TODO(), key, token, TTL)
}

func (r redisService) ValidateToken(key string) bool {
	check := r.Client.Exists(context.TODO(), key)
	if check.Err() != nil {
		return false
	}
	return true
}

func (r redisService) ClearToken(key string) {
	r.Client.Del(context.TODO(), key)
}

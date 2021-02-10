package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"okr/domain"
)

type redisRepository struct {
	resdis *redis.Client
}

func NewRedisRepository(resdis *redis.Client) domain.RedisRepository {
	return &redisRepository{
		resdis: resdis,
	}
}

var ctx = context.Background()

func (rr redisRepository) RedisSet(json string, key string) error {
	err := rr.resdis.Set(key, json, 0).Err()
	if err != nil {
		fmt.Println("redis.Client.Set Error:", err)
		return err
	}
	return nil
}

func (rr redisRepository) RedisGet(key string) (*domain.User, error) {
	userInfoJson, _ := rr.resdis.Get(key).Result()
	var user = new(domain.User)
	if err := json.Unmarshal([]byte(userInfoJson), user); err != nil {
		return nil, err
	}
	return user, nil
}

func (rr redisRepository) RedisDelete(token string) {
	rr.resdis.Del(token)
}
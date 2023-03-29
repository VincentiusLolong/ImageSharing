package database

import (
	"encoding/json"
	"mestorage/models"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis interface {
	Set(key string, data interface{}, time int) error
	Get(key string) ([]byte, error)
	Delete(key string) (bool, error)
}

type rediscon struct {
	redisPool redis.Conn
}

func NewRedis() Redis {
	return &rediscon{}
}

var redisConn *redis.Pool

func Setup() error {
	redisConn = &redis.Pool{
		MaxIdle:     models.RedisSetting.MaxIdle,
		MaxActive:   models.RedisSetting.MaxActive,
		IdleTimeout: models.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", models.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if models.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", models.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", models.RedisSetting.DB); err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

func (r *rediscon) Set(key string, data interface{}, time int) error {
	r.redisPool = redisConn.Get()
	defer r.redisPool.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = r.redisPool.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = r.redisPool.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func (r *rediscon) Get(key string) ([]byte, error) {
	r.redisPool = redisConn.Get()
	defer r.redisPool.Close()

	reply, err := redis.Bytes(r.redisPool.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (r *rediscon) Delete(key string) (bool, error) {
	r.redisPool = redisConn.Get()
	defer r.redisPool.Close()

	return redis.Bool(r.redisPool.Do("DEL", key))
}

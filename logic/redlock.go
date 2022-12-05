package logic

import (
	"time"

	"github.com/go-redis/redis"
)

type (
	RedLockInterface interface {
		SetNX(key string, tag string)
		DelNX(key string, tag string)
		LockState(key string)
	}

	RedLock struct {
		Keys             []string
		Tag              string
		HeartBeatSeconds int
		Redisconn        *redis.Client
	}
)

func NewRedLock(redisconn *redis.Client, heartBeatSeconds int, tag string, keys []string) RedLockInterface {
	if redisconn == nil {

	}

	go heartBeat(tag, heartBeatSeconds, redisconn)

	return &RedLock{
		Keys:             keys,
		Tag:              tag,
		HeartBeatSeconds: heartBeatSeconds,
		Redisconn:        redisconn,
	}
}

func (r *RedLock) SetNX(key string, tag string)

func (r *RedLock) DelNX(key string, tag string)

func (r *RedLock) LockState(key string)

func heartBeat(tag string, heartBeatSeconds int, redisconn *redis.Client) {
	for {
		redisconn.Do("SET", tag, time.Now().Unix())
		time.Sleep(time.Duration(heartBeatSeconds) * time.Second)
	}
}

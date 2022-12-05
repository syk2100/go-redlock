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

	redLock struct {
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

	return &redLock{
		Keys:             keys,
		Tag:              tag,
		HeartBeatSeconds: heartBeatSeconds,
		Redisconn:        redisconn,
	}
}

func (r *redLock) SetNX(key string, tag string)

func (r *redLock) DelNX(key string, tag string)

func (r *redLock) LockState(key string)

func heartBeat(tag string, heartBeatSeconds int, redisconn *redis.Client) {
	for {
		redisconn.Do("SET", tag, time.Now().Unix())
		time.Sleep(time.Duration(heartBeatSeconds) * time.Second)
	}
}

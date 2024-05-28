package locker

import (
	"context"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
	"xuanfeng_gin/pkg/log"
)

type RedisLocker struct {
	cli *redis.Client
}

func NewLocker(cli *redis.Client) *RedisLocker {
	s := &RedisLocker{
		cli,
	}
	return s
}

func (s *RedisLocker) GetLock(key string, ttl time.Duration, ctx context.Context) (*redislock.Lock, error) {
	locker := redislock.New(s.cli)
	lock, err := locker.Obtain(ctx, key, ttl, nil)
	return lock, err
}

func (s *RedisLocker) GetRefreshLock(key string, ctx context.Context) (lock *redislock.Lock, err error) {
	locker := redislock.New(s.cli)
	lock, err = locker.Obtain(ctx, key, time.Second*10, nil)
	if err != nil {
		log.L.Error(err.Error())
		return
	}

	go func() {
		for {
			time.Sleep(time.Second * 5)
			ttl, e := lock.TTL(ctx)
			if e != nil || ttl == 0 {
				return
			}
			e = lock.Refresh(ctx, time.Second*10, nil)
			if e != nil {
				return
			}
		}
	}()

	return
}

func (s *RedisLocker) IsNotObtained(err error) bool {
	if strings.Contains(err.Error(), "not obtained") {
		return true
	}
	return false
}

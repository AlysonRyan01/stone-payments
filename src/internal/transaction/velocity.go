package transaction

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisVelocity struct {
	rdb *redis.Client
}

func NewRedisVelocity(rdb *redis.Client) *RedisVelocity {
	return &RedisVelocity{rdb: rdb}
}

var velocityScript = redis.NewScript(`
local key    = KEYS[1]
local now    = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
redis.call('ZREMRANGEBYSCORE', key, 0, now - window)
local count = redis.call('ZCARD', key)
redis.call('ZADD', key, now, ARGV[3])
redis.call('PEXPIRE', key, window)
return count
`)

func (v *RedisVelocity) IsNotStandard(ctx context.Context, card string, now time.Time) (bool, error) {
	count, err := velocityScript.Run(ctx, v.rdb,
		[]string{"velocity:" + card},
		now.UnixMilli(), 60_000, uuid.NewString(),
	).Int()
	if err != nil {
		return false, err
	}
	return count >= 5, nil
}
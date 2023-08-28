package utils

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

var luaScript = `
local tokens_key = KEYS[1]
local timestamp_key = KEYS[2]
local max_tokens = tonumber(ARGV[1])
local refill_interval = tonumber(ARGV[2])
local current_timestamp = tonumber(ARGV[3])

local last_timestamp = tonumber(redis.call("GET", timestamp_key) or "0")
local tokens = tonumber(redis.call("GET", tokens_key) or "0")

local elapsed_time = current_timestamp - last_timestamp
local tokens_to_add = elapsed_time / refill_interval
tokens = math.min(tokens + tokens_to_add, max_tokens)

local allowed = tokens >= 1 and 1 or 0
if allowed then 
    tokens = tokens - 1
    redis.call("SET", tokens_key, tokens)
    redis.call("SET", timestamp_key, current_timestamp)
end
-- 在 Lua 脚本中，返回 false 时，Go 的 github.com/go-redis/redis/v8 包会将其解释为 nil。
return allowed
`

type LimiterBucket struct {
	limiter          string
	limiteStamp      string
	burst            int64
	rate             float64
	redis            *redis.Client
	currentTimeStamp int64
}

func NewLimiterBucket(limiter, limiterStamp string, burst int64, rate time.Duration, client *redis.Client) *LimiterBucket {

	return &LimiterBucket{
		limiter:          limiter,
		limiteStamp:      limiterStamp,
		burst:            burst,
		rate:             rate.Seconds(),
		redis:            client,
		currentTimeStamp: time.Now().Unix(),
	}
}

func (l *LimiterBucket) Allow() bool {

	luaScriptSHA := l.redis.ScriptLoad(context.Background(), luaScript).Val()
	result, err := l.redis.EvalSha(context.Background(), luaScriptSHA, []string{l.limiter, l.limiteStamp}, l.burst,
		l.rate, l.currentTimeStamp).Result()
	if err != nil {
		fmt.Printf("use script err :%s\n", err)
		return false
	}
	allowd, ok := result.(int64)
	if !ok {
		fmt.Printf("gen token err:%s\n", err)
		return false
	}
	if allowd == 1 {
		return true
	}
	return false
}

func (l *LimiterBucket) Wait() bool {

	return false
}

func (l *LimiterBucket) Drop() bool {

	return false
}

package utils

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

func main1() {
	ctx := context.Background()
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 地址
	})
	// 限流参数
	maxTokens := 10 // 最大令牌数
	refillRate := 2 // 令牌生成速率（每秒）
	refillInterval := time.Second / time.Duration(refillRate)
	fmt.Println("generato", refillInterval)

	// 调用 Lua 脚本进行限流
	luaScript := `
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
	luaScriptSHA := client.ScriptLoad(ctx, luaScript).Val()
	// 模拟请求
	for i := 1; i <= 20; i++ {
		// 获取当前时间戳
		currentTimestamp := time.Now().Unix()
		// 调用 Lua 脚本
		result, err := client.EvalSha(ctx, luaScriptSHA, []string{"limiter_tokens", "limiter_timestamp"},
			maxTokens, refillInterval.Seconds(), currentTimestamp).Result()
		if err != nil {
			fmt.Printf("Lua script: %s,%v\n", err, result)
			return
		}
		allowed, ok := result.(int64)
		if !ok {
			fmt.Println("get tokens", result)
			return
		}
		if allowed == 1 {
			fmt.Printf("Request %d allowed at %s\n", i, time.Now().Format(time.RFC3339))
		} else if allowed == 0 {
			fmt.Printf("Request %d denied at %s\n", i, time.Now().Format(time.RFC3339))
		} else {
			fmt.Printf("Invalid Request  i;%d , at:%s\n", i, time.Now().Format(time.RFC3339))
		}
		time.Sleep(500 * time.Millisecond)
	}
}

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

	result, err := l.redis.EvalSha(context.Background(), luaScript, []string{l.limiter, l.limiteStamp}, l.burst,
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

// 直接在这里封装一个限流中间件 完成对服务端的限流请求的任务，后期可以优化为对限流外的请求进行阻塞等待
func LimiterInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 地址
	})
	// 限流参数
	//maxTokens := 10 // 最大令牌数
	refillRate := 2 // 令牌生成速率（每秒）
	refillInterval := time.Second / time.Duration(refillRate)
	limiter := NewLimiterBucket("bucket_limiter", "limiter_timestamp",
		10, refillInterval, client)
	if !limiter.Allow() {
		fmt.Println("limiter success!")
		return nil, nil
	}
	response, err := handler(ctx, req)
	if err != nil {
		return nil, nil
	}
	return response, nil
}

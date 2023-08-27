package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"time"
)

func main() {
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

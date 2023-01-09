package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// 集群连接
func cluster() {
	db := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	fmt.Println(db)
}

var redis_client *redis.Client

func init() {
	//redi的配置
	redis_option := &redis.Options{
		Addr:         "127.0.0.1:6379",
		DialTimeout:  time.Millisecond * 100,
		ReadTimeout:  time.Millisecond * 100,
		WriteTimeout: time.Millisecond * 200,
		PoolSize:     20,
		MinIdleConns: 3,
		MaxConnAge:   50,
	}
	redis_client = redis.NewClient(redis_option)
}
func get() {
	var ctx = context.Background()
	val, err := redis_client.Get(ctx, "key").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	fmt.Println(val)
}

func scanKeys() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var cursor uint64
	for {
		var keys []string
		var err error
		// 按前缀扫描key
		keys, cursor, err = redis_client.Scan(ctx, cursor, "prefix:*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}
}

func listDemo(ctx context.Context) {

	var err error
	err = redis_client.RPush(ctx, "mylist", "sunweiming").Err()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	err = redis_client.LPush(ctx, "mylist", "yujinling").Err()
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	length, err := redis_client.LLen(ctx, "mylist").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(" Length : %+v\n", length)

	s1, err := redis_client.LSet(ctx, "mylist", 1, "sunweiming").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s1)
	//获取当前当前下标元素
	elem, err := redis_client.LIndex(ctx, "mylist", 1).Result()
	if err != nil {
		fmt.Printf("LIndex : %s\n", err)
		return
	}
	fmt.Println(elem)
	//在当前链表中插入数据
	elem1, err := redis_client.LInsert(ctx, "mylist", "after", "sunweiming", "yujingling").Result()
	if err != nil {
		fmt.Printf("LIndex : %s\n", err)
		return
	}
	fmt.Println(elem1)
	//获取当前所有元素
	res, err := redis_client.LRange(ctx, "mylist", 1, -1).Result() //获取list中所有元素
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("LRange : %+v\n", res)
	//获取所有keys
	keys, err := redis_client.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)

	for i := 0; int64(i) < length; i++ {
		s, err := redis_client.LPop(ctx, "mylist").Result()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(" LPop : %+v\n", s)

	}
	res, err = redis_client.LRange(ctx, "mylist", 1, -1).Result() //获取list中所有元素
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("LRange : %+v\n", res)
}

func zsetDeomo(ctx context.Context) {
	// key
	zsetKey := "language_rank"
	// value
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	// ZADD
	err := redis_client.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := redis_client.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := redis_client.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redis_client.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
func Keys(ctx context.Context) {
	keys, err := redis_client.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
}
func stringDemo(ctx context.Context) {
	var err error
	res, err := redis_client.Set(ctx, "abc", "123", time.Second*60).Result()
	if err != nil {

		fmt.Printf("string set : %s\n", err)
	}
	fmt.Println(res)
	//设置的时候同时设置过期时间
	res, err = redis_client.SetEX(ctx, "a", "97", time.Second*60).Result()
	if err != nil {

		fmt.Printf("string set : %s\n", err)
	}
	fmt.Println(res)

	res, err = redis_client.SetEX(ctx, "b", "98", time.Second*60).Result()
	if err != nil {

		fmt.Printf("string set : %s\n", err)
	}
	fmt.Println(res)

	result10, _ := redis_client.MGet(ctx, "a", "b").Result()
	fmt.Println("result10", result10)

	result := redis_client.Get(ctx, "a")
	result_str, err := result.Bytes()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result_str))

	va, err := redis_client.Incr(ctx, "a").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(va)

	va, err = redis_client.IncrBy(ctx, "a", 12).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(va)

	va, err = redis_client.Decr(ctx, "a").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(va)

	va, err = redis_client.DecrBy(ctx, "a", 10).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(va)
	//在字符串后面追加字符
	ap, err := redis_client.Append(ctx, "abc", "sunweiming").Result()
	if err != nil {
		fmt.Printf("string Append %s\n ", err)
	}
	fmt.Println(ap)

	result1, err := redis_client.Get(ctx, "abc").Result()
	if err != nil {
		fmt.Printf("string Append %s\n ", err)
	}
	fmt.Println(result1)

	//获取字符串长度
	length, err := redis_client.StrLen(ctx, "abc").Result()
	if err != nil {
		fmt.Printf("string StrLen %s\n ", err)

	}
	fmt.Println(length)
}
func getAll(ctx context.Context) {
	keys, err := redis_client.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
}

func hsetDemo(ctx context.Context) {
	//添加元素
	result, err := redis_client.HSet(ctx, "use", "key1", "value1").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	result, err = redis_client.HSet(ctx, "use", "key2", "value2").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	key1Result, err := redis_client.HGet(ctx, "use", "key1").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(key1Result)
	//删除某个元素
	delResult, err := redis_client.HDel(ctx, "use", "key1").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("HDel :%+v\n ", delResult)
	//获取某个元素
	allResult, err := redis_client.HGetAll(ctx, "use").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("HGetAll :%+v\n ", allResult)
	//判断某个元素是否存在
	existResult, err := redis_client.HExists(ctx, "use", "key1").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("HExists :%+v\n ", existResult)
	//获取hset 的长度
	length, err := redis_client.HLen(ctx, "use").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("HElen :%+v\n ", length)
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	//listDemo(ctx)
	//zsetDeomo(ctx)
	stringDemo(ctx)
	//hsetDemo(ctx)
	//Keys(ctx)
}

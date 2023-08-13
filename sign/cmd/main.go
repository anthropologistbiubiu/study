package main

import (
	"os"
	"sign/db/dmysql"
	"sign/utils/log"
)

// 实现这个业务的完整性，在业务层中添加 orm 的过程。
// 添加orm 事务的处理过程。
// 添加 redis 分布式缓存 了解 分布式缓存的特点
// 添加服务的注册与发现  添加负载均衡
// 添加 log 层的日志归档和记录
// 抽象出 aes md5 rsa 等这些 服务接口
// 有 grpc 的认证过程
// 添加配置文件的解析过程
// 掌握分布式的特点
// 加上单元测试

// 1.聚合支付
// 支付服务的 主要逻辑  想想能不能把 kafka 这个消息队列用起来
// 白名单 + 一个限流算法 令牌桶算法 / 滑动窗口算法
// 支付服务当中 怎么把定时任务加进去   要持续的优化这个服务
// es + 实现大数据的查询
// 梳理清楚分库分表的逻辑
// 分布式缓存  + 数据一致性
// 设计模式
// nginx 负载均衡

//2.虚拟货币支付交易
// ihive服务的技术栈 + 多了一个scaner 服务的部署  + kafk 数据的推送的服务 +  transfer 服务的调用(grpc) + 支付服务的主体逻辑 + 预警服
// 这个服务当中重要的一些逻辑就是缓存的处理 (string,hash,zset,list)  还有就是 + 数据精度的处理 + channel + 协程 + 接口

func main() {
	//defer log.Sync()
	// lumberjack
	logger := log.New(os.Stderr, log.WarnLevel, log.AddCaller(), log.AddCallerSkip(2))
	log.ReplaceDefault(logger)
	log.Warn("sunweiming")
	dmysql.InitMysql()

}

// 封装出 access.log error.log info.log
// grpc 怎么封装 access.log ----- 拦截器
// gin 中间件的使用
/*
func main() {
	serviceAddr := "localhost:55001" // 替换为实际的服务器地址
	serviceName := "sign-service"
	etcd.RegisterServiceWithEtcd(serviceName, serviceAddr)

	listen, err := net.Listen("tcp", serviceAddr)
	if err != nil {
		fmt.Println("", err)
	}
	server := grpc.NewServer()
	reflection.Register(server)
	proto.RegisterSignServiceRequestServer(server, &service.SignServer{})
	var wg sync.WaitGroup
	wg.Add(1)
	// 处理优雅退出信号
	go func() {
		defer wg.Done()
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		fmt.Println("\nReceived shutdown signal. Gracefully shutting down...")
		// 关闭gRPC服务器
		server.GracefulStop()
	}()

	fmt.Println("start success!")
	// 启动gRPC服务器
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// 等待所有活动的请求处理完成
	wg.Wait()
	// 在服务器退出时，注销服务
	etcd.UnregisterServiceWithEtcd("", "")
}
*/

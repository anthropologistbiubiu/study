package main

import "github.com/redis/go-redis/v9"

// 直接在这里写一个回调服务算了，通过发布与订阅模式来完成。
// 写一个阻塞的服务来完成读消息队列的数据，不提供接口，但是要实现对服务的优雅退出和重启。不能丢失消费的数据。
/*

根据之前的服务对kafka服务写一个优雅退出过程的服务。
第一 实现对订单状态的数据库修改，第二对回调订单使用协程发送请求。第三，对修改状态后的订单标记offset.
如果标记后从管道读取下一个数据那么并发也就是失去了意义。 第四 加上并发控制。

验证demo1 :设计并发标记数据的时机。如果是全部处理完成，那么数据的处理就是在顺序处理了。

测试如果每标记一个，才会读取下个数据的话，那么是否可实现并发？

第一 实现证明数据的并发消费 在宕机的情况下是不可靠的，会丢失数据，针对并发消费最好只设定好标记时机。

第二 并发消费数据，采取手动标记数据其实是顺序执行的。

第三 针对并发数据的消费要合理设置分区，合理设置标记时机。

第四 要解决的是在优雅退出下，并发请求的优雅处理和消费者的关闭。

*/
// 上面记录的是对消息队列使用存在的疑惑,我这里来一步一步解决,后期来搭建为docker版本的集群模式

var (
	messageQueue = new(MessageQueue)
)

type MessageQueue struct {
	client *redis.Client
}

func init() {
	messageQueue.client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func (m *MessageQueue) read() {

}
func (m *MessageQueue) write() {

}

func Producer() {

}

func Consumer() {

}
func main() {

}

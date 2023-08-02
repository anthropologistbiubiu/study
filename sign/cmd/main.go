package main

import "github.com/gin-gonic/gin"

//grpc框架    +  protobuf  重构签名服务器

// 每天晚上回来扩充一些

// 明确一下需求 现在要写一个签名服务器 需要封装 日志文件 ，日志文件包括log 文件  access 文件,
// 核心是通过 proto 协议定义，生成 rpc 文件，完成调用和请求。
// 使用 gin + endless 网络框架 实现平滑关闭。
// 封装 mysql + cache + orm 来完成数据的处理。
// 后续扩充对日志的归档 golang 我觉得就过了

func main() {

	server := gin.Default()

	// 使用grpc 来调用

}

package main

import "fmt"

func main() {

	fmt.Println("hello world")

}

// 在这里启动一个服务，通过http 的rest请求来完整对grpc 服务的调用
// 1.如果gw.pb.go文件可以完成，先通过这种方式来完成
// 2.再通过http.pb.go文件或者两个文件来共同完成http->grpc 的服务。

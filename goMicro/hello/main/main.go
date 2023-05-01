package main

/*
import (
    "context"
    "github.com/asim/go-micro/plugins/registry/consul/v3"
    "github.com/asim/go-micro/v3"
    "github.com/asim/go-micro/v3/registry"
    proto "go-micro-test/proto"
    "log"
)
type Greeter struct{}
func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
    rsp.Greeting = "Hello " + req.Name
    return nil
}
func main() {
    consulReg := consul.NewRegistry(registry.Addrs(":8500"))
    service := micro.NewService(
        micro.Name("greeter"),
        micro.Registry(consulReg),
    )
    service.Init()
    proto.RegisterGreeterHandler(service.Server(), new(Greeter))
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
*/

package main

import (
	"fmt"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"net/http"
)

func main() {
	cli := client.NewClient(
		client.WithBalancer(balancer.NewRoundRobin()),
		client.WithEndpoint("http://service1"),
	)

	app := kratos.New(
		kratos.Name("http-client"),
		kratos.Server(
			server.NewServer(
				server.Name("http-server"),
				server.Address(":8000"),
			),
		),
		kratos.HttpClient(cli),
	)

	r := app.HTTPRouter()
	r.GET("/", func(ctx context.Context, req *http.Request) (interface{}, error) {
		return "Hello from client", nil
	})

	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}

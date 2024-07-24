// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"payhub/internal/biz"
	"payhub/internal/conf"
	"payhub/internal/data"
	"payhub/internal/server"
	"payhub/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	paymentRepo := data.NewPaymentRepo(dataData, logger)
	paymentOrderUsecase := biz.NewPaymentOrderUsecase(paymentRepo, logger)
	paymentOrderService := service.NewPaymentOrderService(paymentOrderUsecase)
	grpcServer := server.NewGRPCServer(confServer, paymentOrderService, logger)
	httpServer := server.NewHTTPServer(confServer, paymentOrderService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}

package service

import (
	"context"
	v1 "payhub/api/helloworld/v1"
	"payhub/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer
	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

type PaymentOrderService struct {
	v1.UnimplementedPaymentSerivceServer
	uc *biz.PaymentOrderUsecase
}

func NewPaymentOrderService(uc *biz.PaymentOrderUsecase) *PaymentOrderService {
	return &PaymentOrderService{uc: uc}
}

func (s *PaymentOrderService) CreatePaymentOrder(ctx context.Context, in *v1.PaymentCreateRequest) (*v1.PaymentCreateReply, error) {
	return &v1.PaymentCreateReply{
		Status: 200,
		PayUrl: "www.success.com",
	}, nil
	if err := s.uc.CreatePaymentOrder(ctx, &biz.PaymentOrder{MerchantID: in.Merchantid, Amount: in.Amount}); err != nil {
		return &v1.PaymentCreateReply{Status: 401, PayUrl: ""}, err
	}
	return &v1.PaymentCreateReply{Status: 200, PayUrl: "www.baidu.com"}, nil
}

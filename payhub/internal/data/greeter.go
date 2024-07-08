package data

import (
	"context"

	"payhub/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}

type PaymentOrderRepo struct {
	data *Data
	log  *log.Helper
}

func NewPaymentRepo(data *Data, logger log.Logger) biz.PaymentRepo {
	return &PaymentOrderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *PaymentOrderRepo) Save(ctx context.Context, g *biz.PaymentOrder) error {
	// 数据库操作
	return r.data.Mysql.Create(g).Error
}

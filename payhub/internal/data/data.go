package data

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"payhub/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewPaymentRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Mysql *gorm.DB
	Redis *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	cache := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return &Data{Mysql: db, Redis: cache}, cleanup, nil
}

package dmysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sign/utils/log"
)

// 初始化orm 的mysql 方法
// 连接本地数据库 要存储哪些数据,每个请求的数据吗
// 提供路由可以更改 不同类型的签名的sescret
// 查询不同类型的签名的secret

var (
	Orms     = &Orm{}
	err      error
	userName = "sunweiming"
	dbName   = ""
	passWord = ""
)

type Orm struct {
	DB *gorm.DB
}

func InitMysql() {
	dsn := fmt.Sprintf("%s %s %s", userName, dbName, passWord)
	Orms.DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Error("", zap.String("err", fmt.Sprintf("%s", err)))
	}
	// 最后改一下配置
}

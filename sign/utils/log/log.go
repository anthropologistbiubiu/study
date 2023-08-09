package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func log() {

}

// 调用日志等级 函数输出调用调用等级的链路追踪,如行号，调用文件，日志日志初始化，日志文件的堆栈。
// 完成日志的切割与归档
type Level = zapcore.Level

const (
	Infolevel = zap.InfoLevel
)

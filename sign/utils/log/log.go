package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

func log() {

}

// 调用日志等级 函数输出调用调用等级的链路追踪,如行号，调用文件，日志日志初始化，日志文件的堆栈。
// 完成日志的切割与归档
type Level = zapcore.Level

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
)

type Logger struct {
	l  *zap.Logger
	al *zap.AtomicLevel
}

func New(out io.Writer, level Level) *Logger {
	if out == nil {
		out = os.Stderr
	}

	al := zap.NewAtomicLevelAt(level)
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.AddSync(out),
		al,
	)
	return &Logger{l: zap.New(core), al: &al}
}

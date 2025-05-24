package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


func New(logLevel, serviceName string) (*zap.Logger, error) {
	
	cfg := zap.NewProductionConfig()                       
	
	switch logLevel {
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "dpanic":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DPanicLevel)
	case "fatal":
		cfg.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	default:
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	// 禁用栈跟踪
	cfg.DisableStacktrace        = true 
	// 时间戳格式
	cfg.EncoderConfig.EncodeTime = formattedTimeEncoder  
	// 输出格式
	cfg.Encoding                 = "console" 
	// 输出位置 stdout / file
	cfg.OutputPaths              = []string{"stdout"} 
	// 初始化字段
	cfg.InitialFields            = map[string]any{"service-name": serviceName}

	// 构建 logger
	return cfg.Build()
}


// 格式化时间戳格式
func formattedTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}


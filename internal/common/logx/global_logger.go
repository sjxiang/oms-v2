package logx

import (
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 格式化时间戳格式
func formattedTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}


// 全局 logger
var logger *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	config.EncoderConfig.EncodeTime = formattedTimeEncoder  // 自定义时间戳格式
	config.DisableStacktrace = true                         // 打印堆栈
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)  // 日志级别, 过滤
	
	if isLocal, _ := strconv.ParseBool(os.Getenv("LOCAL_ENV")); isLocal {
		// 本地环境, 打印到终端
		config.Encoding = "console"
		config.OutputPaths = []string{"stdout"}
	} else {
		// 线上环境, 打印到文件和终端
		config.Encoding = "json"
		config.OutputPaths = []string{"stdout", "gua.log"}
	}

	config.InitialFields = map[string]any{
		"author": "gua",
	}

	logger, err = config.Build()

	if err != nil {
		panic(err)
	}
}


func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	logger.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	logger.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}




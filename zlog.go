package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(ops ...Option) *zap.SugaredLogger {
	options := newOptions()
	for _, op := range ops {
		op(options)
	}

	writer := zapcore.AddSync(options.output)

	// 格式相关的配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writer, zapcore.Level(options.level))
	return zap.New(core, zap.WithCaller(options.caller)).Sugar()
}

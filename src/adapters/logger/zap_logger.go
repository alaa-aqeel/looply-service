package logger

import (
	"github.com/alaa-aqeel/looply-app/src/core/ports"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	zap *zap.Logger
}

func newZapLogger() (ports.LoggerPort, error) {

	config := zap.NewDevelopmentConfig()
	config.Encoding = "json"
	config.OutputPaths = []string{
		"logs/app.log", // log file
	}
	config.ErrorOutputPaths = []string{"logs/error.log", "stderr"}
	zapLogger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{zap: zapLogger}, nil
}

func (z *ZapLogger) getLevel(level ports.Level) zapcore.Level {
	switch level {
	case ports.Debug:
		return zap.DebugLevel
	case ports.Info:
		return zap.InfoLevel
	case ports.Warn:
		return zap.WarnLevel
	case ports.Error:
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func (z *ZapLogger) Write(level ports.Level, message string, fields ...zapcore.Field) {
	z.zap.Check(z.getLevel(level), message).Write(fields...)
}

func (z *ZapLogger) Log(level ports.Level, log ports.Logger) {
	z.Write(level, log.Message,
		zap.String("tag", log.Tag),
		zap.Any("args", log.Args),
		zap.Duration("duration", log.Duration),
		zap.Error(log.Error),
	)
}

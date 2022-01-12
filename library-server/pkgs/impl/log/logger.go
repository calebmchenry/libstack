package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DefaultLogger creates a default logger
func DefaultLogger(cmd string) (*zap.Logger, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Sampling = nil
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapConfig.Build(zap.Fields(zap.String("job", cmd)))
}

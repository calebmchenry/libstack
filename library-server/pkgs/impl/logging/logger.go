package logging

import (
	"bytes"

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

// BufferLogger can be used in your unit tests. Log messages are written to the
// returned *bytes.Buffer.
func BufferLogger() (*zap.Logger, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	ws := zapcore.AddSync(buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
	enc := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(enc, ws, zapcore.InfoLevel)
	return zap.New(core), buf
}

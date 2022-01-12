package logging

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// define a non-public type to key this logger
type loggerKeyType int

// define a non-public const for keying this logger in the context
const loggerKey loggerKeyType = iota

// NewContext wraps the current logging context into this context
func NewContext(ctx context.Context, l *zap.Logger, fields ...zap.Field) context.Context {
	return context.WithValue(ctx, loggerKey, l.With(fields...))
}

// From extracts the logger from the context
func From(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return logger
	}
	panic(fmt.Errorf("logger missing from context.Context"))
}

package log

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
)

type message struct {
	Message struct {
		Msg string `json:"msg"`
	} `json:"message"`
	Level messageType
}

func newMessage(msgType messageType, msg string) *message {
	return &message{
		Message: struct {
			Msg string `json:"msg"`
		}{Msg: msg},
		Level: msgType,
	}
}

type messageType string

const (
	TRACE messageType = "TRACE"
	DEBUG             = "DEBUG"
	INFO              = "INFO"
	WARN              = "WARN"
	ERROR             = "ERROR"
	FATAL             = "FATAL"
)

type loggerKey struct{}

func FromCtx(ctx context.Context) *zap.Logger {
	value := ctx.Value(loggerKey{})
	if value == nil {
		return nil
	}
	return value.(*zap.Logger)
}

func CtxWithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func CtxWithFields(ctx context.Context, fs ...zap.Field) context.Context {
	logger := FromCtx(ctx)
	if logger == nil {
		return ctx
	}
	return CtxWithLogger(ctx, logger.With(fs...))
}

func doWithLogger(ctx context.Context, action func(*zap.Logger)) {
	logger := FromCtx(ctx)
	if logger == nil {
		_, _ = fmt.Fprint(os.Stderr, "NO LOGGER!\n")
		return
	}
	action(logger)
}

func Debug(ctx context.Context, msg string, fs ...zap.Field) {
	doWithLogger(ctx, func(logger *zap.Logger) {
		logger.Debug(msg, fs...)
	})
}

func Info(ctx context.Context, msg string, fs ...zap.Field) {
	doWithLogger(ctx, func(logger *zap.Logger) {
		logger.Info(msg, fs...)
	})
}

func Warn(ctx context.Context, msg string, fs ...zap.Field) {
	doWithLogger(ctx, func(logger *zap.Logger) {
		logger.Warn(msg, fs...)
	})
}

func Error(ctx context.Context, msg string, fs ...zap.Field) {
	doWithLogger(ctx, func(logger *zap.Logger) {
		logger.Error(msg, fs...)
	})
}

func Sync(ctx context.Context) {
	doWithLogger(ctx, func(logger *zap.Logger) {
		logger.Sync()
	})
}

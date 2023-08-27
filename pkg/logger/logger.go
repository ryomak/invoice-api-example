package logger

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"os"
)

var logger *slog.Logger

type LogKind string

type Log struct {
	Message string
	User    LogUser
}

type LogUser struct {
	UserID string
}

func Build() error {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return nil
}

func InfofWithData(ctx context.Context, data map[string]any, format string, args ...any) {
	logger.InfoContext(
		ctx,
		fmt.Sprintf(format, args...),
		slog.String("userId", User(ctx)),
		slog.Any("data", data),
	)
}

func Infof(ctx context.Context, format string, args ...any) {
	logger.InfoContext(
		ctx,
		fmt.Sprintf(format, args...),
		slog.String("userId", User(ctx)),
	)
}

func Warningf(ctx context.Context, format string, args ...any) {

	logger.WarnContext(
		ctx,
		fmt.Sprintf(format, args...),
		slog.String("userId", User(ctx)),
	)
}

func Errorf(ctx context.Context, format string, args ...any) {

	logger.ErrorContext(
		ctx,
		fmt.Sprintf(format, args...),
		slog.String("userId", User(ctx)),
	)
}

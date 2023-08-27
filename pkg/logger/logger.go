package logger

import (
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

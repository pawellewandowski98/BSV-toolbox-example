package logger

import "log/slog"

type Logger struct {
	slog.Logger
}

func New() *Logger {
	log := slog.Default()

	return &Logger{*log}
}

func (l *Logger) Start(msg string, args ...any) {
	l.Logger.Info("🚀\t"+msg, args...)
}

func (l *Logger) Complete(msg string, args ...any) {
	l.Logger.Info("🚀\t"+msg+" ✅  ✅  ✅", args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.Logger.Info("ℹ️\t"+msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.Logger.Error("❌\t"+msg, args...)
}

package noop

import "cosmossdk.io/log"

var _ log.Logger = (*Logger)(nil)

type Logger struct{}

func NewLogger() log.Logger {
	return &Logger{}
}

func (l *Logger) Debug(_ string, _ ...any) {}
func (l *Logger) Info(_ string, _ ...any)  {}
func (l *Logger) Warn(_ string, _ ...any)  {}
func (l *Logger) Error(_ string, _ ...any) {}
func (l *Logger) With(_ ...any) log.Logger { return l }
func (l *Logger) Impl() any                { return l }

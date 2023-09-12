package xlog

import (
	"errors"
	"strings"
	"sync/atomic"
)

var defaultLogger atomic.Value

func DefaultLogger() *Logger {
	if l := defaultLogger.Load(); l != nil {
		return l.(*Logger)
	}
	SetDefaultLogger(NewLogger("root"))
	return defaultLogger.Load().(*Logger)
}

func SetDefaultLogger(l *Logger) {
	defaultLogger.Store(l)
}

type Logger struct {
	Name      string
	Level     int
	Formatter string
	Handlers  []Handler
	Parent    *Logger
}

func NewLogger(name string) *Logger {
	return &Logger{
		Name:     name,
		Level:    LevelInfo,
		Handlers: []Handler{},
	}
}

func (l *Logger) CreateChild(name string) *Logger {
	logger := NewLogger(name)
	logger.Parent = l
	return logger
}

func (l *Logger) SetFormatter(s string) {
	l.Formatter = s
	for _, h := range l.Handlers {
		h.SetFormatter(l.Formatter)
	}
}

func (l *Logger) AddHandler(h Handler) {
	h.InitFormatter(l.Formatter)
	l.Handlers = append(l.Handlers, h)
}

func (l *Logger) ClearHandlers() {
	l.Handlers = []Handler{}
}

func (l *Logger) Log(r Record) error {
	if l.Parent == nil && len(l.Handlers) == 0 {
		l.AddHandler(NewStdoutHandler())
	}

	errMsgs := []string{}

	// handle record with loaded handlers
	if r.Level >= l.Level {
		for _, h := range l.Handlers {
			if err := h.HandleRecord(r); err != nil {
				errMsgs = append(errMsgs, err.Error())
			}
		}
	}

	// propagate to parent logger
	if l.Parent != nil {
		if err := l.Parent.Log(r); err != nil {
			errMsgs = append(errMsgs, err.Error())
		}
	}

	return errors.New(strings.Join(errMsgs, ", "))
}

func (l *Logger) Info(msg string, args ...any) error {
	return l.Log(NewRecord(l.Name, LevelInfo, msg, args...))
}

func (l *Logger) Warn(msg string, args ...any) error {
	return l.Log(NewRecord(l.Name, LevelWarn, msg, args...))
}

func (l *Logger) Error(msg string, args ...any) error {
	return l.Log(NewRecord(l.Name, LevelError, msg, args...))
}

func (l *Logger) Fatal(msg string, args ...any) error {
	return l.Log(NewRecord(l.Name, LevelFatal, msg, args...))
}

func (l *Logger) Panic(msg string, args ...any) error {
	return l.Log(NewRecord(l.Name, LevelPanic, msg, args...))
}

func (l *Logger) Debug(msg string, args ...any) error {
	return l.Log(NewRecord(l.Name, LevelDebug, msg, args...))
}

func (l *Logger) Trace(msg string, args ...any) error {
	return l.Log(NewRecord(l.Name, LevelTrace, msg, args...))
}

func (l *Logger) ExtraTrace(msg string, args ...any) error {
	return l.Log(NewRecord(l.Name, LevelExtraTrace, msg, args...))
}

func (l *Logger) DeepTrace(msg string, args ...any) error {
	return l.Log(NewRecord(l.Name, LevelDeepTrace, msg, args...))
}

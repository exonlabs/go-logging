package xlog

import (
	"sync/atomic"
)

var (
	defaultFormatter  atomic.Value
	defaultTimeFormat atomic.Value
)

func init() {
	SetDefaultFormatter("{time} {level} {message}")
	SetDefaultTimeFormat("2006-01-02 15:04:05.000000")
}

func DefaultFormatter() string {
	return defaultFormatter.Load().(string)
}

func SetDefaultFormatter(s string) {
	defaultFormatter.Store(s)
}

func DefaultTimeFormat() string {
	return defaultTimeFormat.Load().(string)
}

func SetDefaultTimeFormat(s string) {
	defaultTimeFormat.Store(s)
}

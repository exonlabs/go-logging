package xlog

import "sync/atomic"

// logging levels
const (
	LevelDeepTrace  = -40
	LevelExtraTrace = -30
	LevelTrace      = -20
	LevelDebug      = -10
	LevelInfo       = 0
	LevelWarn       = 10
	LevelError      = 20
	LevelFatal      = 30
	LevelPanic      = 40
)

var defaultLevelMapper atomic.Value

func init() {
	SetDefaultLevelMapper(LevelToText)
}

func DefaultLevelMapper() func(int) string {
	return defaultLevelMapper.Load().(func(int) string)
}

func SetDefaultLevelMapper(m func(int) string) {
	defaultLevelMapper.Store(m)
}

// returns text representation for log level
func LevelToText(level int) string {
	switch level {
	case LevelTrace, LevelExtraTrace, LevelDeepTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelPanic:
		return "PANIC"
	default:
		return ""
	}
}

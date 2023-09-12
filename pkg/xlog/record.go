package xlog

import (
	"fmt"
	"strings"
	"time"
)

type Record struct {
	Timestamp time.Time
	Source    string
	Level     int
	Message   string
	MsgArgs   []any
}

// create new logging record
func NewRecord(src string, level int, msg string, args ...any) Record {
	return Record{
		Timestamp: time.Now().Local(),
		Source:    src,
		Level:     level,
		Message:   msg,
		MsgArgs:   args,
	}
}

func (r Record) Format(formatter string) string {
	if formatter == "" {
		formatter = DefaultFormatter()
	}
	return strings.NewReplacer(
		"{time}", r.Timestamp.Format(DefaultTimeFormat()),
		"{source}", r.Source,
		"{level}", DefaultLevelMapper()(r.Level),
		"{message}", fmt.Sprintf(r.Message, r.MsgArgs...),
	).Replace(formatter)
}

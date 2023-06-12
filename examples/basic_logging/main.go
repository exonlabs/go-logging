package main

import (
	"fmt"

	"github.com/exonlabs/go-logging/pkg/xlog"
)

func log_messages(logger *xlog.Logger) {
	logger.Trace("logging message type: %s", "trace")
	logger.Debug("logging message type: %s", "debug")
	logger.Info("logging message type: %s", "info")
	logger.Warn("logging message type: %s", "warn")
	logger.Error("logging message type: %s", "error")
	logger.Fatal("logging message type: %s", "fatal")
	logger.Panic("logging message type: %s", "panic")
}

func main() {
	logger := xlog.DefaultLogger()

	fmt.Println("\n* logging level: TRACE")
	logger.Level = xlog.LevelTrace
	log_messages(logger)

	fmt.Println("\n* logging level: DEBUG")
	logger.Level = xlog.LevelDebug
	log_messages(logger)

	fmt.Println("\n-- with custom level mapper --")
	xlog.SetDefaultLevelMapper(func(level int) string {
		return fmt.Sprintf("%-5s", xlog.LevelToText(level))
	})

	fmt.Println("\n* logging level: INFO")
	logger.Level = xlog.LevelInfo
	log_messages(logger)

	fmt.Println("\n* logging level: WARN")
	logger.Level = xlog.LevelWarn
	log_messages(logger)

	fmt.Println("\n-- with custom formatter and time format --")
	xlog.SetDefaultFormatter("({time}) {level} [{source}] {message}")
	xlog.SetDefaultTimeFormat("2006/01/02 15:04:05.000000")

	fmt.Println("\n* logging level: ERROR")
	logger.Level = xlog.LevelError
	log_messages(logger)

	fmt.Println("\n* logging level: FATAL")
	logger.Level = xlog.LevelFatal
	log_messages(logger)

	fmt.Println("\n* logging level: PANIC")
	logger.Level = xlog.LevelPanic
	log_messages(logger)

	fmt.Println()
}

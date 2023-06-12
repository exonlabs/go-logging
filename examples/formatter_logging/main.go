package main

import (
	"fmt"

	"github.com/exonlabs/go-logging/pkg/xlog"
)

func log_messages(logger *xlog.Logger) {
	logger.Debug("logging message type: %s", "debug")
	logger.Info("logging message type: %s", "info")
	logger.Warn("logging message type: %s", "warn")
	logger.Error("logging message type: %s", "error")
	logger.Fatal("logging message type: %s", "fatal")
}

func main() {
	logger := xlog.DefaultLogger()
	logger.Level = xlog.LevelDebug

	fmt.Println("\n* with default formatter:", xlog.DefaultFormatter())
	log_messages(logger)

	logger.SetFormatter("{time} {level} [{source}] {message}")
	fmt.Println("\n* logging formatter:", logger.Formatter)
	log_messages(logger)

	hnd2 := xlog.NewStdoutHandler()
	hnd2.SetFormatter("{time} {level} -- {message} -- (hnd2)")
	logger.AddHandler(hnd2)

	fmt.Println("\n* logging with 2 handlers")
	log_messages(logger)

	fmt.Println()
}

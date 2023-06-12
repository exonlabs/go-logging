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
	logger.SetFormatter(
		"{time} {level} [{source}] root handler -- {message}")

	fmt.Println("\n* logging parent logger:", logger.Name)
	log_messages(logger)

	log1 := xlog.NewLogger("child1")
	log1.Parent = logger
	log1.Level = xlog.LevelInfo
	fmt.Println("\n* logging child logger:", log1.Name)
	log_messages(log1)

	log2 := xlog.NewLogger("child2")
	log2.Parent = logger
	log2.Level = xlog.LevelWarn
	log2.SetFormatter(
		"{time} {level} ({source}) child2 handler +++ {message}")
	log2.AddHandler(xlog.NewStdoutHandler())
	fmt.Println("\n* logging child logger (+handlers):", log2.Name)
	log_messages(log2)

	log21 := xlog.NewLogger("child21")
	log21.Parent = log2
	log21.Level = xlog.LevelError
	log21.SetFormatter(
		"{time} {level} ({source}) child21 handler +++ {message}")
	log21.AddHandler(xlog.NewStdoutHandler())
	fmt.Println("\n* logging subchild logger (+handlers):", log21.Name)
	log_messages(log21)

	fmt.Println()
}

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
	logger.SetFormatter("{time} [{source}] {level} -- {message}")

	hnd1 := xlog.NewStdoutHandler()
	logger.AddHandler(hnd1)

	hnd2 := xlog.NewFileHandler("/tmp/foobar.log")
	logger.AddHandler(hnd2)

	fmt.Println("\n* logging stdout and file:", hnd2.FilePath)
	log_messages(logger)

	fmt.Println()
}

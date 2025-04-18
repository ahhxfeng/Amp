/*
* main logger of the project use log/slog
* > stdout, > file.log
 */
package log

import (
	"log/slog"
	"os"
)

func InitLogger(filename string) *slog.Logger {
	// set the log level to info
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("failed to open log file", "error", err)
		return nil
	}
	// create a new logger text handler
	handler := slog.NewTextHandler(f, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})

	// create a new logger
	logger := slog.New(handler)
	// set the logger to the global logger
	slog.SetDefault(logger)
	logger.Info("logger initialized", "filename", filename)

	return logger
}

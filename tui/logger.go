package tui

import (
	"os"

	"github.com/charmbracelet/log"
)

var baseLogger = newBaseLogger()

func newBaseLogger() *log.Logger {
	// will creation standard file io lib later
	f, _ := os.Create("debug.log")
	return log.NewWithOptions(f, log.Options{
		Level: log.DebugLevel,
	})
}

func LoggerFor(component string) *log.Logger {
	return baseLogger.With("component", component)
}

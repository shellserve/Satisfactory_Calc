package tui

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/shellserve/Satisfactory_Calc/internal/domain/file"
)

const debugFile = "debug.log"

var baseLogger, err = newBaseLogger()

func newBaseLogger() (*log.Logger, error) {
	var (
		f   *os.File
		err error
	)

	if !file.FileExists(debugFile) {
		if f, err = os.Create(debugFile); err != nil {
			return nil, err
		}
	}

	return log.NewWithOptions(f, log.Options{
		Level: log.DebugLevel,
	}), nil
}

func LoggerFor(component string) *log.Logger {
	if err != nil {
		panic("Logger failed to initalize!")
	}
	return baseLogger.With("component", component)
}

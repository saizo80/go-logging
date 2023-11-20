package gologging

import (
	"testing"
)

func TestDebug(t *testing.T) {
	log := New(DEBUG, Option{FilePath: "test.log"})
	log.Debug("This is a debug message")
	errorCode := 100
	log.Debug("This is a debug message with error code: %d", errorCode)

	log = New(INFO, Option{FilePath: "test.log"})
	log.Debug("This debug message should not be printed to stdout, but should be printed to file")

	log = New(DEBUG, Option{Stdout: false})
	log.Debug("This debug message should not be printed")
}

func TestInfo(t *testing.T) {
	log := New(INFO)
	log.Info("This is a info message")
	errorCode := 100
	log.Info("This is a info message with error code: %d", errorCode)

	log = New(WARN)
	log.Info("This info message should not be printed")

	log = New(INFO, Option{Stdout: false})
	log.Info("This info message should not be printed")
}

func TestWarn(t *testing.T) {
	log := New(WARN)
	log.Warn("This is a warn message")
	errorCode := 100
	log.Warn("This is a warn message with error code: %d", errorCode)

	log = New(ERROR)
	log.Warn("This warn message should not be printed")

	log = New(WARN, Option{Stdout: false})
	log.Warn("This warn message should not be printed")
}

func TestError(t *testing.T) {
	log := New(ERROR)
	log.Error("This is a error message")
	errorCode := 100
	log.Error("This is a error message with error code: %d", errorCode)

	log = New(ERROR, Option{Stdout: false})
	log.Error("This error message should not be printed")
}

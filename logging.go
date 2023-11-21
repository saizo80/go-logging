package gologging

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	reset  = "\033[0m"
	yellow = "\033[0;33m"
	blue   = "\033[0;34m"
	gray   = "\033[0;95m"
	red    = "\033[0;31m"
	DEBUG  = 0
	INFO   = 10
	WARN   = 20
	ERROR  = 30
	FATAL  = 40
)

type Logger struct {
	level    int
	filePath string
	stdout   bool
}

type Option struct {
	FilePath string
	Stdout   bool
}

func New(level int, option ...Option) *Logger {
	stdout := true
	filePath := ""
	if len(option) > 0 {
		stdout = option[0].Stdout
		filePath = option[0].FilePath
	}
	return &Logger{
		level:    level,
		filePath: filePath,
		stdout:   stdout,
	}
}

func (l *Logger) print(level string, message interface{}, color string, args ...interface{}) {
	now := time.Now()
	if _, ok := message.(string); !ok {
		message = fmt.Sprintf("%v", message)
	}
	if len(args) > 0 {
		message = fmt.Sprintf(message.(string), args...)
	}
	fmt.Printf("[%s%s%s] %s %s\n",
		color,
		level,
		reset,
		now.Format("15:04:05,000"),
		message,
	)
}

func (l *Logger) printToFile(level string, message interface{}, args ...interface{}) {
	now := time.Now()
	if _, ok := message.(string); !ok {
		message = fmt.Sprintf("%v", message)
	}
	if len(args) > 0 {
		message = fmt.Sprintf(message.(string), args...)
	}
	fileObj, err := os.OpenFile(l.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer fileObj.Close()

	// get the file name that is calling this function
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fileName := filepath.Base(exe)

	_, err = fmt.Fprintf(fileObj, "%s - %s - %s - %s\n",
		now.Format("2006-01-02 15:04:05,000"),
		fileName,
		level,
		message,
	)
	if err != nil {
		panic(err)
	}
}

// set log level
func (l *Logger) SetLevel(level int) {
	l.level = level
}

// set log file path
func (l *Logger) SetFilePath(filePath string) {
	l.filePath = filePath
}

// set stdout printing flag
func (l *Logger) SetStdout(stdout bool) {
	l.stdout = stdout
}

// debug log
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	if l.level <= DEBUG {
		if l.stdout {
			l.print("DEBUG", message, gray, args...)
		}
	}
	if l.filePath != "" {
		l.printToFile("DEBUG", message, args...)
	}
}

// info log
func (l *Logger) Info(message interface{}, args ...interface{}) {
	if l.level <= INFO && l.stdout {
		l.print("INFO", message, blue, args...)
	}
	if l.filePath != "" {
		l.printToFile("INFO", message, args...)
	}
}

// warn log
func (l *Logger) Warn(message interface{}, args ...interface{}) {
	if l.level <= WARN && l.stdout {
		l.print("WARN", message, yellow, args...)
	}
	if l.filePath != "" {
		l.printToFile("WARN", message, args...)
	}
}

// error log
func (l *Logger) Error(message interface{}, args ...interface{}) {
	if l.level <= ERROR && l.stdout {
		l.print("ERROR", message, red, args...)
	}
	if l.filePath != "" {
		l.printToFile("ERROR", message, args...)
	}
}

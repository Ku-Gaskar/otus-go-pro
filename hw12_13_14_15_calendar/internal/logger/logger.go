package logger

import (
	"log"
	"os"
	"strings"
)

type Logger struct {
	logger     *log.Logger
	logLevel   int
	logEnabled bool
}

const (
	LevelError = iota
	LevelWarn
	LevelInfo
	LevelDebug
)

// New создает новый логгер с заданным уровнем логирования и настройками.

func New(Enabled bool, level string, prefix string) *Logger {
	logLevel := parseLogLevel(level)

	return &Logger{
		logger:     log.New(os.Stdout, prefix, log.LstdFlags),
		logLevel:   logLevel,
		logEnabled: Enabled,
	}
}

func parseLogLevel(level string) int {
	switch strings.ToLower(level) {
	case "error":
		return LevelError
	case "warn":
		return LevelWarn
	case "info":
		return LevelInfo
	case "debug":
		return LevelDebug
	default:
		log.Printf("Неизвестный уровень логирования: %s, используется 'info'", level)
		return LevelInfo
	}
}

// Error логирует сообщение уровня ERROR.
func (l *Logger) Error(v ...interface{}) {
	if l.logEnabled && l.logLevel >= LevelError {
		l.logger.SetPrefix("[ERROR] ")
		l.logger.Println(v...)
	}
}

// Warn логирует сообщение уровня WARN.
func (l *Logger) Warn(v ...interface{}) {
	if l.logEnabled && l.logLevel >= LevelWarn {
		l.logger.SetPrefix("[WARN] ")
		l.logger.Println(v...)
	}
}

// Info логирует сообщение уровня INFO.
func (l *Logger) Info(v ...interface{}) {
	if l.logEnabled && l.logLevel >= LevelInfo {
		l.logger.SetPrefix("[INFO] ")
		l.logger.Println(v...)
	}
}

// Debug логирует сообщение уровня DEBUG.
func (l *Logger) Debug(v ...interface{}) {
	if l.logEnabled && l.logLevel >= LevelDebug {
		l.logger.SetPrefix("[DEBUG] ")
		l.logger.Println(v...)
	}
}

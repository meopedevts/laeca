package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct{}

var Log *Logger

func InitLogger() {
	Log = &Logger{}
}

func (l *Logger) logMessage(level LogLevel, msg string, args ...any) {
	var style lipgloss.Style
	var prefix string

	switch level {
	case DEBUG:
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#4CAF50"))
		prefix = "[DEBUG]"
	case INFO:
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#2196F3"))
		prefix = "[INFO]"
	case WARN:
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF9800"))
		prefix = "[WARN]"
	case ERROR:
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#F44336"))
		prefix = "[ERROR]"
	case FATAL:
		style = lipgloss.NewStyle().Foreground(lipgloss.Color("#D32F2F"))
		prefix = "[FATAL]"
	}

	message := fmt.Sprintf(msg, args...)
	log.Printf("%s %s\n", style.Render(prefix), message)
}

func Debug(msg string, args ...any) {
	Log.logMessage(DEBUG, msg, args...)
}

func Info(msg string, args ...any) {
	Log.logMessage(INFO, msg, args...)
}

func Warn(msg string, args ...any) {
	Log.logMessage(WARN, msg, args...)
}

func Error(msg string, args ...any) {
	Log.logMessage(ERROR, msg, args...)
}

func Fatal(msg string, args ...any) {
	Log.logMessage(FATAL, msg, args...)
	os.Exit(1)
}

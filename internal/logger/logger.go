package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
)

// LogLevel representa os níveis de log
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Logger é a estrutura principal para o logger customizado
type Logger struct {
	// Level LogLevel
}

// Instância global do logger
var Log *Logger

// InitLogger inicializa o logger
func InitLogger() {
	Log = &Logger{}
}

func (l *Logger) logMessage(level LogLevel, msg string, args ...any) {
	// if level < l.Level {
	// 	return
	// }

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

// Debug imprime uma mensagem de debug
func Debug(msg string, args ...any) {
	Log.logMessage(DEBUG, msg, args...)
}

// Info imprime uma mensagem informativa
func Info(msg string, args ...any) {
	Log.logMessage(INFO, msg, args...)
}

// Warn imprime uma mensagem de aviso
func Warn(msg string, args ...any) {
	Log.logMessage(WARN, msg, args...)
}

// Error imprime uma mensagem de erro
func Error(msg string, args ...any) {
	Log.logMessage(ERROR, msg, args...)
}

// Fatal imprime uma mensagem de erro fatal
func Fatal(msg string, args ...any) {
	Log.logMessage(FATAL, msg, args...)
	os.Exit(1) // Encerra o programa após logar a mensagem fatal
}

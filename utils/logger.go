package utils

import (
	"log"
	"os"
	"time"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *Logger) Info(message string, data ...map[string]interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	if len(data) > 0 && data[0] != nil {
		l.Printf("[%s] INFO: %s | %v\n", timestamp, message, data[0])
	} else {
		l.Printf("[%s] INFO: %s\n", timestamp, message)
	}
}

func (l *Logger) Error(message string, data ...map[string]interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	if len(data) > 0 && data[0] != nil {
		l.Printf("[%s] ERROR: %s | %v\n", timestamp, message, data[0])
	} else {
		l.Printf("[%s] ERROR: %s\n", timestamp, message)
	}
}

func (l *Logger) Warn(message string, data ...map[string]interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	if len(data) > 0 && data[0] != nil {
		l.Printf("[%s] WARN: %s | %v\n", timestamp, message, data[0])
	} else {
		l.Printf("[%s] WARN: %s\n", timestamp, message)
	}
}

func (l *Logger) Debug(message string, data ...map[string]interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	if len(data) > 0 && data[0] != nil {
		l.Printf("[%s] DEBUG: %s | %v\n", timestamp, message, data[0])
	} else {
		l.Printf("[%s] DEBUG: %s\n", timestamp, message)
	}
}

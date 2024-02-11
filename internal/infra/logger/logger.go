package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	*log.Logger
}

func New() *Logger {
	return &Logger{log.Default()}
}

func (logger *Logger) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Error]: %s\n", msg)
}

func (logger *Logger) LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Info]: %s\n", msg)
}

package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)

	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (logger *Logger) Debug(v ...any) {
	logger.debug.Println(v...)
}

func (logger *Logger) Info(v ...any) {
	logger.info.Println(v...)
}

func (logger *Logger) Warn(v ...any) {
	logger.warning.Println(v...)
}

func (logger *Logger) Error(v ...any) {
	logger.err.Println(v...)
}

func (logger *Logger) Debugf(format string, v ...any) {
	logger.debug.Printf(format, v...)
}

func (logger *Logger) Infof(format string, v ...any) {
	logger.info.Printf(format, v...)
}

func (logger *Logger) Warnf(format string, v ...any) {
	logger.warning.Printf(format, v...)
}

func (logger *Logger) Errorf(format string, v ...any) {
	logger.err.Printf(format, v...)
}

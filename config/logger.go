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

func NewLogger() *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{}
}

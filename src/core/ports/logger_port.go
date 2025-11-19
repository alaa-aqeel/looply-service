package ports

import (
	"time"
)

type Level int8

const (
	Debug Level = iota
	Info
	Warn
	Error
)

type Logger struct {
	Message  string
	Tag      string
	Error    error
	Args     any
	Duration time.Duration
}

type LoggerPort interface {
	Log(level Level, log Logger)
}

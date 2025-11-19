package logger

import (
	"sync"

	"github.com/alaa-aqeel/looply-app/src/core/ports"
)

var (
	Log     ports.LoggerPort
	once    sync.Once
	errInit error
)

func NewLogger() (ports.LoggerPort, error) {
	once.Do(func() {
		Log, errInit = newZapLogger()
	})
	return Log, errInit
}

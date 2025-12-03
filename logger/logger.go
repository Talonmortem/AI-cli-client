package logger

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

var logg *log.Logger

// Init инициализация глобального логгера
func Init(filePath string, debug bool) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	logg = log.New()
	logg.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	output := io.Writer(file)
	if debug {
		output = io.MultiWriter(os.Stdout, file)
	}

	logg.SetOutput(output)

	if debug {
		logg.SetLevel(log.DebugLevel)
	} else {
		logg.SetLevel(log.InfoLevel)
	}

	return nil
}

// ---------------- ModuleLogger ----------------

type ModuleLogger struct {
	module string
}

// For возвращает логгер для конкретного модуля
func For(module string) *ModuleLogger {
	return &ModuleLogger{module: module}
}

// Println, Printf — Info
func (m *ModuleLogger) Println(v ...any) {
	if logg != nil {
		logg.Infoln(append([]any{fmt.Sprintf("[%s]", m.module)}, v...)...)
	}
}

func (m *ModuleLogger) Printf(format string, v ...any) {
	if logg != nil {
		format = fmt.Sprintf("[%s] %s", m.module, format)
		logg.Infof(format, v...)
	}
}

// Debugln, Debugf
func (m *ModuleLogger) Debugln(v ...any) {
	if logg != nil {
		logg.Debugln(append([]any{fmt.Sprintf("[%s]", m.module)}, v...)...)
	}
}

func (m *ModuleLogger) Debugf(format string, v ...any) {
	if logg != nil {
		format = fmt.Sprintf("[%s] %s", m.module, format)
		logg.Debugf(format, v...)
	}
}

// Warnln, Warnf
func (m *ModuleLogger) Warnln(v ...any) {
	if logg != nil {
		logg.Warnln(append([]any{fmt.Sprintf("[%s]", m.module)}, v...)...)
	}
}

func (m *ModuleLogger) Warnf(format string, v ...any) {
	if logg != nil {
		format = fmt.Sprintf("[%s] %s", m.module, format)
		logg.Warnf(format, v...)
	}
}

// Errorln, Errorf
func (m *ModuleLogger) Errorln(v ...any) {
	if logg != nil {
		logg.Errorln(append([]any{fmt.Sprintf("[%s]", m.module)}, v...)...)
	}
}

func (m *ModuleLogger) Errorf(format string, v ...any) {
	if logg != nil {
		format = fmt.Sprintf("[%s] %s", m.module, format)
		logg.Errorf(format, v...)
	}
}

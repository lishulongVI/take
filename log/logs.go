package log

import (
	"fmt"
	"io"
	stdlog "log"
	"strings"
)

const (
	Off = iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
)

type Logger struct {
	level  int
	logger *stdlog.Logger
}

var loggers []*Logger

var logLevel = Debug

func NewLogger(out io.Writer) *Logger {
	ret := &Logger{level: logLevel, logger: stdlog.New(out, "", stdlog.Ldate|stdlog.Ltime|stdlog.Lshortfile)}
	loggers = append(loggers, ret)
	return ret
}

func SetLevel(level string) {
	logLevel = getLenvel(level)

	for _, l := range loggers {
		l.SetLevel(level)
	}
}

func getLenvel(level string) int {

	level = strings.ToLower(level)

	switch level {
	case "off":
		return Off
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Info
	}

}

func (l *Logger) SetLevel(level string) {
	l.level = getLenvel(level)
}

func (l *Logger) IsTraceEnabled() bool {
	return l.level <= Trace
}
func (l *Logger) IsDebugEnabled() bool {
	return l.level <= Debug
}
func (l *Logger) IsWarnEnabled() bool {
	return l.level <= Warn
}

func (l *Logger) Trace(v ...interface{}) {
	if Trace < l.level {
		return
	}
	l.logger.SetPrefix("T ")
	l.logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	if Trace < l.level {
		return
	}
	l.logger.SetPrefix("T ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}
func (l *Logger) Debug(v ...interface{}) {
	if Debug < l.level {
		return
	}
	l.logger.SetPrefix("D ")
	l.logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if Debug < l.level {
		return
	}
	l.logger.SetPrefix("D ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}
func (l *Logger) Info(v ...interface{}) {
	if Info < l.level {
		return
	}
	l.logger.SetPrefix("I ")
	l.logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if Info < l.level {
		return
	}
	l.logger.SetPrefix("I ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}
func (l *Logger) Warn(v ...interface{}) {
	if Warn < l.level {
		return
	}
	l.logger.SetPrefix("W ")
	l.logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if Warn < l.level {
		return
	}
	l.logger.SetPrefix("W ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}
func (l *Logger) Error(v ...interface{}) {
	if Error < l.level {
		return
	}
	l.logger.SetPrefix("E ")
	l.logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if Error < l.level {
		return
	}
	l.logger.SetPrefix("E ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}
func (l *Logger) Fatal(v ...interface{}) {
	if Fatal < l.level {
		return
	}
	l.logger.SetPrefix("F ")
	l.logger.Output(2, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if Fatal < l.level {
		return
	}
	l.logger.SetPrefix("F ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

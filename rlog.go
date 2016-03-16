// Package rlog implements a simple logging package. It provides the
// following LogWriter engins: syslog, stdout, stderr
// The definition of log level follows the spec of syslog.
package rlog

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// Severity, refer to syslog
const (
	EMERG  = 0
	ALERT  = 1
	CRIT   = 2
	ERROR  = 3
	WARN   = 4
	NOTICE = 5
	INFO   = 6
	DEBUG  = 7
)

// Logger flags
const (
	PREFIX = 1 << iota
	JSON
)

// Logger represents an instance of log settings
type Logger struct {
	Tag        string    // the app name
	Level      int       // the log level
	WriterType string    // the LogWriter type
	Writer     LogWriter // the LogWriter instance
	Flag       int       // the flag bitmap
}

// LogWriter is an engine which provides logln function
type LogWriter interface {
	logln(level int, msg string)
}

// Severity is the string map for prefix
var Severity = map[int]string{
	EMERG:  "EMERG ",
	ALERT:  "ALERT ",
	CRIT:   "CRIT  ",
	ERROR:  "ERROR ",
	WARN:   "WARN  ",
	NOTICE: "NOTICE",
	INFO:   "INFO  ",
	DEBUG:  "DEBUG ",
}

// NewLogger creates a Logger. The writerType specifies the LogWriter egnine.
// The tag is usually used as the program name. The logs of level higher then
// level will be ignored. The flag can be used to format the log message.
func NewLogger(writerType, tag string, level int, flag int) (*Logger, error) {
	var err error
	logger := Logger{Tag: tag, Level: level, WriterType: writerType, Flag: flag}

	switch writerType {
	case "syslog":
		logger.Writer, err = newSyslogWriter(tag)
	case "stdout":
		logger.Writer, err = newIoWriter(os.Stdout)
	case "stderr":
		logger.Writer, err = newIoWriter(os.Stderr)
	default:
		return nil, errors.New("Invalid writer type" + writerType)
	}

	if err != nil {
		return nil, err
	}
	return &logger, nil
}

// NewIoLogger provides the flexibility that users can provide their io.Writer
func NewIoLogger(w io.Writer, tag string, level int, flag int) (*Logger, error) {
	var err error
	logger := Logger{Tag: tag, Level: level, WriterType: "io", Flag: flag}

	logger.Writer, err = newIoWriter(w)

	if err != nil {
		return nil, err
	}
	return &logger, nil
}

// SetLevel sets Logger level
func (l *Logger) SetLevel(level int) {
	l.Level = level
}

// SetFlag sets Logger flag
func (l *Logger) SetFlag(flag int) {
	l.Flag = flag
}

// Emergf logs formatted messages of emergency level. Newline is appended.
func (l *Logger) Emergf(format string, v ...interface{}) {
	l.logfln(EMERG, format, v...)
}

// Alertf logs formatted messages of alert level. Newline is appended.
func (l *Logger) Alertf(format string, v ...interface{}) {
	l.logfln(ALERT, format, v...)
}

// Critf logs formatted messages of critical level. Newline is appended.
func (l *Logger) Critf(format string, v ...interface{}) {
	l.logfln(CRIT, format, v...)
}

// Errorf logs formatted messages of error level. Newline is appended.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logfln(ERROR, format, v...)
}

// Warnf logs formatted messages of warning level. Newline is appended.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.logfln(WARN, format, v...)
}

// Noticef logs formatted messages of notice level. Newline is appended.
func (l *Logger) Noticef(format string, v ...interface{}) {
	l.logfln(NOTICE, format, v...)
}

// Infof logs formatted messages of information level. Newline is appended.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.logfln(INFO, format, v...)
}

// Debugf logs formatted messages of debug level. Newline is appended.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logfln(DEBUG, format, v...)
}

// Emergln logs messages of Emergency level. Newline is appended
func (l *Logger) Emergln(v ...interface{}) {
	l.logln(EMERG, v...)
}

// Alertln logs messages of Alert level. Newline is appended
func (l *Logger) Alertln(v ...interface{}) {
	l.logln(ALERT, v...)
}

// Critln logs messages of critical level. Newline is appended
func (l *Logger) Critln(v ...interface{}) {
	l.logln(CRIT, v...)
}

// Errorln logs messages of error level. Newline is appended
func (l *Logger) Errorln(v ...interface{}) {
	l.logln(ERROR, v...)
}

// Warnln logs messages of warning level. Newline is appended
func (l *Logger) Warnln(v ...interface{}) {
	l.logln(WARN, v...)
}

// Noticeln logs messages of notice level. Newline is appended
func (l *Logger) Noticeln(v ...interface{}) {
	l.logln(NOTICE, v...)
}

// Infoln logs messages of information level. Newline is appended
func (l *Logger) Infoln(v ...interface{}) {
	l.logln(INFO, v...)
}

// Deubgln logs messages of debug level. Newline is appended
func (l *Logger) Debugln(v ...interface{}) {
	l.logln(DEBUG, v...)
}

// logfln generates the message by fmt.Sprintf if level should be outputed
func (l *Logger) logfln(level int, format string, v ...interface{}) {
	if l.Level < level {
		return
	}

	l.logMsg(level, fmt.Sprintf(format, v...))
}

// logln generates the message by fmt.Sprint if level should be outputed
func (l *Logger) logln(level int, v ...interface{}) {
	if l.Level < level {
		return
	}

	l.logMsg(level, strings.TrimRight(fmt.Sprintln(v...), "\n"))
}

// logMsg uses the LoggerWriter to output the message. The message may be
// modified according to the flag setting.
func (l *Logger) logMsg(level int, msg string) {
	if l.Flag&JSON != 0 {
		msg = `{"t":"slog","msg":"` + msg + `"}`
	}

	if l.Flag&PREFIX != 0 {
		timestamp := time.Now().Format(time.Stamp)
		msg = fmt.Sprintf("%s|%s|%s|%s", timestamp, l.Tag, Severity[level], msg)
	}

	l.Writer.logln(level, msg)
}

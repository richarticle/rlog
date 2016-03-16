package rlog

import "log/syslog"

// syslogWriter implements LogWritter interface. syslog is the ouptut.
type syslogWriter struct {
	writer *syslog.Writer
}

// newSyslogWriter creates a syslog writer
func newSyslogWriter(tag string) (*syslogWriter, error) {
	w, err := syslog.New(syslog.LOG_INFO|syslog.LOG_USER, tag)
	if err != nil {
		return nil, err
	}

	return &syslogWriter{w}, nil
}

// logln uses syslog.Writer to output the message
func (sw *syslogWriter) logln(level int, msg string) {
	switch level {
	case EMERG:
		sw.writer.Emerg(msg)
	case ALERT:
		sw.writer.Alert(msg)
	case CRIT:
		sw.writer.Crit(msg)
	case ERROR:
		sw.writer.Err(msg)
	case WARN:
		sw.writer.Warning(msg)
	case NOTICE:
		sw.writer.Notice(msg)
	case INFO:
		sw.writer.Info(msg)
	case DEBUG:
		sw.writer.Debug(msg)
	default:
		return
	}
}

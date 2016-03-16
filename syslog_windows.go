package rlog

import "fmt"

// syslogWriter implements LogWritter interface. syslog is the ouptut.
type syslogWriter struct {
}

// newSyslogWriter creates a syslog writer
func newSyslogWriter(tag string) (*syslogWriter, error) {
	return &syslogWriter{}, nil
}

// logln uses syslog.Writer to output the message
func (sw *syslogWriter) logln(level int, msg string) {
	fmt.Println(msg)
}

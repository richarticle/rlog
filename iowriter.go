package rlog

import (
	"fmt"
	"io"
)

// ioWriter implements LogWritter interface. io.Writer is the ouptut.
type ioWriter struct {
	writer io.Writer
}

// newIoWriter creates an IO writer
func newIoWriter(w io.Writer) (*ioWriter, error) {
	return &ioWriter{w}, nil
}

// logln writes the message with newline
func (iw *ioWriter) logln(level int, msg string) {
	fmt.Fprintln(iw.writer, msg)
}

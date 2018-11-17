package io

import (
	_io "io"
	"time"
)

func TrySetReadDeadline(t time.Time, s _io.Reader) {
	if sd, ok := s.(DeadlineReader); ok {
		sd.SetReadDeadline(t)
	}
}

func TrySetWriteDeadline(t time.Time, s _io.Writer) {
	if sd, ok := s.(DeadlineWriter); ok {
		sd.SetWriteDeadline(t)
	}
}

func CopyDeadline(time time.Time, dst _io.Writer, src _io.Reader) (int64, error) {
	TrySetReadDeadline(time, src)
	TrySetWriteDeadline(time, dst)
	return _io.Copy(dst, src)
}

package io

import (
	_io "io"
	"time"
)

func TrySetReadDeadline(t time.Time, s _io.Reader) error {
	if sd, ok := s.(DeadlineReader); ok {
		if err := sd.SetReadDeadline(t); err != nil {
			return err
		}
	}
	return nil
}

func TrySetWriteDeadline(t time.Time, s _io.Writer) error {
	if sd, ok := s.(DeadlineWriter); ok {
		if err := sd.SetWriteDeadline(t); err != nil {
			return err
		}
	}
	return nil
}

func CopyDeadline(time time.Time, dst _io.Writer, src _io.Reader) (int64, error) {
	if err := TrySetReadDeadline(time, src); err != nil {
		return 0, err
	}
	if err := TrySetWriteDeadline(time, dst); err != nil {
		return 0, err
	}
	return _io.Copy(dst, src)
}

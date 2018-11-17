package io

import (
	_io "io"
	"time"
)

type CloseReader interface {
	_io.Reader
	CloseRead() error
}
type CloseWriter interface {
	_io.Writer
	CloseWrite() error
}
type DeadlineReader interface {
	_io.Reader
	SetReadDeadline(t time.Time) error
}
type DeadlineWriter interface {
	_io.Reader
	SetWriteDeadline(t time.Time) error
}

package io

import _io "io"

type CloseReader interface {
	_io.Reader
	CloseRead() error
}
type CloseWriter interface {
	_io.Writer
	CloseWrite() error
}

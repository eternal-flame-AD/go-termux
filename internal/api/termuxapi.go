package api

import (
	"context"
	"strings"
	"time"

	"github.com/eternal-flame-AD/go-termux/internal/intent"

	"github.com/eternal-flame-AD/go-termux/internal/unix"
)

const termuxApiComponent = "com.termux.api/.TermuxApiReceiver"

type Call struct {
	Method       string
	Args         map[string]interface{}
	Action       string
	Data         string
	pipeToRemote *unix.Socket
	pipeToMe     *unix.Socket
}

func (c *Call) Call(ctx context.Context) {
	p1, p2 := unix.NewSocket(), unix.NewSocket()
	c.pipeToMe, c.pipeToRemote = &p1, &p2

	bc := intent.Broadcast{
		User:      0,
		Component: termuxApiComponent,
		Data:      c.Data,
		Action:    c.Action,
	}
	bc.AddString("api_method", c.Method)
	bc.AddString("socket_input", strings.Trim(c.pipeToRemote.Name(), "@"))
	bc.AddString("socket_output", strings.Trim(c.pipeToMe.Name(), "@"))
	if c.Args != nil {
		for key, val := range c.Args {
			switch val.(type) {
			case string:
				bc.AddString(key, val.(string))
			case bool:
				bc.AddBool(key, val.(bool))
			case int:
				bc.AddInt(key, val.(int))
			case []int32:
				bc.AddLongs(key, val.([]int32))
			case []string:
				bc.AddStrings(key, val.([]string))
			default:
				panic("Unsupported arg type")
			}
		}
	}
	bc.Send(ctx)
}

func (c Call) SetReadDeadline(t time.Time) error {
	return c.pipeToMe.SetReadDeadline(t)
}

func (c Call) Read(p []byte) (n int, err error) {
	return c.pipeToMe.Read(p)
}

func (c Call) SetWriteDeadline(t time.Time) error {
	return c.pipeToRemote.SetWriteDeadline(t)
}

func (c Call) Write(p []byte) (n int, err error) {
	return c.pipeToRemote.Write(p)
}

func (c Call) Close() error {
	if err := c.CloseRead(); err != nil {
		return err
	}
	if err := c.CloseWrite(); err != nil {
		return err
	}
	return nil
}

func (c Call) CloseRead() error {
	return c.pipeToMe.Close()
}

func (c Call) CloseWrite() error {
	return c.pipeToRemote.Close()
}

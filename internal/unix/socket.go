package unix

import (
	"net"
	"time"

	"github.com/twinj/uuid"
)

func NewSocket() Socket {
	sock := getUnixSocket()
	l, err := net.Listen("unix", sock.Name)
	if err != nil {
		panic(err)
	}
	return Socket{
		underlyingAddr:     sock,
		underlyingListener: l,
	}
}

type Socket struct {
	underlyingAddr     net.UnixAddr
	underlyingListener net.Listener
	underlyingConn     *net.UnixConn
}

func (c *Socket) Name() string {
	return c.underlyingAddr.Name
}

func (c *Socket) connect() {
	conn, err := c.underlyingListener.Accept()
	if err != nil {
		panic(err)
	}
	c.underlyingConn = conn.(*net.UnixConn)
}

func (c *Socket) SetReadDeadline(t time.Time) error {
	return c.underlyingConn.SetReadDeadline(t)
}

func (c *Socket) SetWriteDeadline(t time.Time) error {
	return c.underlyingConn.SetWriteDeadline(t)
}

func (c *Socket) Read(p []byte) (n int, err error) {
	if c.underlyingConn == nil {
		c.connect()
	}
	return c.underlyingConn.Read(p)
}

func (c *Socket) Write(p []byte) (n int, err error) {
	if c.underlyingConn == nil {
		c.connect()
	}
	return c.underlyingConn.Write(p)
}

func (c *Socket) Close() error {
	if c.underlyingConn == nil {
		return nil
	}
	return c.underlyingConn.Close()
}

func getUnixSocket() net.UnixAddr {
	return net.UnixAddr{Name: "@" + uuid.NewV4().String(), Net: "unix"}
}

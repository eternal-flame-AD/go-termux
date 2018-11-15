package chanbuf

type BufToChan struct {
	C chan<- []byte
}

func (c BufToChan) Write(p []byte) (int, error) {
	c.C <- p
	return len(p), nil
}

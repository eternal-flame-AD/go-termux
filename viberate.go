package termux

import "bytes"

func Viberate(ms int, force bool) error {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "Viberate", map[string]interface{}{
		"duration_ms": ms,
		"force":       force,
	}, "")
	return checkErr(buf.Bytes())
}

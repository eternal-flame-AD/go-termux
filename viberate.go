package termux

import "bytes"

// Viberate creates a viberation on the device, force means whether to viberate even if the device is set to silent mode
func Viberate(ms int, force bool) error {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "Viberate", map[string]interface{}{
		"duration_ms": ms,
		"force":       force,
	}, "")
	return checkErr(buf.Bytes())
}

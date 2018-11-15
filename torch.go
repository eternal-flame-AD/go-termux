package termux

import "bytes"

func TorchEnable(enabled bool) error {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "Torch", map[string]interface{}{
		"enabled": enabled,
	}, "")
	return checkErr(buf.Bytes())
}

package termux

import "bytes"

// TorchEnable sets the current state of the device flashlight
func TorchEnable(enabled bool) error {
	buf := bytes.NewBuffer([]byte{})
	if err := exec(nil, buf, "Torch", map[string]interface{}{
		"enabled": enabled,
	}, ""); err != nil {
		return err
	}
	return checkErr(buf.Bytes())
}

package termux

import (
	"bytes"
)

// Brightness sets the current brightness level from 0 to 255
func Brightness(level uint8) error {
	buf := bytes.NewBuffer([]byte{})
	if err := exec(nil, buf, "Brightness", map[string]interface{}{
		"brightness": int(level),
	}, ""); err != nil {
		return err
	}
	res := buf.Bytes()
	return checkErr(res)
}

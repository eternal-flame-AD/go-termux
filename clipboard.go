package termux

import (
	"bytes"
)

func ClipboardGet() (string, error) {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "Clipboard", nil, "")
	return buf.String(), nil
}

func ClipboardSet(val string) error {
	inbuf := bytes.NewBufferString(val)
	buf := bytes.NewBuffer([]byte{})
	exec(inbuf, buf, "Clipboard", map[string]interface{}{
		"set":         true,
		"api_version": "2",
	}, "")
	return checkErr(buf.Bytes())
}

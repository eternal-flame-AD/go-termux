package termux

import (
	"bytes"
)

// ClipboardGet gets the current content of the clipboard
func ClipboardGet() (string, error) {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "Clipboard", nil, "")
	return buf.String(), nil
}

// ClipboardSet sets the clipboard to the given value
func ClipboardSet(val string) error {
	inbuf := bytes.NewBufferString(val)
	buf := bytes.NewBuffer([]byte{})
	exec(inbuf, buf, "Clipboard", map[string]interface{}{
		"set":         true,
		"api_version": "2",
	}, "")
	return checkErr(buf.Bytes())
}

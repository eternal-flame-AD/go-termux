package termux

import (
	"bytes"
	"path/filepath"
)

// ShareType represent the action of the share, defaults to edit
type ShareType = string

const (
	// View share type
	View ShareType = "view"
	// Edit share type
	Edit ShareType = "edit"
	// Send share type
	Send ShareType = "send"
)

func shareParam(title string, useDefault bool, actionType ShareType) map[string]interface{} {
	if actionType == "" {
		actionType = "edit"
	}
	return map[string]interface{}{
		"title":            title,
		"default-receiver": useDefault,
		"action-type":      string(actionType),
	}
}

// ShareFile shares a file with MIME type determined by the file extension, useDefault determines whether to use the default app if availables
func ShareFile(title string, path string, useDefault bool, actionType ShareType) error {
	param := shareParam(title, useDefault, actionType)
	buf := bytes.NewBuffer([]byte{})

	realpath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	param["file"] = realpath

	if err := exec(nil, buf, "Share", param, ""); err != nil {
		return err
	}
	return checkErr(buf.Bytes())
}

// Share shares raw data bytes with the given content type, useDefault determines whether to use the default app if availables
func Share(title string, data []byte, contentType string, useDefault bool, actionType ShareType) error {
	if contentType == "" {
		contentType = "text/plain"
	}
	param := shareParam(title, useDefault, actionType)
	param["content-type"] = contentType

	in := bytes.NewBuffer(data)
	buf := bytes.NewBuffer([]byte{})
	if err := exec(in, buf, "Share", param, ""); err != nil {
		return err
	}
	return checkErr(buf.Bytes())
}

package termux

import (
	"bytes"
	"path/filepath"
)

type ShareType = string

const (
	View ShareType = "view"
	Edit ShareType = "edit"
	Send ShareType = "send"
)

func shareParam(title string, useDefault bool, actionType ShareType) map[string]interface{} {
	return map[string]interface{}{
		"title":            title,
		"default-receiver": useDefault,
		"action-type":      actionType,
	}
}

func ShareFile(title string, path string, useDefault bool, actionType ShareType) error {
	param := shareParam(title, useDefault, actionType)
	buf := bytes.NewBuffer([]byte{})

	realpath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	param["file"] = realpath

	exec(nil, buf, "Share", param, "")
	return checkErr(buf.Bytes())
}

func Share(title string, data []byte, contentType string, useDefault bool, actionType ShareType) error {
	if contentType == "" {
		contentType = "text/plain"
	}
	param := shareParam(title, useDefault, actionType)
	param["content-type"] = contentType

	in := bytes.NewBuffer(data)
	buf := bytes.NewBuffer([]byte{})
	exec(in, buf, "Share", param, "")
	return checkErr(buf.Bytes())
}

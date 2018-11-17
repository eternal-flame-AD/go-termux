package termux

import (
	"bytes"
	"encoding/json"
)

// ContactPiece represents one contact info
type ContactPiece struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

// ContactList gets all contact from the device
func ContactList() ([]ContactPiece, error) {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "ContactList", nil, "")
	res := buf.Bytes()
	if err := checkErr(res); err != nil {
		return nil, err
	}
	records := make([]ContactPiece, 0)
	if err := json.Unmarshal(res, &records); err != nil {
		return nil, err
	}
	return records, nil
}

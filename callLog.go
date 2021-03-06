package termux

import (
	"bytes"
	"encoding/json"
)

// CallLogPiece represents one piece of call log
type CallLogPiece struct {
	Name     string `json:"name"`
	Number   string `json:"phone_number"`
	Type     string `json:"type"`
	Date     string `json:"date"`
	Duration string `json:"duration"`
}

// CallLog acquires call logs with a given limit and offset
func CallLog(limit int, offset int) ([]CallLogPiece, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := exec(nil, buf, "CallLog", map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}, ""); err != nil {
		return nil, err
	}
	res := buf.Bytes()
	if err := checkErr(res); err != nil {
		return nil, err
	}
	records := make([]CallLogPiece, 0)
	if err := json.Unmarshal(res, &records); err != nil {
		return nil, err
	}
	return records, nil
}

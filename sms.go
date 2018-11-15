package termux

import (
	"bytes"
	"encoding/json"
)

type SMSBoxType int

const (
	All    SMSBoxType = 0
	Inbox  SMSBoxType = 1
	Sent   SMSBoxType = 2
	Draft  SMSBoxType = 3
	Outbox SMSBoxType = 4
)

type SMS struct {
	ThreadID int    `json:"threadid"`
	Type     string `json:"type"`
	Read     bool   `json:"read"`
	Sender   string `json:"sender"`
	Number   string `json:"number"`
	Received string `json:"received"`
	Body     string `json:"body"`
}

func SMSList(limit int, offset int, box SMSBoxType) ([]SMS, error) {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "SmsInbox", map[string]interface{}{
		"type":   box,
		"limit":  limit,
		"offset": offset,
	}, "")
	res := buf.Bytes()
	if err := checkErr(res); err != nil {
		return nil, err
	}
	l := make([]SMS, 0)
	if err := json.Unmarshal(res, l); err != nil {
		return nil, err
	}
	return l, nil
}

func SMSSend(numbers []string, text string) error {
	in := bytes.NewBufferString(text)
	buf := bytes.NewBuffer([]byte{})
	exec(in, buf, "SmsSend", map[string]interface{}{
		"recipients": numbers,
	}, "")
	return checkErr(buf.Bytes())
}

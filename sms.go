package termux

import (
	"bytes"
	"encoding/json"
)

// SMSBoxType enumberates available sms box types
type SMSBoxType int

const (
	// All sms box type
	All SMSBoxType = 0
	// Inbox sms box type
	Inbox SMSBoxType = 1
	// Sent sms box type
	Sent SMSBoxType = 2
	// Draft sms box type
	Draft SMSBoxType = 3
	// Outbox sms box type
	Outbox SMSBoxType = 4
)

// SMS represents a piece of received SMS
type SMS struct {
	ThreadID int    `json:"threadid"`
	Type     string `json:"type"`
	Read     bool   `json:"read"`
	Sender   string `json:"sender"`
	Number   string `json:"number"`
	Received string `json:"received"`
	Body     string `json:"body"`
}

// SMSList acquires a list of the received SMS in the given SMS box with a given limit and offset
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

// SMSSend sends a text message to the given recipient numbers
func SMSSend(numbers []string, text string) error {
	in := bytes.NewBufferString(text)
	buf := bytes.NewBuffer([]byte{})
	exec(in, buf, "SmsSend", map[string]interface{}{
		"recipients": numbers,
	}, "")
	return checkErr(buf.Bytes())
}

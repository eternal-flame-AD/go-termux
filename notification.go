package termux

import "bytes"

// NotificationPriority enumerates the priority level of the notification
type NotificationPriority string

const (
	// Max notification priority
	Max NotificationPriority = "max"
	// High notification priority
	High NotificationPriority = "high"
	// Default notification priority
	Default NotificationPriority = "default"
	// Low notification priority
	Low NotificationPriority = "low"
	// Min notification priority
	Min NotificationPriority = "min"
)

// NotificationButton represents a button shown in the notification bar
type NotificationButton struct {
	Text   string
	Action string
}

// NotificationOpt represents the options of a notification
type NotificationOpt struct {
	Content  string
	Sound    bool
	Title    string
	Viberate []int32
	Priority NotificationPriority
	LED      struct {
		Color string
		On    int
		Off   int
	}
	Action       string
	DeleteAction string
	Btn1         NotificationButton
	Btn2         NotificationButton
	Btn3         NotificationButton
}

// NotificationRemove removes the notification with the given id
func NotificationRemove(id string) error {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "NotificationRemove", map[string]interface{}{
		"id": id,
	}, "")
	return checkErr(buf.Bytes())
}

// Notification creates a new notification with the given id and options
func Notification(id string, opt NotificationOpt) error {
	in := bytes.NewBuffer([]byte{})
	out := bytes.NewBuffer([]byte{})
	in.WriteString(opt.Content)

	param := map[string]interface{}{
		"priority": opt.Priority,
		"title":    opt.Title,
		"sound":    opt.Sound,
		"id":       id,
	}
	if opt.LED.On > 0 && opt.LED.Off > 0 {
		param["led-on"] = opt.LED.On
		param["led-off"] = opt.LED.Off
	}
	if opt.Viberate != nil {
		param["viberate"] = opt.Viberate
	}
	for key, val := range map[string]string{
		"led-color":        opt.LED.Color,
		"action":           opt.Action,
		"on_delete_action": opt.DeleteAction,
		"button_text_1":    opt.Btn1.Text,
		"button_action_1":  opt.Btn1.Action,
		"button_text_2":    opt.Btn2.Text,
		"button_action_2":  opt.Btn2.Action,
		"button_text_3":    opt.Btn3.Text,
		"button_action_3":  opt.Btn3.Action,
	} {
		if val == "" {
			continue
		}
		param[key] = val
	}

	if err := exec(in, out, "Notification", param, ""); err != nil {
		return err
	}
	return checkErr(out.Bytes())
}

package termux

import "bytes"

type NotificationPriority string

const (
	Max     NotificationPriority = "max"
	High    NotificationPriority = "high"
	Default NotificationPriority = "default"
	Low     NotificationPriority = "low"
	Min     NotificationPriority = "min"
)

type NotificationButton struct {
	Text   string
	Action string
}

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

func NotificationRemove(id string) error {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "NotificationRemove", map[string]interface{}{
		"id": id,
	}, "")
	return checkErr(buf.Bytes())
}

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

	exec(in, out, "Notification", param, "")
	return checkErr(out.Bytes())
}

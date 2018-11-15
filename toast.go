package termux

import "bytes"

type ToastPosition string

const (
	Top    ToastPosition = "top"
	Middle ToastPosition = "middle"
	Bottom ToastPosition = "bottom"
)

type ToastOption struct {
	BGColor   string
	FontColor string
	Position  ToastPosition
	Short     bool
}

func Toast(text string, opts ToastOption) error {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "Toast", map[string]interface{}{
		"text_color": opts.FontColor,
		"background": opts.BGColor,
		"short":      opts.Short,
		"gravity":    string(opts.Position),
	}, "")
	return checkErr(buf.Bytes())
}

package termux

import "bytes"

// ToastPosition enumerates current position of the toast
type ToastPosition string

const (
	// Top lifts the toast ot the top of the screen
	Top ToastPosition = "top"
	// Middle is the default position of the toast
	Middle ToastPosition = "middle"
	// Bottom puts the toast to the bottom of the screen
	Bottom ToastPosition = "bottom"
)

// ToastOption represents the optional options to a toast
type ToastOption struct {
	BGColor   string
	FontColor string
	Position  ToastPosition
	Short     bool
}

// Toast creates a toast on the device
func Toast(text string, opts ToastOption) error {
	in := bytes.NewBufferString(text)
	buf := bytes.NewBuffer([]byte{})
	params := map[string]interface{}{}
	if opts.FontColor != "" {
		params["text_color"] = opts.FontColor
	}
	if opts.BGColor != "" {
		params["background"] = opts.BGColor
	}
	if opts.Short {
		params["short"] = true
	}
	if opts.Position != "" {
		params["gravity"] = string(opts.Position)
	}
	if err := exec(in, buf, "Toast", params, ""); err != nil {
		return err
	}
	return checkErr(buf.Bytes())
}

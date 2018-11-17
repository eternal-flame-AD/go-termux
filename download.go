package termux

import "bytes"

// Download calls the system download manager to download an URL
func Download(desc string, title string, url string) error {
	buf := bytes.NewBuffer([]byte{})
	if err := exec(nil, buf, "Download", map[string]interface{}{
		"title":       title,
		"description": desc,
	}, url); err != nil {
		return err
	}
	return checkErr(buf.Bytes())
}

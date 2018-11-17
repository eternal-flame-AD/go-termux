package termux

import "bytes"

// Download calls the system download manager to download an URL
func Download(desc string, title string, url string) error {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "Download", map[string]interface{}{
		"title":       title,
		"description": desc,
	}, url)
	return checkErr(buf.Bytes())
}

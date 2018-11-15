package termux

import (
	"bytes"
	"path/filepath"
)

func WallpaperFile(path string, lockscreen bool) error {
	realpath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "Wallpaper", map[string]interface{}{
		"file":       realpath,
		"lockscreen": lockscreen,
	}, "")
	return checkErr(buf.Bytes())
}

func WallpaperURL(url string, lockscreen bool) error {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "Wallpaper", map[string]interface{}{
		"url":        url,
		"lockscreen": lockscreen,
	}, "")
	return checkErr(buf.Bytes())
}

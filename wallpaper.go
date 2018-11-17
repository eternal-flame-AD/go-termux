package termux

import (
	"bytes"
	"path/filepath"
)

// WallpaperFile sets the current wallpaper, main screen if lockscreen is false
func WallpaperFile(path string, lockscreen bool) error {
	realpath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer([]byte{})
	if err := exec(nil, buf, "Wallpaper", map[string]interface{}{
		"file":       realpath,
		"lockscreen": lockscreen,
	}, ""); err != nil {
		return err
	}
	return checkErr(buf.Bytes())
}

// WallpaperURL sets the current wallpaper, main screen if lockscreen is false
func WallpaperURL(url string, lockscreen bool) error {
	buf := bytes.NewBuffer([]byte{})
	if err := exec(nil, buf, "Wallpaper", map[string]interface{}{
		"url":        url,
		"lockscreen": lockscreen,
	}, ""); err != nil {
		return err
	}
	return checkErr(buf.Bytes())
}

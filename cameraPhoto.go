package termux

import (
	"bytes"
	"os"
	"path"
)

func TakePhoto(cameraID string, outfile string) error {
	buf := bytes.NewBuffer([]byte{})
	if !path.IsAbs(outfile) {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		outfile = path.Join(wd, outfile)
	}
	exec(nil, buf, "CameraPhoto", map[string]interface{}{
		"camera": cameraID,
		"file":   outfile,
	}, "")
	res := buf.Bytes()
	return checkErr(res)
}

package termux

import (
	"bytes"
	"os"
	"path"
)

// TakePhoto uses the device camera to take a photo and save the result to outfile
// cameraID is the id of the camera specified by the return of CameraInfo
func TakePhoto(cameraID string, outfile string) error {
	buf := bytes.NewBuffer([]byte{})
	if !path.IsAbs(outfile) {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		outfile = path.Join(wd, outfile)
	}
	if err := exec(nil, buf, "CameraPhoto", map[string]interface{}{
		"camera": cameraID,
		"file":   outfile,
	}, ""); err != nil {
		return err
	}
	res := buf.Bytes()
	return checkErr(res)
}

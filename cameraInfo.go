package termux

import (
	"bytes"
	"encoding/json"
)

// CameraInfoPiece represent the info of one camera on the device
type CameraInfoPiece struct {
	ID              string `json:"id"`
	Facing          string `json:"facing"`
	JPEGOutputSizes []struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"jpeg_output_sizes"`
	FocalLengths []float64 `json:"focal_lengths"`
	AEModes      []string  `json:"auto_exposure_modes"`
	PhysicalSize struct {
		Width  float64 `json:"width"`
		Height float64 `json:"height"`
	} `json:"physical_size"`
	Capabilities []string `json:"capabilities"`
}

// CameraInfo gets the information of available cameras on the device
func CameraInfo() ([]CameraInfoPiece, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := exec(nil, buf, "CameraInfo", nil, ""); err != nil {
		return nil, err
	}
	res := buf.Bytes()
	if err := checkErr(res); err != nil {
		return nil, err
	}
	records := make([]CameraInfoPiece, 0)
	if err := json.Unmarshal(res, &records); err != nil {
		return nil, err
	}
	return records, nil
}

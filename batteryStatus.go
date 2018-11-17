package termux

import (
	"bytes"
	"encoding/json"
)

// BatteryStatusResponse represents the current battery status
type BatteryStatusResponse struct {
	Health      string  `json:"health"`
	Percentage  int     `json:"percentage"`
	Plugged     string  `json:"plugged"`
	Status      string  `json:"status"`
	Temperature float64 `json:"temperature"`
}

// BatteryStatus acquires the corrent audio info
func BatteryStatus() (*BatteryStatusResponse, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := exec(nil, buf, "BatteryStatus", nil, ""); err != nil {
		return nil, err
	}
	res := buf.Bytes()
	if err := checkErr(res); err != nil {
		return nil, err
	}
	ret := new(BatteryStatusResponse)
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

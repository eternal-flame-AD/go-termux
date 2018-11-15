package termux

import (
	"bytes"
	"encoding/json"
)

type BatteryStatusResponse struct {
	Health      string  `json:"health"`
	Percentage  int     `json:"percentage"`
	Plugged     string  `json:"plugged"`
	Status      string  `json:"status"`
	Temperature float64 `json:"temperature"`
}

func BatteryStatus() (*BatteryStatusResponse, error) {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "BatteryStatus", nil, "")
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

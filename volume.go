package termux

import (
	"bytes"
	"encoding/json"
)

type AudioStreamState struct {
	Name      string `json:"stream"`
	Volume    int    `json:"volume"`
	MaxVolume int    `json:"max_volume"`
}

func AudioStreams() ([]AudioStreamState, error) {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "Volume", nil, "")
	res := buf.Bytes()

	if err := checkErr(res); res != nil {
		return nil, err
	}
	l := make([]AudioStreamState, 0)
	if err := json.Unmarshal(res, l); err != nil {
		return nil, err
	}
	return l, nil
}

func AudioStreamVolume(name string, volume int) error {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "Volume", map[string]interface{}{
		"stream": name,
		"volume": volume,
	}, "")
	return checkErr(buf.Bytes())
}

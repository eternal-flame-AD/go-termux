package termux

import (
	"bytes"
	"encoding/json"
)

type AudioInfoResponse struct {
	JavaOutputSampleRate       string `json:"PROPERTY_OUTPUT_SAMPLE_RATE"`
	JavaOutputFramesPerBuffer  string `json:"PROPERTY_OUTPUT_FRAMES_PER_BUFFER"`
	AudioTrackOutputSampleRate int    `json:"AUDIOTRACK_NATIVE_OUTPUT_SAMPLE_RATE"`
	BluetoothA2DP              bool   `json:"BLUETOOTH_A2DP_IS_ON"`
	WiredHeadsetConnected      bool   `json:"WIREDHEADSET_IS_CONNECTED"`
}

func AudioInfo() (*AudioInfoResponse, error) {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "AudioInfo", nil, "")
	res := buf.Bytes()
	if err := checkErr(res); err != nil {
		return nil, err
	}
	ret := new(AudioInfoResponse)
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

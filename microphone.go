package termux

import (
	"bytes"
	"encoding/json"
	"os"
	"path"
)

// RecordInfo represent the current state of recording process
type RecordInfo struct {
	IsRecording bool   `json:"isRecording"`
	OutputFile  string `json:"outputFile"`
}

// Record records from microphone, duration must be in seconds
func Record(duration int, outFile string) error {
	buf := bytes.NewBuffer([]byte{})
	if !path.IsAbs(outFile) {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		outFile = path.Join(wd, outFile)
	}
	if err := execAction("MicRecorder", nil, buf, "record", map[string]interface{}{
		"limit": duration * 1000,
		"file":  outFile,
	}); err != nil {
		return err
	}
	res := buf.Bytes()
	return checkErr(res)
}

// GetRecordInfo returns the current state of the recording process
func GetRecordInfo() (RecordInfo, error) {
	var info RecordInfo
	buf := bytes.NewBuffer([]byte{})
	if err := execAction("MicRecorder", nil, buf, "info", nil); err != nil {
		return RecordInfo{}, err
	}
	res := buf.Bytes()
	if err := checkErr(res); err != nil {
		return RecordInfo{}, err
	}
	if err := json.Unmarshal(res, &info); err != nil {
		return RecordInfo{}, err
	}
	return info, nil
}

// StopRecord stop the recording process
func StopRecord() error {
	buf := bytes.NewBuffer([]byte{})
	if err := execAction("MicRecorder", nil, buf, "quit", nil); err != nil {
		return err
	}
	res := buf.Bytes()
	if err := checkErr(res); err != nil {
		return err
	}
	return nil
}

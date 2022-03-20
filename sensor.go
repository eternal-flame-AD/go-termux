package termux

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/eternal-flame-AD/go-termux/internal/chanbuf"
)

// SensorList acquires a list of available sensors on the device
func SensorList() ([]string, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := execAction("Sensor", nil, buf, "list", nil); err != nil {
		return nil, err
	}
	res := buf.Bytes()

	if err := checkErr(res); err != nil {
		return nil, err
	}
	l := new(struct {
		Sensors []string `json:"sensors"`
	})
	if err := json.Unmarshal(res, l); err != nil {
		return nil, err
	}
	return l.Sensors, nil
}

// SensorWatchOpt represents the options to a Sensor call
type SensorWatchOpt struct {
	Limit      int
	DelayMS    int
	SensorList []string
}

// Sensor starts a sensor watch in a given context and options
// returns raw data bytes encooded with JSON
func Sensor(ctx context.Context, opt SensorWatchOpt) (<-chan []byte, error) {
	response := make(chan []byte)
	param := map[string]interface{}{}
	if opt.SensorList == nil {
		param["all"] = true
	} else {
		param["sensors"] = strings.Join(opt.SensorList, ",")
	}
	if opt.DelayMS != 0 {
		param["dalay"] = opt.DelayMS
	}
	if opt.Limit != 0 {
		param["limit"] = opt.Limit
	}
	if err := execContext(ctx, nil, chanbuf.BufToChan{
		C: response,
	}, "Sensor", param, ""); err != nil {
		return nil, err
	}

	go func() {
		defer execAction("Sensor", nil, bytes.NewBuffer([]byte{}), "cleanup", nil)
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}()

	return response, nil
}

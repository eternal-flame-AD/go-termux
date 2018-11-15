package termux

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/eternal-flame-AD/go-termux/internal/chanbuf"
)

func SensorList() ([]string, error) {
	buf := bytes.NewBuffer([]byte{})
	execAction("Sensor", nil, buf, "list")
	res := buf.Bytes()

	if err := checkErr(res); res != nil {
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

type SensorWatchOpt struct {
	Limit      int
	DelayMS    int
	SensorList []string
}

func Sensor(ctx context.Context, opt SensorWatchOpt) <-chan []byte {
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
	execContext(ctx, nil, chanbuf.BufToChan{response}, "Sensor", param, "")

	go func() {
		defer execAction("Sensor", nil, bytes.NewBuffer([]byte{}), "cleanup")
		for {
			select {
			case <-ctx.Done():
				return
			}
		}
	}()

	return response
}

package termux

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/eternal-flame-AD/go-termux/internal/chanbuf"
)

type LocationProvider string

const (
	GPS     LocationProvider = "gps"
	Network LocationProvider = "network"
	Passive LocationProvider = "passive"
)

type LocationRecord struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Accuracy  float64 `json:"accuracy"`
	Bearing   float64 `json:"bearing"`
	Speed     float64 `json:"speed"`
	ElapsedMS int     `json:"elapsedMs"`
	Provider  string  `json:"provider"`
}

func location(ctx context.Context, t string, provider LocationProvider) (*LocationRecord, error) {
	buf := bytes.NewBuffer([]byte{})
	execContext(ctx, nil, buf, "Location", map[string]interface{}{
		"provider": string(provider),
		"request":  t,
	}, "")
	res := buf.Bytes()
	if err := checkErr(res); err != nil {
		return nil, err
	}
	r := new(LocationRecord)
	if err := json.Unmarshal(res, r); err != nil {
		return nil, err
	}
	return r, nil
}

func LastLocation(ctx context.Context, provider LocationProvider) (*LocationRecord, error) {
	return location(ctx, "last", provider)
}

func Location(ctx context.Context, provider LocationProvider) (*LocationRecord, error) {
	return location(ctx, "once", provider)
}

func UpdatedLocation(ctx context.Context, provider LocationProvider) <-chan struct {
	Location *LocationRecord
	Error    error
} {
	response := make(chan []byte)
	go execContext(ctx, nil, chanbuf.BufToChan{response}, "Location", map[string]interface{}{
		"provider": string(provider),
		"request":  "updates",
	}, "")
	ret := make(chan struct {
		Location *LocationRecord
		Error    error
	})
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ret)
				return
			case data := <-response:
				if err := checkErr(data); err != nil {
					ret <- struct {
						Location *LocationRecord
						Error    error
					}{nil, err}
					continue
				}
				l := new(LocationRecord)
				if err := json.Unmarshal(data, l); err != nil {
					ret <- struct {
						Location *LocationRecord
						Error    error
					}{nil, err}
					continue
				}
				ret <- struct {
					Location *LocationRecord
					Error    error
				}{l, nil}
			}
		}
	}()
	return ret
}

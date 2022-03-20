package termux

import (
	"bytes"
	"encoding/json"
)

// WifiEnable sets the current enabled state of the device wifi
func WifiEnable(enabled bool) error {
	buf := bytes.NewBuffer([]byte{})
	exec(nil, buf, "WifiEnable", map[string]interface{}{
		"enabled": enabled,
	}, "")
	res := buf.Bytes()
	return checkErr(res)
}

// WifiConnection represents the current connection info of the device wifi interface
type WifiConnection struct {
	BSSID           string `json:"bssid"`
	FreqMHZ         int    `json:"frequency_mhz"`
	IP              string `json:"ip"`
	LinkSpeedMbps   int    `json:"link_speed_mbps"`
	MACAddr         string `json:"mac_address"`
	NetworkID       int    `json:"network_id"`
	RSSI            int    `json:"rssi"`
	SSID            string `json:"ssid"`
	SSIDHidden      bool   `json:"ssid_hidden"`
	SupplicantState string `json:"supplicant_state"`
}

// WifiConnectionState returns the current wifi connection state of the device
func WifiConnectionState() (*WifiConnection, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := exec(nil, buf, "WifiConnectionInfo", nil, ""); err != nil {
		return nil, err
	}
	res := buf.Bytes()
	if err := checkErr(res); err != nil {
		return nil, err
	}
	ret := new(WifiConnection)
	if err := json.Unmarshal(res, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// WifiAP represents a discovered AP in a wifi scan
type WifiAP struct {
	BSSID     string `json:"bssid"`
	FreqMHZ   int    `json:"frequency_mhz"`
	RSSI      int    `json:"rssi"`
	SSID      string `json:"ssid"`
	TimeStamp int64  `json:"timestamp"`
	// CenterFreqMHZ not used for 20Mhz bands
	CenterFreqMHZ    int    `json:"center_frequency_mhz"`
	ChannelBandwidth string `json:"channel_bandwidth_mhz"`
}

// WifiScan scans for available networks
func WifiScan() ([]WifiAP, error) {
	buf := bytes.NewBuffer([]byte{})
	execAction("WifiScanInfo", nil, buf, "list")
	res := buf.Bytes()

	if err := checkErr(res); err != nil {
		return nil, err
	}
	l := make([]WifiAP, 0)
	if err := json.Unmarshal(res, l); err != nil {
		return nil, err
	}
	return l, nil
}

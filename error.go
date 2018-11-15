package termux

import (
	"encoding/json"
	"errors"
)

func checkErr(data []byte) error {
	x := struct {
		Error    *string `json:"error"`
		APIError *string `json:"API_ERROR"`
	}{}
	json.Unmarshal(data, &x)
	if x.Error != nil {
		return errors.New(*(x.Error))
	}
	if x.APIError != nil {
		return errors.New(*(x.APIError))
	}
	return nil
}

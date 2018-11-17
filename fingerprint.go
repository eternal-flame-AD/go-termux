package termux

import "errors"

// Fingerprint uses the fingerprint sensor on the device for authentication
// Not implemented as the original API does not work on my device
func Fingerprint() error {
	return errors.New("not implemented")
}

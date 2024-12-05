package providers

import (
	"errors"
)

func decodeChromecast(typeQuery string, id string) (string, error) {
	switch typeQuery {
	case "pr":
		return "https://ccp-lh.googleusercontent.com/proxy/" + id + "=w0", nil

	case "pp":
		return "https://ccp-lh.googleusercontent.com/chromecast-private-photos/" + id + "=w0", nil

	default:
		return "", errors.New("invalid type")
	}
}

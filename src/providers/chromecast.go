package providers

import (
	"errors"
	"net/url"
)

func decodeChromecast(queries url.Values) (string, error) {
	switch queries.Get("type") {
	case "pr":
		return "https://ccp-lh.googleusercontent.com/proxy/" + queries.Get("id") + "=w0", nil

	case "pp":
		return "https://ccp-lh.googleusercontent.com/chromecast-private-photos/" + queries.Get("id") + "=w0", nil

	default:
		return "", errors.New("invalid type")
	}
}

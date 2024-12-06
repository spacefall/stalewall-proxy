package providers

import (
	"errors"
	"github.com/spacefall/stalewall-proxy/src/utils"
)

func decodeChromecast(typeQuery string, id string) ([]byte, error) {
	var url string
	switch typeQuery {
	case "pr":
		url = "https://ccp-lh.googleusercontent.com/proxy/" + id + "=w0"

	case "pp":
		url = "https://ccp-lh.googleusercontent.com/chromecast-private-photos/" + id + "=w0"

	default:
		return nil, errors.New("invalid type")
	}
	return utils.Fetcher(url)
}

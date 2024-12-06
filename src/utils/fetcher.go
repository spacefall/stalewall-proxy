package utils

import (
	"errors"
	"io"
	"net/http"
)

func Fetcher(url string) ([]byte, error) {
	// fetch image
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("error fetching image")
	}
	// this doesn't really return a fixable error, so I guess we can ignore it (?)
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("error fetching image: " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

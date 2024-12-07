package providers

import "github.com/spacefall/stalewall-proxy/src/utils"

func decodeApod(_ string, id string) ([]byte, error) {
	url := "https://apod.nasa.gov/apod/image/" + id + ".jpg"
	return utils.Fetcher(url)
}

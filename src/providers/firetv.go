package providers

import "github.com/spacefall/stalewall-proxy/src/utils"

func decodeFireTV(_ string, id string) ([]byte, error) {
	url := "https://d21m0ezw6fosyw.cloudfront.net/" + id + ".jpg"
	return utils.Fetcher(url)
}

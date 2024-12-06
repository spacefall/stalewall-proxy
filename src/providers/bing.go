package providers

import "github.com/spacefall/stalewall-proxy/src/utils"

func decodeBing(_ string, id string) ([]byte, error) {
	url := "https://bing.com/th?id=" + id + ".jpg&p=0&pid=hp&qlt=100"
	return utils.Fetcher(url)
}

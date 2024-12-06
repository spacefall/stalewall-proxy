package providers

import (
	"errors"
	"github.com/spacefall/stalewall-proxy/src/utils"
)

func decodeSpotlight(typeQuery string, id string) ([]byte, error) {
	var url string
	switch typeQuery {
	case "ip":
		url = "https://img-prod-cms-rt-microsoft-com.akamaized.net/cms/api/am/imageFileData/" + id

	case "qp":
		url = "https://query.prod.cms.rt.microsoft.com/cms/api/am/binary/" + id

	case "rp":

		url = "https://res.public.onecdn.static.microsoft/creativeservice/" + id + ".jpg"

	default:
		return nil, errors.New("invalid type")
	}
	return utils.Fetcher(url)
}

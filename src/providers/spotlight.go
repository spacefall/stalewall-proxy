package providers

import "errors"

func decodeSpotlight(typeQuery string, id string) (string, error) {
	switch typeQuery {
	case "ip":
		return "https://img-prod-cms-rt-microsoft-com.akamaized.net/cms/api/am/imageFileData/" + id, nil

	case "qp":
		return "https://query.prod.cms.rt.microsoft.com/cms/api/am/binary/" + id, nil

	case "rp":
		return "https://res.public.onecdn.static.microsoft/creativeservice/" + id + ".jpg", nil

	default:
		return "", errors.New("invalid type")
	}
}

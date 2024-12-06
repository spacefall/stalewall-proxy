package providers

import (
	"encoding/base64"
	"github.com/spacefall/stalewall-proxy/src/utils"
	"github.com/tidwall/gjson"
	"strings"
)

func decodeEarthView(_ string, id string) ([]byte, error) {
	url := "https://www.gstatic.com/prettyearth/assets/data/v3/" + id + ".json"
	jsonBytes, err := utils.Fetcher(url)
	if err != nil {
		return nil, err
	}

	json := gjson.GetBytes(jsonBytes, "dataUri")
	// it's always a jpeg
	pureB64 := strings.TrimPrefix(json.String(), "data:image/jpeg;base64,")
	return base64.StdEncoding.DecodeString(pureB64)
}

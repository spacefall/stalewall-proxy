package providers

import "net/url"

var Providers = map[string]func(url.Values) (string, error){
	"chromecast": decodeChromecast,
}

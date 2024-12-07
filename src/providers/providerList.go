package providers

var Providers = map[string]func(string, string) ([]byte, error){
	"chromecast": decodeChromecast,
	"firetv":     decodeFireTV,
	"spotlight":  decodeSpotlight,
	"bing":       decodeBing,
	"earthview":  decodeEarthView,
	"apod":       decodeApod,
}

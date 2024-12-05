package providers

var Providers = map[string]func(string, string) (string, error){
	"chromecast": decodeChromecast,
	"firetv":     decodeFireTV,
}

package src

import (
	"encoding/base64"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	pr "github.com/spacefall/stalewall-proxy/src/providers"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"strconv"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// just syntactic sugar
	queries := r.URL.Query()

	// get provider
	prov := queries.Get("prov")
	if prov == "" {
		http.Error(w, "no provider specified", http.StatusBadRequest)
		return
	}

	// check if provider is valid
	if _, valid := pr.Providers[prov]; !valid {
		http.Error(w, "provider specified ("+prov+") is invalid", http.StatusBadRequest)
		return
	}

	// b64 id -> plain text id
	base64id := queries.Get("id")
	id, err := base64.StdEncoding.DecodeString(base64id)
	if err != nil {
		http.Error(w, "error decoding base64 id", http.StatusBadRequest)
		return
	}

	// decode url with provider decoder
	url, err := pr.Providers[prov](queries.Get("type"), string(id))
	if err != nil {
		http.Error(w, "error with provider "+prov+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	println(url)

	// fetch image
	res, err := http.Get(url)
	if err != nil {
		http.Error(w, "error fetching image", http.StatusInternalServerError)
		return
	}
	if res.StatusCode != http.StatusOK {
		http.Error(w, "error fetching image: "+res.Status, res.StatusCode)
		return
	}

	// this doesn't really return a fixable error, so I guess we can ignore it (?)
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	// allows cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Content-Type", "image/jpeg")

	if hStr, wStr := queries.Get("h"), queries.Get("w"); hStr != "" && wStr != "" {
		// decode height and width
		height, err := strconv.Atoi(hStr)
		if err != nil || height <= 0 {
			http.Error(w, "invalid height", http.StatusBadRequest)
			return
		}
		width, err := strconv.Atoi(wStr)
		if err != nil || width <= 0 {
			http.Error(w, "invalid width", http.StatusBadRequest)
			return
		}

		// decode the image
		img, _, err := image.Decode(res.Body)
		if err != nil {
			http.Error(w, "error decoding image", http.StatusInternalServerError)
			return
		}

		// resize setup
		resizer := nfnt.NewDefaultResizer()
		analyzer := smartcrop.NewAnalyzer(resizer)

		// find best crop
		topCrop, _ := analyzer.FindBestCrop(img, width, height)
		// crop
		img = img.(SubImager).SubImage(topCrop)
		// resize
		if img.Bounds().Dx() > width || img.Bounds().Dy() > height {
			img = resizer.Resize(img, uint(width), uint(height))
		}

		// encode to response
		err = jpeg.Encode(w, img, nil)
		if err != nil {
			http.Error(w, "error encoding image for response", http.StatusInternalServerError)
			return
		}
		return
	}

	// if not cropped, copy image to response
	_, err = io.Copy(w, res.Body)
	if err != nil {
		http.Error(w, "error responding with image", http.StatusInternalServerError)
		return
	}
}

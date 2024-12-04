package src

import (
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

	// decode url with provider decoder
	url, err := pr.Providers[prov](queries)
	if err != nil {
		http.Error(w, "error with provider "+prov+": "+err.Error(), http.StatusInternalServerError)
		return
	}

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

	defer res.Body.Close()

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

package handler

import (
	"github.com/spacefall/stalewall-proxy/src"
	"net/http"
)

//goland:noinspection GoUnusedExportedFunction
func Handler(w http.ResponseWriter, r *http.Request) {
	src.Handler(w, r)
}

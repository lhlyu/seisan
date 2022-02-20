package api

import (
	"fmt"
	"net/http"
)

const version = "v0.0.1"

func HandlerVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"version":"%s"}`, version)
}

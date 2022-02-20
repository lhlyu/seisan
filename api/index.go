package seisan

import (
	"fmt"
	"net/http"
	"seisan/internal/core"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"hello":"%s"}`, core.Version)
}

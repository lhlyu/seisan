package handler

import (
	"fmt"
	"net/http"
	"seisan/api/internal/core"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"hello":"%s"}`, core.Version)
}

package handler

import (
	"fmt"
	"github.com/lhlyu/seisan/internal/core"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"hello":"%s"}`, core.Version)
}

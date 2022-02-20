package handler

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	io.WriteString(w, `{"hello":"yu"}`)
}

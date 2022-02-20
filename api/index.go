package api

import (
	"io"
	"net/http"
	"seisan/internal/demo"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	values := r.URL.Query()
	d := values.Get("d")
	v, _ := strconv.Atoi(d)
	io.WriteString(w, demo.GetLabel(v))
}

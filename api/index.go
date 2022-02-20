package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seisan/internal/demo"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	header, _ := json.Marshal(r.Header)
	fmt.Fprintf(w, `{"a": "%s","hello":"%s", "header": "%s"}`, demo.Demo, demo.S, string(header))
}

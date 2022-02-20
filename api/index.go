package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var s string

func init() {
	b, _ := ioutil.ReadFile("./demo.txt")
	s = string(b)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	header, _ := json.Marshal(r.Header)
	fmt.Fprintf(w, `{"hello":"%s", "header"}`, s, string(header))
}

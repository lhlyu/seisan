package api

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
)

var s string

const a = 1

func init() {
	b, err := ioutil.ReadFile("./demo.txt")
	if err != nil {
		s = err.Error()
		return
	}
	s = string(b)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	header, _ := json.Marshal(r.Header)
	fmt.Fprintf(w, `{"a": "%s","hello":"%s", "header": "%s"}`, cast.ToString(a), s, string(header))
}
